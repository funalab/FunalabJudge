package util

import "os"

/*
* This function should be used when we get file path from datbase.
* Path stored in db should be represented as fragment.
* Fragment means "Absolutepath - projectRootPath"
* */

func OpenFileFromDB(pfp string) (*os.File, error) {
	absPath := ParseIntoAbsolutePath(pfp)
	f, err := os.Open(absPath)
	return f, err
}

func ParseIntoAbsolutePath(pathFragment string) string {
	return GetProjectRoot() + pathFragment
}
