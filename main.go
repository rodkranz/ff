package main

import (
	"github.com/fatih/color"

	"flag"
	"strings"
	"path/filepath"
	"os"
	"bufio"
	"fmt"
	"errors"
	"regexp"
)

/*********************************************************************************/
type Search struct {
	File  string
	Text  string
    Path  string
	Reach int
}

func (s *Search) GetFile() (error, string) {
	if len(s.File) == 0 {
		return errors.New("File not defined!"), s.File
	}
	return nil, s.File;
}

func (s *Search) GetText() (error, string) {
	if len(s.Text) == 0 {
		return errors.New("Text not defined!"), s.Text
	}
	return nil, s.Text;
}

func (s *Search) GetPath() string {
	if len(s.Path) == 0 {
		return "./"
	}
	return s.Path;
}


func (s *Search) keepFile(path string) bool {
	if len(s.File) != 0 {
		if !strings.Contains(path, searching.File) {
			return false
		}
	}
	return true
}
func (s *Search) Range(line string, i int) string {
	var ii, ie int
	index := strings.Index(line, s.Text)

	word := line[index:index+len(s.Text)]

	ii = index - i ;
	ie = len(s.Text) + index + i;

	if ii < 0 { ii = 0 }
	if ie > len(line) { ie = len(line) }

	fontWord := line[ii:index]
	endWord  := line[index+len(s.Text):ie]

	red := color.New(color.FgRed).SprintFunc()
	return fmt.Sprintf("%s%s%s", fontWord, red(word), endWord)
}

func (s *Search) hasText(path string) (bool, map[int]string) {
	var lineNumber = make(map[int]string)
	var i int;

	if len(s.Text) != 0 {
		file, _ := os.Open(path)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			i++
			line := scanner.Text()

			if strings.Contains(line, s.Text) {
				lineNumber[i] = s.Range(line, searching.Reach)
			}
		}

		return len(lineNumber) != 0, lineNumber;
	}
	return true, lineNumber;
}
/*********************************************************************************/

/*********************************************************************************/
type Item struct {
	name 	string
	path 	string
	comment map[int]string
}
type Store struct {
	ListOfFiles []Item
}
func (s *Store) addFile(item Item) {
	s.ListOfFiles = append(s.ListOfFiles, item)
}
/*********************************************************************************/

var searching Search
var storage   Store = Store{}

func init() {

	// if ask about version
	re := regexp.MustCompile("^(-v|--version)$")
	for _, arg := range os.Args {
		if len(re.FindAllString(arg, -1)) != 0 {
			showVersion()
		}
	}

	path 		:= flag.String("p", 	 "./", 	"path string")
	text 		:= flag.String("t", 	 "", 	"the word that I have to looking for.")
	file 		:= flag.String("f", 	 "", 	"the file name that I have to looking for.")
	reach	    := flag.Int("r", 		 10, 	"range between start and end of the line")
	flagNoColor := flag.Bool("no-color", false, "Disable color output")

	flag.Parse()

	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}

	searching = Search{File: *file, Text: *text, Path: *path, Reach: *reach}
}

func findFilesInPath() {
	filepath.Walk(searching.Path, visitor)
}
// walk in each file
func visitor(path string, file os.FileInfo, _ error) error {
	// checkj if can keep the file
	if !searching.keepFile(path) {
		return nil
	}

	// check if has text that I am looking for.
	hasText, comments := searching.hasText(path);
	if !hasText {
		return nil;
	}

	storage.addFile(Item{file.Name(), path, comments})
	return nil
}

func showResult() {
	green 	:= color.New(color.FgGreen).SprintFunc()
	blue	:= color.New(color.FgBlue).SprintFunc()
	cyan 	:= color.New(color.FgCyan).SprintFunc()

	var nl = func () {
		color.Cyan("%s\n", strings.Repeat("-", 100))
	}

	nl()
	title := fmt.Sprintf("%s: %s", green("Path"),  cyan(searching.Path))
	if err, file := searching.GetFile(); err == nil {
		title = fmt.Sprintf("%s\n%s: %s", title, green("File"),  cyan(file))
	}
	if err, text := searching.GetText(); err == nil {
		title = fmt.Sprintf("%s\n%s: %s", title, green("Text"),  cyan(text))
	}

	fmt.Printf("%s\n", title)
	nl()
	for _, s := range storage.ListOfFiles {
		fmt.Printf("[%s] %s \n", green("File"), blue(s.path))
		for line, comment := range s.comment {
			fmt.Printf("\t[%s] %s\n", green(line), comment)
		}
		if len(s.comment) > 0 {
			nl()
		}
	}

}

func showVersion() {
	text := `--------------------------------------------------------------------------------------
 This program has writen by Rodrigo Lopes <dev.rodrigo.lopes@gmail.com>
 just to learn little more about GO language.
--------------------------------------------------------------------------------------
 Version : 1.0.0
 Language: Go Language
 Project : https://bitbucket.org/rkranz/gofindfileortext
 Contact : dev.rodrigo.lopes@gmail.com
 Linkedin: https://www.linkedin.com/in/rodrigo-lopes-76533724
--------------------------------------------------------------------------------------
`
	fmt.Println(text)
	os.Exit(0)
}

func main() {
	findFilesInPath()
	showResult()
}