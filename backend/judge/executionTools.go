package judge

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func compareWithAnswer(output string, answer string) bool {
	fixedOutput := strings.TrimRight(output, "\n")
	fixedAnswer := strings.TrimRight(answer, "\n")
	return fixedOutput == fixedAnswer
}

func execCommand(sId primitive.ObjectID, command string, execTime int) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(execTime))
	defer cancel()
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = filepath.Join(os.Getenv("EXEC_DIR"), sId.Hex())
	output, err := cmd.CombinedOutput()

	return output, err
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
