package check

import (
	"os"
	"runtime"
	"strings"
)

func CheckFile(path string) (os.FileInfo, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file, nil
	}
	if os.IsNotExist(err) {
		return nil, nil
	}
	return nil, err
}

func CheckWinPath(str string) string {
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(str, "\\", "/")
	}
	return str
}
