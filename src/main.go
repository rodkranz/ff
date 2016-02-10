package main

import (
	"flag"
	"github.com/rodkranz/ff/src/search"
	"regexp"
	"github.com/rodkranz/ff/src/output"
	"github.com/rodkranz/ff/src/storage"
	"github.com/fatih/color"
	"runtime"
)

var (
	searching     search.Search
	showVersion   bool
	CPUNum		  int
)

func init() {
	flag.IntVar(&CPUNum, "cpu", runtime.NumCPU(), "Number of CPU")
	flag.StringVar(&searching.Text, "t", "", "Text searching")
	flag.StringVar(&searching.File, "f", "", "Filter by file name")
	flag.StringVar(&searching.Path, "d", "./", "Text searching")
	flag.IntVar(&searching.Reach, "a", 10, "Range around of the word")
	flag.BoolVar(&searching.WithRegex, "r", false, "Search by this Regex")
	flag.BoolVar(&searching.CaseSensitive, "u", true, "Use case sensitive")
	flag.BoolVar(&color.NoColor, "no-color", false, "Disable color output")
	flag.BoolVar(&showVersion, "version", false, "Show the version")

	if searching.WithRegex {
		searching.Regex = regexp.MustCompile(searching.Text)
	}

	flag.Parse()

	// Check if has parameters
	if flag.NArg() == 1 {
		if par := flag.Arg(0); par[0] != 45 {
			searching.Text = flag.Arg(0)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(CPUNum)

	if showVersion {
		output.ShowVersion()
	}

	// Find Files filtering by name
	storage := *storage.NewStorage()
	searching.SetStorage(&storage)
	searching.FindFiles()
	searching.SearchByText()

	output.ShowPretty(&searching)
}
