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
	err = mergeChunks(tempFiles, output,opts)
	if err != nil {
		return fmt.Errorf("failed to merge chunks: %w", err)
	}
	return nil
}

func splitAndSortChunks(reader *bufio.Reader, opts *flags.Options) ([]string, error) {
	var tempFiles []string
	lines := make([]string, 0, opts.ChunkSize)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read line: %w", err)
		}
		if len(line) > 0 {
			lines = append(lines, strings.TrimSpace(line))
		}
		if len(lines) >= opts.ChunkSize || (err == io.EOF && len(lines) > 0) {
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


func Getcomparator(opts* flags.Options) func(string,string) bool{
	comparator := func(x,y string) bool { return x < y }

	if opts.Numeric {
		comparator = func(x,y string)bool {
			num1, err1 := strconv.Atoi(x)
			num2, err2 := strconv.Atoi(y)
			if err1 != nil && err2 != nil {
				return x<y
			}
			return num1 < num2
		}
	}
	if opts.Reverse {
		comparator = func(x,y string) bool { return x>y }
	}
	if opts.IgnoreCase {
		comparator = func(x,y string) bool { return strings.ToLower(x) < strings.ToLower(y) }
	}
	return comparator
}


func sortLines(lines []string, opts *flags.Options) {
	comparator:=Getcomparator(opts) 
	sort.Slice(lines, func(i, j int) bool { return comparator(lines[i], lines[j]) })	
	if opts.Unique {
		lines = utils.RemoveDuplicates(lines)
	}
	fmt.Println(lines)
}


func mergeChunks(tempFiles []string, output io.Writer,opts *flags.Options) error {
	var readers []*bufio.Reader
	for _, tempFile := range tempFiles {
		file, err := os.Open(tempFile)
		if err != nil {
			return fmt.Errorf("failed to open temp file: %w", err)
		}
		defer file.Close()
		readers = append(readers, bufio.NewReader(file))
	}
	pq := newPriorityQueue(readers,opts)
	heap.Init(pq)
	for !pq.Empty() {
		line, err := pq.Peek()
		if err != nil {
			return fmt.Errorf("failed to peek: %w", err)
		}
		_, err = fmt.Fprintln(output, line)
		if err != nil {
			return fmt.Errorf("failed to write line: %w", err)
		}
		_ = pq.Pop()
	}
	return nil
}

type priorityQueue struct {
	readers []*bufio.Reader
	opts *flags.Options
}

func newPriorityQueue(readers []*bufio.Reader,opts *flags.Options) *priorityQueue {
	return &priorityQueue{
		readers: readers,
		opts: opts,
	}
}

func (pq *priorityQueue) Len() int {
	return len(pq.readers)
}

func (pq *priorityQueue) Less(i, j int) bool {
	comparator:=Getcomparator(pq.opts)
	line1, err1 := pq.readers[i].ReadString('\n')
	line2, err2 := pq.readers[j].ReadString('\n')
	if err1 == io.EOF {
		return false
	}
	if err2 == io.EOF {
		return true
	}
	return comparator(strings.TrimSpace(line1), strings.TrimSpace(line2))
}

func (pq *priorityQueue) Swap(i, j int) {
	pq.readers[i], pq.readers[j] = pq.readers[j], pq.readers[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	pq.readers = append(pq.readers, x.(*bufio.Reader))
}

func (pq *priorityQueue) Pop() interface{} {
	old:=pq.readers
	n:=len(old)
	x:=old[n-1]
	pq.readers=old[0:n-1]
	return x
}

func (pq *priorityQueue) Peek() (string, error) {
	if pq.Len() == 0 {
		return "", io.EOF
	}
	line, err := pq.readers[0].ReadString('\n')
	if err == io.EOF {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("failed to read line: %w", err)
	}
	return strings.TrimSpace(line), nil
}

func (pq *priorityQueue) Empty() bool {
	return pq.Len() == 0
}

