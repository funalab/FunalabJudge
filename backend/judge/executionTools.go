package judge

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func updateSubmissionStatus(c *gin.Context, sId int, status string) {
	client, exists := c.Get("mongoClient")
	if !exists {
		log.Fatal("DB client is not available.")
	}
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(usrCol)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": sId}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func updateSubmissionResult(c *gin.Context, sId int, tId int, status string) {
	client, exists := c.Get("mongoClient")
	if !exists {
		log.Fatal("DB client is not available.")
	}
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := (client.(*mongo.Client)).Database(dbName).Collection(subCol)

	filter := bson.M{"id": sId}
	update := bson.M{
		"$set": bson.M{
			"results.$[elem].status": status,
		},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.testCaseId": tId}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update, &options.UpdateOptions{
		ArrayFilters: &arrayFilters,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func compareWithAnswer(output string, answer string) bool {
	fixedOutput := strings.TrimRight(output, "\n")
	fixedAnswer := strings.TrimRight(answer, "\n")
	return fixedOutput == fixedAnswer
}

func execCommand(sId int, command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = filepath.Join(os.Getenv("EXEC_DIR"), strconv.Itoa(sId))

	output, err := cmd.CombinedOutput()
	return string(output), err
}

func execCommandWithInput(sId int, command string, input string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Microsecond)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = filepath.Join(os.Getenv("EXEC_DIR"), strconv.Itoa(sId))

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	stdin.Write([]byte(input + "\n"))
	stdin.Close()
	output, err := cmd.CombinedOutput()

	return string(output), err
}

func searchExecutableFile(sId int) (string, error) {
	var executableFiles []string
	targDir := filepath.Join(os.Getenv("EXEC_DIR"), strconv.Itoa(sId))
	err := filepath.Walk(targDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Mode().Perm()&0111 != 0 {
			executableFiles = append(executableFiles, info.Name())
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if len(executableFiles) > 1 {
		return "", errors.New("not only one executable files exists")
	} else if len(executableFiles) == 0 {
		return "", errors.New("executable files does not found")
	}
	return executableFiles[0], nil
}
