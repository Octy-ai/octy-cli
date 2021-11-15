package utils

import (
	"testing"
)

func TestDecodeB64String(t *testing.T) {

	want := "test token"
	inputStr := "dGVzdCB0b2tlbg=="

	decodedStr, err := DecodeB64String(inputStr)
	if decodedStr != want || err != nil {
		t.Errorf("DecodeB64String() did not decode b64 string correctly, Incorrect string returned, got: %v, want: %v.", decodedStr, want)
	}

}

func TestDirectoryExists(t *testing.T) {

	tables := []struct {
		path string
		res  bool
	}{
		{"/Users/bengoodenough/Documents", true},
		{"/Users/bengoodenough/Documents/does/not/exist", false},
	}

	for _, table := range tables {
		res := DirectoryExists(table.path)
		if res != table.res {
			t.Errorf("DirectoryExists() misidentified the existence of a directory, got: %v, want: %v.", res, table.res)
		}
	}
}
