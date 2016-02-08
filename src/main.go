package main

import (
	"bitbucket.org/rkranz/gofind/src/search"
	"flag"
	"fmt"
)

var (
	searching 		search.Search
	showWithColor 	bool
	showVersion 	bool
)

func init() {
	flag.StringVar(&searching.Text,  		"t", 		  "", 			"Text searching")
	flag.StringVar(&searching.File,  		"f", 		  "", 			"Filter by file name")
	flag.StringVar(&searching.Path,  		"d",		  "./",			"Text searching")
	flag.IntVar(&searching.Reach, 			"a", 		  10,			"Range around of the word")
	flag.BoolVar(&searching.Regex, 			"r", 		  false, 		"Search by this Regex")
	flag.BoolVar(&searching.CaseSensitive, 	"u", 		  true,			"Use case sensitive")
	flag.BoolVar(&showWithColor, 			"no-color",   false,		"Disable color output")
	flag.BoolVar(&showVersion, 				"version",    false,		"Show the version")

	flag.Parse()
}


func main() {
	// Find Files filtering by name
	var storage search.Storage

	searching.FindFiles(&storage);
	searching.FindText(&storage);

	for i, s := range storage.Files {
		fmt.Printf("%v \t %v \n", i, s)
	}

	fmt.Printf("\n\n")
}