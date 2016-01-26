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
	"time"
)

var myDebugMode = false

var red	= color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var blue= color.New(color.FgBlue).SprintFunc()
var cyan = color.New(color.FgCyan).SprintFunc()

/*********************************************************************************/
type Search struct {
	File  		string
	Text  		string
    Path  		string
	WithRegex 	bool
	Reach 		int
	Regexp 		*regexp.Regexp
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

func (s *Search) hasText(path string) (bool, map[int]string) {
	var lineNumber = make(map[int]string)
	var i int;

	if len(s.Text) == 0 {
		return true, lineNumber;
	}

	file, err := os.Open(path)
	if err != nil {
		lineNumber[-1] = fmt.Sprintf("%s", err.Error())
		return false, lineNumber;
	}
	defer file.Close()

	scanner  := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if s.WithRegex {
			words := searching.Regexp.FindAllString(line, -1)
			if len(words) > 0 {
				for _, v := range words {
					lineNumber[i] = Range(v, line, searching.Reach)
				}
			}
		} else {
			if strings.Contains(line, s.Text) {
				lineNumber[i] = Range(s.Text, line, searching.Reach)
			}
		}
	}

	return len(lineNumber) != 0, lineNumber;
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
	if len(re.FindAllString(strings.Join(os.Args[1:], " "), -1)) != 0 {
		showVersion()
	}

	path 		:= flag.String("p", 	 "./", 	"The directory path")
	text 		:= flag.String("t", 	 "", 	"what I am searching")
	regex 		:= flag.Bool("r", 	 	 false, "Search by regex")
	file 		:= flag.String("f", 	 "", 	"Filter by name of file or the file name")
	reach	    := flag.Int("a", 		 10, 	"Range around of the word that I found.")
	debugMode	:= flag.Bool("d", 	 	 false, "Show Debug Mode")
	flagNoColor := flag.Bool("no-color", false, "Disable color output")

	flag.Parse()

	myDebugMode = *debugMode

	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}

	searching = Search{
		File: *file,
		Text: *text,
		Path: *path,
		Reach: *reach,
		WithRegex: *regex,
	}

	if *regex {
		searching.Regexp = regexp.MustCompile(*text);
	}

	if len(os.Args) == 2 {
		if len(searching.Text) != 0 {
			return
		}

		if os.Args[1][0] != 45 {
			searching.Text = os.Args[1]
		}
	}
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
	var nl = func () {
		color.Cyan("\r%s\n", strings.Repeat("-", 100))
	}

	nl()
	title := fmt.Sprintf("%s : %s", green("Path"),  cyan(searching.Path))
	if err, file := searching.GetFile(); err == nil {
		title = fmt.Sprintf("%s\n%s : %s", title, green("File"),  cyan(file))
	}
	if err, text := searching.GetText(); err == nil {
		t := green("Text ")
		if searching.WithRegex {
			t = green("Regex")
		}
		title = fmt.Sprintf("%s\n%s: %s", title, t,  cyan(text))
	}

//	 To store the keys in slice in sorted order
	keys := make([]int, 0)
	for k := range storage.ListOfFiles {
		keys = append(keys, k)
	}

	fmt.Printf("%s\n", title)
	nl()
	for _, k := range keys {
		fmt.Printf("[%s] %s \n", green("File"), blue(storage.ListOfFiles[k].path))
		for line, comment := range storage.ListOfFiles[k].comment {
			if line == -1 {
				fmt.Printf("\t[%s] %s\n", green("ERROR"), red(comment))
				continue
			}
			fmt.Printf("\t[%s] %s\n", green(line), comment)
		}
		if len(storage.ListOfFiles[k].comment) > 0 {
			nl()
		}
	}
}

func Range(text, line string, i int) string {
	var ii, ie int
	index := strings.Index(line, text)

	word := line[index:index+len(text)]

	ii = index - i ;
	ie = len(text) + index + i;

	if ii < 0 { ii = 0 }
	if ie > len(line) { ie = len(line) }

	fontWord := line[ii:index]
	endWord  := line[index+len(text):ie]

	red := color.New(color.FgRed).SprintFunc()
	return fmt.Sprintf("%s%s%s", fontWord, red(word), endWord)
}

func showVersion() {
	text := `--------------------------------------------------------------------------------------
 This program has writen by Rodrigo Lopes <dev.rodrigo.lopes@gmail.com>
 just to learn little more about GO language.
--------------------------------------------------------------------------------------
 Version : 1.0.1
 Language: Go Language
 License : ISC
 Project : https://bitbucket.org/rkranz/gofindfileortext
 Contact : dev.rodrigo.lopes@gmail.com
 Linkedin: https://www.linkedin.com/in/rodrigo-lopes-76533724
--------------------------------------------------------------------------------------
`
	fmt.Println(text)
	os.Exit(0)
}

func main() {
	start := time.Now()
	fmt.Printf("\rSearching please wait...")

	findFilesInPath()
	showResult()

	if myDebugMode {
		fmt.Printf("final Execution took %s\n", green(time.Since(start)))
	}
}