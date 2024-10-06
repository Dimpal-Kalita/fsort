package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)


func WriteTempFile(lines []string) (string, error) {
	tmpfile, err := ioutil.TempFile("", "fsort_temp_*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}

	defer tmpfile.Close()

	writer:= bufio.NewWriter(tmpfile)
	for _, line := range lines {
		_, err := writer.WriteString(line+"\n")
		if err != nil {
			return "", fmt.Errorf("failed to write to temp file: %w", err)
		}
	}
	writer.Flush()
	return tmpfile.Name(), nil
}


func DeleteFiles(files []string){
	for _,file:= range files{
		os.Remove(file)
	}
}

func RemoveDuplicates(lines []string) []string {
	seen := make(map[string]struct{})
	var result []string
	for _, line := range lines {
		if _, ok := seen[line]; !ok {
			seen[line] = struct{}{}
			result = append(result, line)
		}
	}
	return result
}