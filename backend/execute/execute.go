package execute

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
* Description
* This function is helper function.
* This function is responsible for executing one sample case.
* After execute make cmd, call this function with paramse.
*
* Params
* exfp that represents execute file path that was made with make cmd.
* fin that represents input file path for each testcase.
* fout that represents output file path for each testcase.
*
*
* Return
* This function would return (status of the testcase, error)
* */

func executeSample(exfp string, fin string, fanswer string) (bool, error) {
	ex := exec.Command(exfp)
	stdin, err := ex.StdinPipe()
	if err != nil {
		log.Println("Failed to open stdin pipe.")
		return false, err
	}
	defer stdin.Close()
	input, err := readInput(fin)
	if err != nil {
		log.Println("Failed to read input.")
		return false, err
	}
	stdin.Write([]byte(input))

	err = ex.Run()
	if err != nil {
		log.Println("Failed to run the testcase. RE is caused.")
		return false, err
	}
	output, err := ex.Output()
	if err != nil {
		log.Println("Failed to run the testcase. RE is caused.")
		return false, err
	}

	status, err := compareWithAnswer(string(output), fanswer)
	if err != nil {
		log.Println("Failed to run the testcase. RE is caused.")
		return false, err
	}
	return status, nil
}

/* This function is responsible for reading input.*/
func readInput(fin string) (string, error) {
	return readIOFile(fin)
}

/* This function is responsible for reading output. */
func readOutput(fout string) (string, error) {
	return readIOFile(fout)
}

func readIOFile(fio string) (string, error) {
	file, err := os.Open(fio)
	if err != nil {
		log.Println("Failed to read io testcase file.")
		return "", err
	}
	defer file.Close()

	/*TODO: We have two types io file.
					 * 1. Word as token.
					 * 2. Line as token.
			     *
				   * They can be processed with bufio.Scanner
			     * But we should know type of the file.
		       *
		       * Switch tokenize function for type of token.
	         * 1. tokenizeAsLine (default)
	         * 2. tokenizeAsWord
					 * */
	content, err := tokenizeAsLine(file)
	if err != nil {
		return "", err
	}
	return content, nil
}

func tokenizeAsLine(f io.Reader) (string, error) {
	var b strings.Builder
	sc := bufio.NewScanner(f)
	return readToken(&b, sc)
}

func tokenizeAsWord(f io.Reader) (string, error) {
	var b strings.Builder
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)
	return readToken(&b, sc)
}

func readToken(b *strings.Builder, sc *bufio.Scanner) (string, error) {
	for sc.Scan() {
		b.WriteString(sc.Text())
		b.WriteString("\n")
	}
	if err := sc.Err(); err != nil {
		log.Printf("Failed to scan io file.")
		return "", err
	}
	return b.String(), nil
}

/* This function would be responsible for validating answer. */
func compareWithAnswer(output string, fanswer string) (bool, error) {
	ans, err := readIOFile(fanswer)
	stat, err := compareHelper(output, ans)
	return stat, err
}

/*Validation*/
func compareHelper(output string, answer string) (bool, error) {
	return true, nil
}
