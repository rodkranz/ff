package main

import (
	"flag"
	"fmt"
	"github.com/rodkranz/ff/src/search"
	"regexp"
	"github.com/rodkranz/ff/src/output"
	"github.com/rodkranz/ff/src/storage"
	"github.com/fatih/color"
)

var (
	searching     search.Search
	showVersion   bool
)

func init() {
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
	if showVersion {
		output.ShowVersion()
	}

	// Find Files filtering by name
	var storage storage.Storage

	searching.SetStorage(&storage)
	searching.FindFiles()
	searching.SearchByText()

	output.ShowPretty(&searching)

	fmt.Printf("\n\n")
}
