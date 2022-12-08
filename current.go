package filedir

import (
	"os"
	"runtime"
)

// __FILE__
func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

// __DIR__
func GetCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

func GetCurrentExecutable() string {
	path, err := os.Executable()
	if err != nil {
		return ""
	}
	return path
}
