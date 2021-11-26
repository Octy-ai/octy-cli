package utils

import (
	"encoding/base64"
	"os"
	"strings"
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

// AfterStr : Get the substring after the specified string x.
func AfterStr(value string, x string) string {
	pos := strings.LastIndex(value, x)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(x)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

// BeforeStr : Get the substring before the specified string x.
func BeforeStr(value string, x string) string {

	pos := strings.Index(value, x)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

// InSlice : determines if the specified string is in the provided slice of strings
func InSlice(v string, s []string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

// ValueInIntStrMap: determines if the specified string is equal to any of the provided maps values
func ValueInIntStrMap(v string, m map[int]string) bool {
	for _, x := range m {
		if x == v {
			return true
		}
	}
	return false
}
