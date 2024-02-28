package util

import "path/filepath"

func GetProjectRoot() string {
	return filepath.Dir(GetBackendProjectRoot())
}
