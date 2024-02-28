package judge

import (
	"context"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func updateSubmissionStatus(client *mongo.Client, sId primitive.ObjectID, status string) {
	dbName := os.Getenv("DB_NAME")
	usrCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(usrCol)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": sId}, bson.M{"$set": bson.M{"status": status}})
	if err != nil {
		log.Fatal(err.Error())
	}
}

func updateSubmissionResult(client *mongo.Client, sId primitive.ObjectID, tId int, status string) {
	dbName := os.Getenv("DB_NAME")
	subCol := os.Getenv("SUBMISSION_COLLECTION")
	collection := client.Database(dbName).Collection(subCol)

	filter := bson.M{"_id": sId}
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

func compareWithAnswer(output string, answer string) bool {
	fixedOutput := strings.TrimRight(output, "\n")
	fixedAnswer := strings.TrimRight(answer, "\n")
	return fixedOutput == fixedAnswer
}

func execCommand(sId primitive.ObjectID, command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = filepath.Join(os.Getenv("EXEC_DIR"), sId.Hex())

	output, err := cmd.CombinedOutput()
	return string(output), err
}

func execCommandWithInput(sId primitive.ObjectID, command string, input string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = filepath.Join(os.Getenv("EXEC_DIR"), sId.Hex())

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	stdin.Write([]byte(input + "\n"))
	stdin.Close()
	output, err := cmd.CombinedOutput()

	return string(output), err
}

func searchExecutableFile(sId primitive.ObjectID) (string, error) {
	var executableFiles []string
	targDir := filepath.Join(os.Getenv("EXEC_DIR"), sId.Hex())
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
