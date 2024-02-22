package execute

import (
	"fmt"
	"go-test/myTypes"
	"go-test/util"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

/* It might move into .env file. */
var makefile string = "Makefile"
var compileResourceDirPath = filepath.Join(util.GetProjectRoot(), "compile_resource")

func GenerateMakefile(c *gin.Context) {
	// 既にファイルは保存されている
	// Makefileがなければ作成する

	if !isHaveMakeFile(names) {
		m, err := writeMakeFile(names)
		if err != nil {
			log.Fatalf("Failed to write make file")
		}
		f = append(f, m)
	}

	/* 3. Execute compile by running make */
	err := execMake()
	if err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{"problemId": problemId})
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

func writeMakeFile(names []string) (*os.File, error) {
	path := filepath.Join(compileResourceDirPath, "Makefile")
	makefile, err := os.Create(path)
	if err != nil {
		log.Println("Failed to write make file.")
		return nil, myTypes.NewGenerateMakefileErr(fmt.Sprintf("Failed to generate makefile: %v\n", err.Error()))
	}
	defer makefile.Close()

	err = writeOptions(makefile, names)
	if err != nil {
		log.Println("Failed to write makefile header.")
		return nil, myTypes.NewGenerateMakefileErr(fmt.Sprintf("Failed to write makefile header: %v\n", err.Error()))
	}
	err = writeTargets(makefile)
	if err != nil {
		log.Println("Failed to write make targets.")
		return nil, myTypes.NewGenerateMakefileErr(fmt.Sprintf("Failed to write makefile targets: %v\n", err.Error()))
	}
	return makefile, nil
}

func writeOptions(mf *os.File, names []string) error {
	err := writePROG(mf)
	if err != nil {
		return err
	}
	err = writeOBJS(mf, names)
	if err != nil {
		return err
	}
	err = writeCC(mf)
	if err != nil {
		return err
	}
	err = writeCFLAGS(mf)
	if err != nil {
		return err
	}
	err = writeLDFLAGS(mf)
	if err != nil {
		return err
	}
	return nil
}

func writeTargets(mf *os.File) error {
	aa := ".PHONY: all\n"
	ac := "all: $(PROG)\n"
	_, err := io.WriteString(mf, aa)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	_, err = io.WriteString(mf, ac)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}

	sa := ".SUFFIXES: .o .c\n"
	sc := ".c.o:\n"
	_, err = io.WriteString(mf, sa)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	_, err = io.WriteString(mf, sc)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	cf := "\t$(CC) $(CFLAGS) -c $<\n"
	pr := "$(PROG): $(OBJS)\n"
	_, err = io.WriteString(mf, cf)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	_, err = io.WriteString(mf, pr)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}

	cc := "\t$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)\n"
	_, err = io.WriteString(mf, cc)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	cla := ".PHONY: clean\n"
	cl := "clean:\n"
	clc := "\trm -rf $(OBJS) $(PROG)\n"
	_, err = io.WriteString(mf, cla)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	_, err = io.WriteString(mf, cl)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	_, err = io.WriteString(mf, clc)
	if err != nil {
		log.Println("Failed to write targets")
		return err
	}
	return nil
}

func writePROG(mf *os.File) error {
	_, err := io.WriteString(mf, "PROG = final\n")
	if err != nil {
		return err
	}
	return nil
}

func writeOBJS(mf *os.File, names []string) error {
	_, err := io.WriteString(mf, "OBJS = ")
	if err != nil {
		return err
	}
	for _, name := range names {
		main := strings.Split(name, ".")[0]
		obj := main + ".o" + " "
		_, err := io.WriteString(mf, obj)
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(mf, "\n")
	if err != nil {
		return err
	}
	return nil
}

func writeCC(mf *os.File) error {
	_, err := io.WriteString(mf, "CC = gcc\n")
	if err != nil {
		return err
	}
	return nil
}

func writeCFLAGS(mf *os.File) error {
	_, err := io.WriteString(mf, "CFLAGS = -Wall\n")
	if err != nil {
		return err
	}
	return nil
}

func writeLDFLAGS(mf *os.File) error {
	_, err := io.WriteString(mf, "LDFLAGS = -lm\n")
	if err != nil {
		return err
	}
	return nil
}

func execMake() error {
	err := os.Chdir(compileResourceDirPath)
	if err != nil {
		log.Printf("Failed to change directory: %v\n", err.Error())
		return myTypes.NewMakeFailErr(fmt.Sprintf("Failed to change directory: %v\n", err.Error()))
	}
	cmd := exec.Command("make")
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Failed to execute make command: %v\n", err.Error())
		return myTypes.NewMakeFailErr(fmt.Sprintf("Failed to execute make command: %v\n", err.Error()))
	}
	/*Confirm output of make command*/
	fmt.Printf("Make output: %v\n", string(output))
	return nil
}
