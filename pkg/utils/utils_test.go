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

func TestAfterStr(t *testing.T) {
	tables := []struct {
		value string
		x     string
		res   string
	}{
		{"hello**there", "**", "there"},
		{"test&**^string", "**^", "string"},
	}
	for _, table := range tables {
		res := AfterStr(table.value, table.x)
		if res != table.res {
			t.Errorf("AfterStr() failed to split string on specififed string, got: %v, want: %v.", res, table.res)
		}
	}
}

func TestBeforeStr(t *testing.T) {
	tables := []struct {
		value string
		x     string
		res   string
	}{
		{"hello**there", "**", "hello"},
		{"test&**^string", "&**^", "test"},
	}
	for _, table := range tables {
		res := BeforeStr(table.value, table.x)
		if res != table.res {
			t.Errorf("BeforeStr() failed to split string on specififed string, got: %v, want: %v.", res, table.res)
		}
	}
}

func TestInSlice(t *testing.T) {
	tables := []struct {
		value string
		array []string
		res   bool
	}{
		{"hello", []string{"hello", "there"}, true},
		{"test", []string{"hello", "there"}, false},
	}

	for _, table := range tables {
		res := InSlice(table.value, table.array)
		if res != table.res {
			t.Errorf("In() failed to detect if specified string was in array, got: %v, want: %v.", res, table.res)
		}
	}

}

func TestValueInIntStrMap(t *testing.T) {
	tables := []struct {
		value     string
		intStrMap map[int]string
		res       bool
	}{
		{"hello", map[int]string{1: "hello", 2: "there"}, true},
		{"test", map[int]string{1: "hello", 2: "there"}, false},
	}

	for _, table := range tables {
		res := ValueInIntStrMap(table.value, table.intStrMap)
		if res != table.res {
			t.Errorf("ValueInIntStrMap() failed to detect if specified string was in map values, got: %v, want: %v.", res, table.res)
		}
	}

}
