package utils

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestWriteTempFile(t *testing.T) {
    lines := []string{"apple", "banana", "cherry"}
    tempFilePath, err := WriteTempFile(lines)
    if err != nil {
        t.Fatalf("WriteTempFile failed: %v", err)
    }
    defer os.Remove(tempFilePath) 

    content, err := ioutil.ReadFile(tempFilePath)
    if err != nil {
        t.Fatalf("Failed to read temp file: %v", err)
    }

   	expectedContent := "apple\nbanana\ncherry\n"
    if string(content) != expectedContent {
        t.Errorf("Expected file content:\n%s\nGot:\n%s", expectedContent, string(content))
    }
}

func TestRemoveDuplicates(t *testing.T) {
    testCases := []struct {
        input    []string
        expected []string
    }{
        {
            input:    []string{"apple", "apple", "banana", "banana", "cherry"},
            expected: []string{"apple", "banana", "cherry"},
        },
        {
            input:    []string{"apple", "banana", "cherry"},
            expected: []string{"apple", "banana", "cherry"},
        },
        {
            input:    []string{},
            expected: []string{},
        },
        {
            input:    []string{"onlyone"},
            expected: []string{"onlyone"},
        },
    }

    for _, tc := range testCases {
        result := RemoveDuplicates(tc.input)
		if len(result)==0 && len(tc.expected)==0{
			continue
		}
        if !reflect.DeepEqual(result, tc.expected) {
            t.Errorf("RemoveDuplicates(%v) expected %v, got %v", tc.input, tc.expected, result)
        }
    }
}
