package sorter

import (
	"io"

	"github.com/Dimpal-Kalita/fsort/internal/flags"
)


type Sorter struct{
	opts *flags.Options
}

func NewSorter(opts *flags.Options) *Sorter{
	return &Sorter{opts: opts}
}

func (s *Sorter) Sort(inputFile string,output io.Writer) error{
	return ExternalSort(inputFile,output,s.opts)
}