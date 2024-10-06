package flags

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Numeric    bool
	Reverse    bool
	Unique     bool
	IgnoreCase bool
	ChunkSize  int
}

func ParseFlags() Options {
	var opts Options
	pflag.BoolVarP(&opts.Numeric, "numeric", "n", false, "compare according to string numerical value")
	pflag.BoolVarP(&opts.Reverse, "reverse", "r", false, "reverse the result of comparisons")
	pflag.BoolVarP(&opts.Unique, "unique", "u", false, "output only the first of an equal run")
	pflag.BoolVarP(&opts.IgnoreCase, "ignore-case", "i", false, "fold lower case to upper case characters")
	pflag.IntVarP(&opts.ChunkSize, "chunk-size", "c", 100000, "use SIZE bytes per read")
	pflag.Parse()
	return opts
}

func IsFlag(arg string) bool {
	return len(arg) > 1 && arg[0] == '-'
}
