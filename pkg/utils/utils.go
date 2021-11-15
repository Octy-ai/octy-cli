package utils

import (
	"encoding/base64"
	"os"
)

// DecodeB64String : returns decoded base64 encoded string
func DecodeB64String(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// DirectoryExists : returns whether the given file or directory exists
func DirectoryExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
