package judge

import (
	"fmt"
	"go-test/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func isHaveMakeFile(sId string) bool {
	files, err := os.ReadDir(filepath.Join(os.Getenv("EXEC_DIR"), sId))
	if err != nil {
		log.Fatalf(err.Error())
		return false
	}
	for _, file := range files {
		if file.Name() == os.Getenv("MAKEFILE_NAME") {
			return true
		}
	}
	return false
}

func writeMakeFile(sId string) error {
	path := filepath.Join(os.Getenv("EXEC_DIR"), sId, "Makefile")
	makefile, err := os.Create(path)
	if err != nil {
		log.Println("Failed to write make file.")
		return util.NewGenerateMakefileErr(fmt.Sprintf("Failed to generate makefile: %v\n", err.Error()))
	}
	defer makefile.Close()

	err = writeOptions(makefile, sId)
	if err != nil {
		log.Println("Failed to write makefile header.")
		return util.NewGenerateMakefileErr(fmt.Sprintf("Failed to write makefile header: %v\n", err.Error()))
	}
	err = writeTargets(makefile)
	if err != nil {
		log.Println("Failed to write make targets.")
		return util.NewGenerateMakefileErr(fmt.Sprintf("Failed to write makefile targets: %v\n", err.Error()))
	}
	return nil
}

func writeOptions(mf *os.File, sId string) error {
	err := writePROG(mf)
	if err != nil {
		return err
	}
	err = writeOBJS(mf, sId)
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
	_, err := io.WriteString(mf, fmt.Sprintf("PROG = %s\n", os.Getenv("MAKEFILE_PROG_DEFAULT")))
	if err != nil {
		return err
	}
	return nil
}

func writeOBJS(mf *os.File, sId string) error {
	_, err := io.WriteString(mf, "OBJS = ")
	if err != nil {
		return err
	}
	files, err := os.ReadDir(filepath.Join(os.Getenv("EXEC_DIR"), sId))
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.Name() == os.Getenv("MAKEFILE_NAME") {
			continue
		}
		main := strings.Split(file.Name(), ".")[0]
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
