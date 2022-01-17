package cfg

import (
	"os"
	"runtime"
)

// FileExist checks if the file exists.
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

// GetHomedir gets home directory.
func GetHomedir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return h
}

// GetSeparate gets the separation of a file path.
func GetSeparate() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	}
	return ""
}
