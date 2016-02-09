package search
import (
	"github.com/fatih/color"

	"fmt"
	"strings"
)

const lineSize = 100


var ColorSearchText	= color.New(color.FgRed).SprintFunc()
var ColorNumbers 	= color.New(color.FgHiYellow).SprintFunc()
var ColorFileName	= color.New(color.FgBlue).SprintFunc()
var ColorTitles 	= color.New(color.FgGreen).SprintFunc()

func ShowPretty(search *Search)  {
	storage := search.GetStorage()

	ShowHeader(search)
	for _, file := range storage.Files {
		ShowFile(file.File.Name())
		for line, comment := range file.Comment {
			ShowComments(line, comment)
		}
		if len(file.Comment) > 0 {
			ShowLine()
		}
	}
	ShowFooter(search)
}

func ShowHeader(search *Search) {
	ShowLine()
	fmt.Printf("%s : %s\n", ColorTitles("Path"), search.Path)
	fmt.Printf("%s : %s\n", ColorTitles("File"), search.File)

	if len(search.Text) > 0 {
		searchBy := "Text "
		if search.WithRegex {
			searchBy = "Regex"
		}
		fmt.Printf("%s: %s\n", ColorTitles(searchBy), search.Text)
	}

	ShowLine()
}

func ShowFooter(search *Search) {
	if len(search.Text) == 0 {
		ShowLine()
	}
}

func ShowFile(fileName string) {
	fmt.Printf("[%s] %s \n", ColorTitles("File"), ColorFileName(fileName))
}

func ShowComments(line int, comment string) {
	fmt.Printf("\t[%s] \t %s \n", ColorNumbers(line), comment)
}

func ShowLine() {
	fmt.Printf("%s\n", strings.Repeat("-", lineSize))
}