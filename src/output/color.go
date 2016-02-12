package output

import (
	"github.com/fatih/color"

	"fmt"
	"strings"

	"github.com/rodkranz/ff/src/search"
	"github.com/rodkranz/ff/src/file"
)

var (
	LineSize        = 100
	ColorNumbers    = color.New(color.FgHiYellow).SprintFunc()
	ColorFileName   = color.New(color.FgBlue).SprintFunc()
	ColorTitles     = color.New(color.FgGreen).SprintFunc()
)

func ShowPretty(search *search.Search) {
	storage := search.GetStorage()

	ShowHeader(search)
	for _, file := range storage.Files {
		if !file.Enabled {
			continue
		}

		ShowFile(file)
		for _, line := range file.GetCommentSorted() {
			ShowComments(line, file.Comment[line])
		}

		if len(file.Comment) > 0 {
			ShowLine()
		}
	}
	ShowFooter(search)
}
func ShowHeader(search *search.Search) {
	ShowLine()
	fmt.Printf("%s  : %s\n", ColorTitles("Path"), search.Path)

	if len(search.File) != 0 {
		fmt.Printf("%s  : %s\n", ColorTitles("File"), search.File)
	}

	if len(search.Exclude) != 0 {
		fmt.Printf("%s: %v\n", ColorTitles("Ignore"), search.Exclude)
	}

	if len(search.Text) > 0 {
		searchBy := "Text "
		if search.WithRegex {
			searchBy = "Regex "
		}
		fmt.Printf("%s: %s\n", ColorTitles(searchBy), search.Text)
	}

	ShowLine()
}

func ShowFooter(search *search.Search) {
	if len(search.Text) == 0 {
		ShowLine()
	}
}

func ShowFile(file file.File) {

	fmt.Printf("[%s] %s\n", ColorTitles("File"), ColorFileName(file.Path))
}

func ShowComments(line int, comment string) {
	fmt.Printf("\t[%s] \t %s\n", ColorNumbers(line), comment)
}

func ShowLine() {
	fmt.Printf("%s\n", strings.Repeat("-", LineSize))
}

func ShowVersion(version string) {
	ShowLine();
	fmt.Printf("\tThis program has written by %s <%s>.\n", ColorTitles("Rodrigo Lopes"), ColorFileName("dev.rodrigo.lopes@gmail.com"))
    fmt.Printf("\tOnly for academic purposes\n")
	ShowLine();

	fmt.Printf("  Version : %s\n", ColorNumbers(version))
	fmt.Printf("  Language: %s\n", ColorTitles("GO Language"))
	fmt.Printf("  License : %s\n", ColorTitles("ISC"))
	fmt.Printf("  Project : %s\n", ColorFileName("https://github.com/rodkranz/ff"))
	fmt.Printf("  Contact : %s\n", ColorFileName("dev.rodrigo.lopes@gmail.com"))
	ShowLine()
}

func ShowUpdate(newVersion string, hasUpdate bool, UrlRepo string) {
	ShowLine()
	if hasUpdate {
		fmt.Printf("Exist a new update available\n")
		fmt.Printf("New version is %s.\n", ColorNumbers(newVersion))
		fmt.Printf("Please check the link \"%s\"\n", ColorTitles(UrlRepo))
	} else {
		fmt.Printf("You have the latest version\n")
		fmt.Printf("Current version is %s.\n", ColorNumbers(newVersion))
	}
	ShowLine()
}