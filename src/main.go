package main

import (
	"flag"
	"regexp"
	"runtime"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodkranz/ff/src/search"
	"github.com/rodkranz/ff/src/output"
	"github.com/rodkranz/ff/src/storage"
	"github.com/rodkranz/ff/src/update"
)

var (
	searching     search.Search
	showVersion   bool
	checkUpdate   bool
	CPUNum		  int
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
	// allow to use N of CPUs available.
	runtime.GOMAXPROCS(CPUNum)

	if showVersion || checkUpdate {
		// create a instance of update
		up := update.NewUpdate("v1.1.2",
			`https://github.com/rodkranz/ff/releases`,
			`href=\"/rodkranz/ff/releases/tag/[ANY]"`)

		// show the current vertion of program
		if showVersion {
			output.ShowVersion(up.Version)
		}

		// Check if it has update.
		if checkUpdate {
			newVer, has := up.Check()
			output.ShowUpdate(newVer, has, up.Url)
		}

		// finalize the application correctly
		os.Exit(0)
	}

	// Show text 'searching'
	output.ShowWaitSearching()

	// create a virtual storage
	storage := *storage.NewStorage()
	searching.SetStorage(&storage)

	// search the files/text
	searching.FindFiles()
	searching.SearchByText()

	//show result pretty
	output.ShowPretty(&searching)

	// finalize the application correctly
	os.Exit(0)
}
