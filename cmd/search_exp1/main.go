package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"runtime/trace"
	"strings"
	"time"
)

const (
	SymbolDirectory = "[D]"
	SymbolFile      = "[F]"

	TracePath  = "./tmp/trace/"
)

var Prefix = "01"

type Config struct {
	Dir      string
	FileName string
	Text     string
	Avoid    []string

	output io.Writer
	ctx    context.Context
}

type Item struct {
	fileInfo os.FileInfo
	skip     bool
	lines    int
	path     string
	err      error
	found    map[int]string
}

func (i Item) GetType() string {
	if i.fileInfo.IsDir() {
		return SymbolDirectory
	}

	return SymbolFile
}

func main() {

	traceFileName := fmt.Sprintf("%s%s_%d.trace", TracePath, Prefix, time.Now().UnixNano())
	if traceFile, err := os.Create(traceFileName); err != nil {
		fmt.Println("Debug is not supported: ", err.Error())
	} else {
		if err := trace.Start(traceFile); err != nil {
			fmt.Println("Cannot close successfully the trace traceFile: ", err.Error())
			if err := traceFile.Close(); err != nil {
				fmt.Println("Cannot close successfully the trace traceFile: ", err.Error())
			}
		} else {
			defer func() {
				trace.Stop()
				if err := traceFile.Close(); err != nil {
					fmt.Println("Cannot close successfully the trace traceFile: ", err.Error())
				}
			}()
		}
	}

	cfg := Config{
		Avoid:  []string{".git", ".idea"},
		output: os.Stdout,
		ctx:    context.Background(),
	}

	cpus := flag.Int("cpus", runtime.NumCPU(), "number os virtual cores")
	flag.StringVar(&cfg.Dir, "dir", "", "directory where ff will search")
	flag.StringVar(&cfg.Text, "text", "", "write the text you need to find")
	flag.Parse()

	if cpus != nil {
		runtime.GOMAXPROCS(*cpus)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	items := make(chan Item)

	// Start filter files
	filtered := make(chan Item)
	{
		filterFiles := func() {
			for item := range items {
				if cfg.FileName == "" {
					filtered <- item
				}

				if strings.Contains(cfg.FileName, item.fileInfo.Name()) {
					filtered <- item
				}
			}
			close(filtered)
		}
		go filterFiles()
	}
	// End filter files

	// search_exp3 by text inside files
	ptln := make(chan Item)
	{
		searchText := func() {
			for item := range filtered {
				func() {
					file, err := os.OpenFile(item.path, os.O_RDONLY, 0666)
					if err != nil {
						item.err = err
						ptln <- item
						return
					}
					defer file.Close()

					// if f.config.CaseInsensitive {
					// 	f.config.Text = strings.ToLower(f.config.Text)
					// }

					foundLines := make(map[int]string)

					numLine := 0
					scanner := bufio.NewScanner(file)
					for scanner.Scan() {
						numLine++

						line := scanner.Text()

						if strings.Contains(line, cfg.Text) {
							foundLines[numLine] = line
						}

						// 	if f.config.CaseInsensitive {
						// 		line = strings.ToLower(line)
						// 	}
						//
						// 	f.searchByText(e, numLine, line)
						// 	f.searchByRegex(e, numLine, line)
					}
					item.lines = numLine
					item.found = foundLines

					ptln <- item
				}()
			}
			close(ptln)
		}
		go searchText()
	}
	// End filter files

	// Printer
	{
		printer := func() {
			w := cfg.output

			for item := range ptln {
				fmt.Fprintf(w, "%s %v\n", item.GetType(), item.path)
				for i, l := range item.found {
					fmt.Fprintf(w, "[%d] %v\n", i, l)
				}

				if item.err != nil {
					fmt.Fprintf(w, " ---> %s <--- \n", item.err)
				}
			}
		}
		go printer()
	}
	// End Printer

	//
	walkFn := func(path string, info os.FileInfo, err error) error {
		items <- Item{
			fileInfo: info,
			path:     path,
		}
		return nil
	}

	err := filepath.Walk(cfg.Dir, walkFn)
	close(items)

	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
