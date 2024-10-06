package flags

import (
	"os"
	"reflect"
	"testing"

	"github.com/spf13/pflag"
)

func TestParseFlags(t *testing.T) {
    tests:= [] struct {
        name string
        args []string
        want Options
        wantInput string
        wantErr bool
    }{
        {
            name: "no flags",
            args: []string{"fsort", "file.txt"},
            want: Options{
                Numeric: false,
                Reverse: false,
                Unique: false,
                IgnoreCase: false,
                ChunkSize: 100000,
            },
        },
        {
            name: "numeric and reverse flags",
            args: []string{"fsort", "-n", "-r", "file.txt"},
            want: Options{
                Numeric: true,
                Reverse: true,
                Unique: false,
                IgnoreCase: false,
                ChunkSize: 100000,
            },
        },
        {
            name: "all flags",
            args: []string{"fsort", "-n", "-r", "-u", "-i", "-c", "100", "file.txt"},
            want: Options{
                Numeric: true,
                Reverse: true,
                Unique: true,
                IgnoreCase: true,
                ChunkSize: 100,
            },
        },
      {
            name: "no file provided read from stdin",
            args: []string{"fsort","-u"},
            want: Options{
                Numeric: false,
                Reverse: false,
                Unique: true,
                IgnoreCase: false,
                ChunkSize: 100000,
            },    
      },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            originalArgs := os.Args
            defer func() { os.Args = originalArgs }()
            os.Args = tt.args
            pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
            got := ParseFlags()
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ParseFlags() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestIsFlag(t *testing.T) {
    tests := []struct {
        name string
        arg  string
        want bool
    }{
        {
            name: "flag",
            arg:  "-n",
            want: true,
        },
        {
            name: "not a flag",
            arg:  "file.txt",
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := IsFlag(tt.arg)
            if got != tt.want {
                t.Errorf("IsFlag() = %v, want %v", got, tt.want)
            }
        })
    }
}
