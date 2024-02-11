package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var goModFile string = "go.mod"

type FindRootDirErr struct {
	msg string
}

type GetCurrentDirErr struct {
	msg string
}

func (f *FindRootDirErr) Error() string {
	return f.msg
}

func (g *GetCurrentDirErr) Error() string {
	return g.msg
}

func NewFindRootDirErr(msg string) error {
	return &FindRootDirErr{msg: msg}
}

func NewGetCurrentDirErr(msg string) error {
	return &GetCurrentDirErr{msg: msg}
}

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
		return "", NewFindRootDirErr(fmt.Sprintf("Failed to get root path."))
	}

	return findRootDir(parentDir)
}
