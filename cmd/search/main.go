package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime/trace"
	"strings"
	"time"
)

const (
	SymbolDirectory = "[D]"
	SymbolFile      = "[F]"
)

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

	cfg := struct {
		Dir      string
		FileName string
		Text     string
		Avoid    []string
	}{

		Dir: "./tmp/test-dir/",
		// FileName: "Gopkg.toml",
		Text: "MinimumNArgs",
		Avoid: []string{
			".git", ".idea",
		},
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

	// search by text inside files
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

					// if traceFile.config.CaseInsensitive {
					// 	traceFile.config.Text = strings.ToLower(traceFile.config.Text)
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

						// 	if traceFile.config.CaseInsensitive {
						// 		line = strings.ToLower(line)
						// 	}
						//
						// 	traceFile.searchByText(e, numLine, line)
						// 	traceFile.searchByRegex(e, numLine, line)
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
			b := &bytes.Buffer{}
			w := bufio.NewWriter(b)

			for item := range ptln {
				fmt.Fprintf(w, "%s %v\n", item.GetType(), item.path)
				for i, l := range item.found {
					fmt.Fprintf(w, "[%d] %v\n", i, l)
				}

				if item.err != nil {
					fmt.Fprintf(w, " ---> %s <--- \n", item.err)
				}
			}

			outputFileName := fmt.Sprintf("./tmp/output/%d.output", time.Now().UTC().Unix())

			traceFile, err := os.Create(outputFileName);
			if err != nil {
				print("Debug is not supported: ", err.Error())
				return
			}

			if err := w.Flush(); err != nil {
				print("Writter Flush: ", err.Error())
				panic(err)
			}

			if _, err := traceFile.Write(b.Bytes()); err != nil {
				print("traceFile Write: ", err.Error())
				panic(err)
			}

			if err := traceFile.Close(); err != nil {
				print("traceFile Close: ", err.Error())
				panic(err)
			}

			// println(b.String())
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

	time.Sleep(time.Second * 2)
	os.Exit(0)
}
