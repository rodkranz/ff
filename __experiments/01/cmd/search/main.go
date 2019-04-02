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
	trace.Start(os.Stdout)
	defer trace.Stop()

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

			outputFileName := fmt.Sprintf("./tmp/output/%d.log", time.Now().UTC().Unix())
			f, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				panic(err)
			}
			if err := f.Close(); err != nil {
				panic(err)
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
