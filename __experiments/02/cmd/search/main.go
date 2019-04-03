package main

import (
	"bufio"
	"bytes"
	"context"
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
)

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
	runtime.GOMAXPROCS(2)

	traceFileName := fmt.Sprintf("./tmp/trace/%v.trace", time.Now().UTC().Unix())

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
		// Dir: "./tmp/test-dir/",
		Dir: "/Users/rodkranz/Projects/Go/src/git.naspersclassifieds.com/olxeu/ecosystem/libs",
		// FileName: "Gopkg.toml",
		Text: "MinimumNArgs",
		Avoid: []string{
			".git", ".idea",
		},
		output: os.Stdout,
		ctx:    context.Background(),
	}

	// Start filter files
	items := make(chan Item)
	filtered := FilterNames(cfg, items)
	// End filter files

	// search by text inside files
	output := Output(cfg, filtered)
	// End filter files

	// Printer
	PrintOutput(cfg, output)
	// End Printer

	// function walk by all the files
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

	time.Sleep(time.Second * 2)
}

func Output(cfg Config, items chan Item) chan Item {
	output := make(chan Item)

	go func() {
		for item := range items {
			func() {
				ctx, tt := trace.NewTask(cfg.ctx, "SearchContent-"+item.path)
				defer tt.End()

				reg := trace.StartRegion(ctx, "OpenFile")
				file, err := os.OpenFile(item.path, os.O_RDONLY, 0666)
				if err != nil {
					item.err = err
					output <- item
					return
				}
				defer file.Close()
				reg.End()

				// if traceFile.config.CaseInsensitive {
				// 	traceFile.config.Text = strings.ToLower(traceFile.config.Text)
				// }

				foundLines := make(map[int]string)

				numLine := 0
				scanner := bufio.NewScanner(file)

				reg = trace.StartRegion(ctx, "ScanLines")
				for scanner.Scan() {
					numLine++

					line := scanner.Text()

					if strings.Contains(line, cfg.Text) {
						foundLines[numLine] = line
					}

					// 	if traceFile.config.CaseInsensitive {
					// 		line = strings.ToLower(line)
					// 	}
					//
					// 	traceFile.searchByText(e, numLine, line)
					// 	traceFile.searchByRegex(e, numLine, line)
				}
				reg.End()

				item.lines = numLine
				item.found = foundLines

				output <- item
			}()
		}
		close(output)
	}()

	return output
}

func FilterNames(cfg Config, items chan Item) chan Item {
	filtered := make(chan Item)
	go func() {
		for item := range items {
			func() {
				ctx, tt := trace.NewTask(cfg.ctx, "FilterByName"+item.path)
				defer tt.End()

				reg := trace.StartRegion(ctx, "OpenFile")
				defer reg.End()

				if cfg.FileName == "" {
					filtered <- item
				}

				if strings.Contains(cfg.FileName, item.fileInfo.Name()) {
					filtered <- item
				}
			}()
		}

		close(filtered)
	}()
	return filtered
}

func PrintOutput(cfg Config, output chan Item) {
	go func() {
		ctx, tt := trace.NewTask(cfg.ctx, "Output")
		defer tt.End()

		var w io.Writer
		b := &bytes.Buffer{}

		// control where ff will write the output.
		if cfg.output != nil {
			w = cfg.output
		} else {
			w = bufio.NewWriter(b)
		}

		reg := trace.StartRegion(ctx, "PrintOutput")
		for item := range output {
			fmt.Fprintf(w, "%s %v\n", item.GetType(), item.path)
			for i, l := range item.found {
				fmt.Fprintf(w, "[%d] %v\n", i, l)
			}

			if item.err != nil {
				fmt.Fprintf(w, " ---> %s <--- \n", item.err)
			}
		}
		reg.End()
	}()
}
