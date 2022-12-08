package filedir

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

/*
get directory that calls this function
*/
func GetInfrastructureDir() (string, error) {
	if os.Getenv("RUNTIME_ENV") == "local" {
		pc, file, line, ok := runtime.Caller(1) //"1" means that calls this function
		if !ok {
			return "", fmt.Errorf("Error occurs on GetFileDir(): pc:%v file:%s line:%d", pc, file, line)
		}
		return filepath.Dir(file), nil
	} else {
		return GetCurrentDir() + "/infrastructure", nil
	}
}
