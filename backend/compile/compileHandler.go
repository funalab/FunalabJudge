package compile

import (
	"fmt"
	"go-test/types"
	"go-test/util"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

/* It might move into .env file. */
var makefile string = "Makefile"
var compileResourceDirPath = filepath.Join(util.GetProjectRoot(), "compile_resource")

/*
* PROJECT_ROOT/compile_resource is a directory for executing compile.
* Here, in short, call this directory as CDIR.
*
* Logic is here.
*
* 1. Clean CDIR.
* 2. Create CDIR.
* 3. Execute compile.
*
* */

func CompileHandler(c *gin.Context) {
	var cr types.CompileRequest
	if err := c.BindJSON(&cr); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	names := cr.Names
	contents := cr.Contents

	if !isSameNumber(names, contents) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Your input is not valid: isSameNumber()",
		})
		return
	}

	/*1 Clean CDIR*/
	cleanCompileResourceDir()

	/* 2 Create CDIR and files for compile*/
	/* If not have make file, generate make file and append filelist, in short, f. */
	createCompileResourceDir()

	f := createFiles(names, contents)

	//TODO: Write make file
	if !isHaveMakeFile(names) {
		m := generateMakeFile(names)
		f = append(f, m)
	}

	/* 3. Execute compile by running make */
	//TODO: Work well, but in some case connection would be failed.
	err := execMake()
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "Make success"})
}

func isSameNumber(names []string, contents []string) bool {
	return len(names) == len(contents)
}

func isHaveMakeFile(names []string) bool {
	for _, name := range names {
		if name == makefile {
			return true
		}
	}
	return false
}

func translateIntoCompileResourceDirPath(pathFragment string) string {
	return filepath.Join(compileResourceDirPath, pathFragment)
}

func cleanCompileResourceDir() {
	os.RemoveAll(compileResourceDirPath)

	_, err := os.Stat(compileResourceDirPath)
	if err == nil {
		fmt.Printf("%v exists\n", compileResourceDirPath)
	} else {
		if os.IsNotExist(err) {
			fmt.Printf("%v does not exist\n", compileResourceDirPath)
		} else {
			fmt.Printf("Any other error: %v\n", err.Error())
		}
	}
}

func createCompileResourceDir() {
	var permission os.FileMode = 0755
	err := os.Mkdir(compileResourceDirPath, permission)
	if err != nil {
		if _, error := os.Stat(compileResourceDirPath); error != nil {
			log.Fatalf("Failed to create %v\n", compileResourceDirPath)
		}
	}
}

func createFiles(names []string, contents []string) []*os.File {
	var files []*os.File
	for index, name := range names {
		parsedName := translateIntoCompileResourceDirPath(name)
		file, err := os.Create(parsedName)
		if err != nil {
			log.Fatalf("Failed to create file: ", err.Error())
			return nil
		}
		defer file.Close()

		_, err = file.WriteString(contents[index])
		if err != nil {
			log.Fatalf("Failed to write string into file.")
			return nil
		}
		files = append(files, file)
	}
	return files
}

/* TODO: Prepare make file template */
func generateMakeFile(names []string) *os.File {
	return nil
}

func execMake() error {
	err := os.Chdir(compileResourceDirPath)
	if err != nil {
		log.Printf("Failed to change directory: %v\n", err.Error())
		return types.NewMakeFailErr(fmt.Sprintf("Failed to change directory: %v\n", err.Error()))
	}
	cmd := exec.Command("g++", "test.cpp")
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Failed to execute make command: %v\n", err.Error())
		return types.NewMakeFailErr(fmt.Sprintf("Failed to execute make command: %v\n", err.Error()))
	}
	/*Confirm output of make command*/
	fmt.Printf("Make output: %v\n", output)
	return nil
}
