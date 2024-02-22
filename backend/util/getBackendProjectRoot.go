package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go-test/myTypes"
)

var goModFile string = "go.mod"

func GetBackendProjectRoot() string {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		log.Fatalf("Failed to get current working directory's absolute path: %v\n", err.Error())
	}

	root, err := findRootDir(currentDir)
	if err != nil {
		log.Fatalf("Failed to get project root dir: %v\n", err.Error())
	}

	return root
}

func findRootDir(currentDir string) (string, error) {
	goModPath := filepath.Join(currentDir, goModFile)
	_, err := os.Stat(goModPath)
	if err == nil {
		return currentDir, nil
	}

	parentDir := filepath.Dir(currentDir)

	if parentDir == currentDir {
		return "", myTypes.NewFindRootDirErr(fmt.Sprintf("Failed to get root path."))
	}

	return findRootDir(parentDir)
}
