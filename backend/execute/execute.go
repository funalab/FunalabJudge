package execute

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
* Description
* This function is helper function.
* This function is responsible for executing one sample case.
* After execute make cmd, call this function with params.
*
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
	input, err := readFileToString(fin)
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

	answer, err := readFileToString(fanswer)
	if err != nil {
		log.Println("Failed to run the testcase. RE is caused.")
		return false, err
	}
	return compareWithAnswer(string(output), answer), nil
}

func readFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

/* This function would be responsible for validating answer. */
func compareWithAnswer(output string, answer string) bool {
	fixedOutput := strings.TrimRight(output, "\n")
	fixedAnswer := strings.TrimRight(answer, "\n")
	return fixedOutput == fixedAnswer
}
