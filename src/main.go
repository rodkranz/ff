package main

import (
	"flag"
	"github.com/rodkranz/ff/src/search"
	"regexp"
	"github.com/rodkranz/ff/src/output"
	"github.com/rodkranz/ff/src/storage"
	"github.com/fatih/color"
	"runtime"
	"github.com/rodkranz/ff/src/update"
	"fmt"
	"os"
	"strings"
)

var (
	searching     search.Search
	showVersion   bool
	checkUpdate   bool
	CPUNum		  int
	Version       = "v1.1.2"
	UrlRepo       = "https://github.com/rodkranz/ff/releases"
)


func init() {
	flag.IntVar(&CPUNum, "cpu", 1, fmt.Sprintf("Number of CPU you have %d available", runtime.NumCPU()))
	flag.StringVar(&searching.Text, "t", "", "Text searching")
	flag.StringVar(&searching.File, "f", "", "Filter by file name")
	flag.StringVar(&searching.Path, "d", "./", "Directory searching")
	flag.IntVar(&searching.Reach, "a", 10, "Range around of the word")
	flag.BoolVar(&searching.CaseSensitive, "u", true, "Use case sensitive")
	flag.BoolVar(&color.NoColor, "-no-color", false, "Disable color output")
	flag.BoolVar(&showVersion, "-version", false, "Show the version")
	flag.BoolVar(&checkUpdate, "up", false, "Check update")
	flag.BoolVar(&searching.WithRegex, "regex",  false, "Search by this Regex")
	exclude     := *flag.String("-exclude-dir", ".bzr,CVS,.git,.hg,.svn", "Exclude dir from reader")

	flag.Parse()

	if len(exclude) > 0 {
		searching.Exclude = strings.Split(exclude, ",")
	}

	if searching.WithRegex {
		searching.Regex = regexp.MustCompile(searching.Text)
	}

	if CPUNum > runtime.NumCPU() {
		CPUNum = runtime.NumCPU()
	}

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
		output.ShowVersion(Version)
		os.Exit(0)
	}

	if checkUpdate {
		newVer, has := update.NewUpdate(UrlRepo, Version)
		output.ShowUpdate(newVer, has, UrlRepo)
		os.Exit(0)
	}

	// Find Files filtering by name
	storage := *storage.NewStorage()
	searching.SetStorage(&storage)
	searching.FindFiles()
	searching.SearchByText()

	output.ShowPretty(&searching)
	os.Exit(0)
}
