package sorter

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Dimpal-Kalita/fsort/internal/flags"
	"github.com/Dimpal-Kalita/fsort/internal/utils"
)



func ExternalSort(inputFile string, output io.Writer, opts *flags.Options) error {
	var reader *bufio.Reader
	if inputFile == "" {
		reader = bufio.NewReader(os.Stdin)	
	} else {
		file, err := os.Open(inputFile)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	}

	tempFiles, err := splitAndSortChunks(reader, opts)
	if err != nil {
		return fmt.Errorf("failed to split and sort chunks: %w", err)
	}
	defer utils.DeleteFiles(tempFiles)
	err = mergeChunks(tempFiles, output)
	if err != nil {
		return fmt.Errorf("failed to merge chunks: %w", err)
	}
	return nil
}

func splitAndSortChunks(reader *bufio.Reader, opts *flags.Options) ([]string, error) {
	var tempFiles []string
	lines:= make([]string,0,opts.ChunkSize)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read line: %w", err)
		}
		if len(line)>0{
			lines = append(lines, strings.TrimSpace(line))
		}	
		if len(lines)>= opts.ChunkSize || (err == io.EOF && len(lines)>0){
			sortLines(lines, opts)
			tempFile, err := utils.WriteTempFile(lines)
			if err != nil {
				return nil, fmt.Errorf("failed to write temp file: %w", err)
			}
			tempFiles = append(tempFiles, tempFile)
			lines = make([]string, 0, opts.ChunkSize)	
		}
		if err == io.EOF {
			break
		}
	}
	return tempFiles, nil
}

func sortLines(lines []string, opts *flags.Options) {
	comparator := func(i, j int) bool { return lines[i] < lines[j] }

	if opts.Numeric {
		comparator = func(i, j int) bool {
			num1, err1 := strconv.Atoi(lines[i])
			num2, err2 := strconv.Atoi(lines[j])
			if err1 != nil && err2 != nil {
				return lines[i] < lines[j]
			}
			return num1 < num2
		}
	}
	if opts.Reverse {
		comparator = func(i, j int) bool { return lines[i]>lines[j] }
	}
	if opts.IgnoreCase {
		comparator = func(i, j int) bool { return strings.ToLower(lines[i])<strings.ToLower(lines[j]) }
	}
	sort.Slice(lines, func(i, j int) bool { return comparator(i, j) })

	if opts.Unique {
		lines = utils.RemoveDuplicates(lines)
	}
}

func mergeChunks(tempFiles []string, output io.Writer) error {
	var readers []*bufio.Reader
	for _, file := range tempFiles {
		f, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open temp file: %w", err)
		}
		defer f.Close()
		readers = append(readers, bufio.NewReader(f))
	}

	var h minHeap
	for i, reader := range readers {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read line from temp file: %w", err)
		}
		if len(line) > 0 {
			heap.Push(&h, heapItem{line: strings.TrimSpace(line), fileIdx: i})
		}
	}

	for h.Len() > 0 {
		line := heap.Pop(&h).(string)
		fmt.Fprintln(output, line)

		reader := readers[h[0].fileIdx]
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read line from temp file: %w", err)
		}
		if len(line) > 0 {
			heap.Push(&h, heapItem{line: strings.TrimSpace(line), fileIdx: h[0].fileIdx})
		}
	}
	return nil
}

type heapItem struct {
	line   string
	fileIdx int
}

type minHeap []heapItem

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool {
	return h[i].line < h[j].line
}
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(heapItem))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item.line
}
