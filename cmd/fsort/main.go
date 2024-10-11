package main

import (
	"flag"
	"log"
	"os"

	"github.com/Dimpal-Kalita/fsort/internal/flags"
	"github.com/Dimpal-Kalita/fsort/internal/sorter"
)


func main(){
	opts:= flags.ParseFlags()

	inputFile:=""

	if flag.NArg() > 0 {
		inputFile = flag.Arg(0)
	}

	s:= sorter.NewSorter(&opts)	

	err:= s.Sort(inputFile, os.Stdout)
	if err!=nil{
		log.Fatalf("Error while sorting: %v", err)
	}
}