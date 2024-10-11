package sorter

import (
	"testing"

	"github.com/Dimpal-Kalita/fsort/internal/flags"
)

func TestSortLines_StringSort(t *testing.T) {
	lines := []string{"banana", "apple", "orange", "mango", "date", "Cherry", "Date"}
	expected := []string{"Cherry", "Date", "apple", "banana", "date", "mango", "orange"}

	opts := &flags.Options{
		Numeric:    false,
		Reverse:    false,
		Unique:     false,
		IgnoreCase: false,
	}

	sortLines(lines, opts)

	if !compareSlices(lines, expected) {
		t.Errorf("String sort failed.\n Expected: %v\n Got: %v", expected, lines)
	}
}

func TestSortLines_CaseInsensitiveSort(t *testing.T) {
	lines := []string{"banana", "apple", "orange", "mango", "date", "Cherry", "Date", "Apple", "Apple"}
	expected := []string{"Apple", "Apple", "apple", "banana", "Cherry", "Date", "date", "mango", "orange"}
	opts := &flags.Options{
		Numeric:    false,
		Reverse:    false,
		Unique:     false,
		IgnoreCase: true,
	}

	sortLines(lines, opts)

	if !compareSlices(lines, expected) {
		t.Errorf("Case-Insensitive Sort failed.\n Expected: %v\n Got: %v", expected, lines)
	}
}

func TestSortLines_NumericSort(t *testing.T) {
	lines := []string{"10", "2", "1", "20", "5", "100", "1000", "1000"}
	expected := []string{"1", "2", "5", "10", "20", "100", "1000", "1000"}
	opts := &flags.Options{
		Numeric:    true,
		Reverse:    false,
		Unique:     false,
		IgnoreCase: false,
	}

	sortLines(lines, opts)

	if !compareSlices(lines, expected) {
		t.Errorf("Numeric Sort failed.\n Expected: %v\n Got: %v", expected, lines)
	}
}

func TestSortLines_UniqueSort(t *testing.T) {
	lines := []string{"apple", "apple", "orange", "mango", "date", "Cherry", "Date", "Apple", "Apple"}
	expected := []string{"Apple", "Cherry", "Date", "apple", "date", "mango", "orange"}
	opts := &flags.Options{
		Numeric:    false,
		Reverse:    false,
		Unique:     true,
		IgnoreCase: false,
	}

	sortLines(lines, opts)

	if !compareSlices(lines, expected) {
		t.Errorf("Unique Sort failed.\n Expected: %v\n Got: %v", expected, lines)
	}
}

func TestSortLines_ReverseSort(t *testing.T) {
	lines := []string{"apple", "apple", "orange", "mango", "date", "Cherry", "Date", "Apple", "Apple"}
	expected := []string{"orange", "mango", "date", "apple", "apple", "Apple", "Cherry", "Date"}
	opts := &flags.Options{
		Numeric:    false,
		Reverse:    true,
		Unique:     false,
		IgnoreCase: false,
	}

	sortLines(lines, opts)

	if !compareSlices(lines, expected) {
		t.Errorf("Reverse Sort failed.\n Expected: %v\n Got: %v", expected, lines)
	}
}

func compareSlices(lines, expected []string) bool {
	if len(lines) != len(expected) {
		return false
	}
	for i := range lines {
		if lines[i] != expected[i] {
			return false
		}
	}
	return true
}
