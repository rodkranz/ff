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
	"sync"
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
	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)

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
		Text: "WhomPointer",
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

	// Printer
	go PrintOutput(cfg, filtered)
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

	// time.Sleep(time.Second * 2)
}

func FilterNames(cfg Config, items chan Item) chan Item {
	filtered := make(chan Item)
	go func() {
		for item := range items {
			func() {
				ctx, tt := trace.NewTask(cfg.ctx, item.path)
				defer tt.End()

				// filtering by name
				// -------------------
				reg := trace.StartRegion(ctx, "ValidatingName")
				if cfg.FileName != "" && !strings.Contains(cfg.FileName, item.fileInfo.Name()) {
					reg.End()
					return
				}
				reg.End()
				// -------------------

				// search inside files
				// -------------------
				reg = trace.StartRegion(ctx, "OpenFile")
				file, err := os.OpenFile(item.path, os.O_RDONLY, 0666)
				if err != nil {
					reg.End()

					item.err = err
					filtered <- item
					return
				}
				defer file.Close()
				reg.End()

				// Variable to help us to collect information
				foundLines := make(map[int]string)
				numLine := 0
				scanner := bufio.NewScanner(file)

				// loop by lines looking for a text
				reg = trace.StartRegion(ctx, "ScanLines")
				for scanner.Scan() {
					numLine++

					line := scanner.Text()

					if strings.Contains(line, cfg.Text) {
						foundLines[numLine] = line
					}
				}
				reg.End()

				item.lines = numLine
				item.found = foundLines

				// -------------------
				filtered <- item
			}()
		}

		close(filtered)
	}()
	return filtered
}

func PrintOutput(cfg Config, outputs ...chan Item) {
	var wg sync.WaitGroup
	output := func(c chan Item) {
		ctx, tt := trace.NewTask(cfg.ctx, "Output")
		defer func() {
			tt.End()
			wg.Done()
		}()

		var w io.Writer
		b := &bytes.Buffer{}

		// control where ff will write the output.
		if cfg.output != nil {
			w = cfg.output
		} else {
			w = bufio.NewWriter(b)
		}

		reg := trace.StartRegion(ctx, "PrintOutput")
		for item := range c {
			fmt.Fprintf(w, "%s %v\n", item.GetType(), item.path)
			for i, l := range item.found {
				fmt.Fprintf(w, "[%d] %v\n", i, l)
			}

			if item.err != nil {
				fmt.Fprintf(w, " ---> %s <--- \n", item.err)
			}
		}
		reg.End()
	}

	wg.Add(len(outputs))
	for _, c := range outputs {
		go output(c)
	}

	wg.Wait()
}
