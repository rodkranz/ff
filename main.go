package main

import (
	"fmt"
	"flag"
	"path/filepath"
	"os"
	"strings"
	"bufio"
	"sort"
)

var _path *string
var _text *string
var _file *string

func init() {
	_path = flag.String("path", "./", "path string")
	_text = flag.String("text", "", "the word that I have to looking for.")
	_file = flag.String("file", "", "the file name that I have to looking for.")

	flag.Parse()
	fmt.Printf("Directory: %s \nFile: %s \nText: %s\n-----------------\n", *_path, *_file, *_text)
}



func main() {
	filepath.Walk(*_path, visitor)
}

func visitor(path string, file os.FileInfo, _ error) error {

	// check if has filter to file name
	if len(*_file) != 0 {
		// if name has no in file path skip this file.
		if !strings.Contains(path, *_file) {
			return nil
		}
	}

	findTextInFile(path)

	return nil
}

func findTextInFile(path string) {
	file, _ := os.Open(path)
	defer file.Close()

	var lineNumber = make(map[int]string)
	var i int;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		line := scanner.Text()

		if strings.Contains(line, *_text) {
			lineNumber[i] = line
		}
	}

	if len(lineNumber) > 0 {
		fmt.Printf("[File: %s] \n", file.Name())
		// To store the keys in slice in sorted order
		var keys []int
		for k := range lineNumber {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		// To perform the opertion you want
		for _, k := range keys {
			fmt.Printf("\t%d: %s \n", k, lineNumber[k])
		}
	}
}