package sorter

import (
	"bytes"
	"github.com/Dimpal-Kalita/fsort/internal/flags"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	mx = 1000000
)

func TestExternalSort_SmallFile(t *testing.T){
	inputContent:="banana\napple\norange\nmango\ndate\nCherry\nDate\n"
	expectedContent:="Cherry\nDate\napple\nbanana\ndate\nmango\norange\n"

	tmpInput,err:= os.CreateTemp("","input*.txt")
	if err!=nil{
		t.Fatalf("Error creating temp file: %v",err)
	}
	defer os.Remove(tmpInput.Name())

	tmpInput.WriteString(inputContent)
	tmpInput.Close()

	var output bytes.Buffer

	opts:=&flags.Options{
		Numeric:    false,
		Reverse:    false,
		Unique:     false,
		IgnoreCase: false,
		ChunkSize: 2,
	}
	err=ExternalSort(tmpInput.Name(),&output,opts)
	if err!=nil{
		t.Fatalf("(External Sort) Error sorting file: %v",err)
	}
	if output.String()!=expectedContent{
		t.Errorf("(External Sort) Expected:\n%s\nGot:\n%s",expectedContent,output.String())
	}
}

func TestExternalSort_SmallFile_Uniq(t *testing.T){
	inputContent:="banana\napple\norange\nmango\ndate\nCherry\nDate\napple\nbanana\ndate\nmango\norange\n"
	expectedContent:="Cherry\nDate\napple\nbanana\ndate\nmango\norange\n"

	tmpInput,err:= os.CreateTemp("","input*.txt")
	if err!=nil{
		t.Fatalf("Error creating temp file: %v",err)
	}
	defer os.Remove(tmpInput.Name())

	tmpInput.WriteString(inputContent)
	tmpInput.Close()

	var output bytes.Buffer

	opts:=&flags.Options{
		Numeric:    false,
		Reverse:    false,
		Unique:     true,
		IgnoreCase: false,
		ChunkSize: 2,
	}
	err=ExternalSort(tmpInput.Name(),&output,opts)
	if err!=nil{
		t.Fatalf("(External Sort) Error sorting file: %v",err)
	}
	if output.String()!=expectedContent{
		t.Errorf("(External Sort) Expected: %s\nGot: %s",expectedContent,output.String())
	}
}



func TestExternalSort_LargeFile(t *testing.T) {
	numLines := mx
	var inputBuilder strings.Builder
	expectedNumbers := make([]int, numLines)

	for i := 0; i < numLines; i++ {
		num := numLines - i
		expectedNumbers[i] = num
		inputBuilder.WriteString(strconv.Itoa(num) + "\n")
	}

	tmpInput, err := os.CreateTemp("", "input*.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tmpInput.Name())

	tmpInput.WriteString(inputBuilder.String())
	tmpInput.Close()

	var output bytes.Buffer

	opts := &flags.Options{
		Numeric:    true,
		Reverse:    false,
		Unique:     false,
		IgnoreCase: false,
		ChunkSize:  1000,
	}
	err = ExternalSort(tmpInput.Name(), &output, opts)
	if err != nil {
		t.Fatalf("Error sorting file: %v", err)
	}
	sortedLines := strings.Split(output.String(), "\n")
	sortedLines = sortedLines[:len(sortedLines)-1]
	sortedNumbers := make([]int, len(sortedLines))
	for i, line := range sortedLines {
		num, err := strconv.Atoi(line)
		if err != nil {
			t.Fatalf("Error parsing line %s: %v", line, err)
		}
		sortedNumbers[i] = num
	}

	for i := 1; i < len(sortedNumbers); i++ {
		if sortedNumbers[i] < sortedNumbers[i-1] {
			t.Errorf("Sorting failed at index %d: %d < %d", i, sortedNumbers[i], sortedNumbers[i-1])
		}
	}

}
