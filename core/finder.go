package ff

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sort"
	"bufio"
	"errors"
)

/**
 * Configuration of searchers
 */
type Config struct {
	Text            string
	Regex           *regexp.Regexp

	FilterFile      string
	Directory       string

	ExcludeFiles    []string
	CaseInsensitive bool

	Reach           int
}

/**
 * Finder struct with owner properties
 */
type Finder struct {
	config Config       `json: "config"`
	list   []Element    `json: "list"`
	err    error		`json: "error"`
}

/**
 * Return a new instance of Finder
 */
func NewFinder(conf Config) *Finder {
	return &Finder{config: conf }
}

/**
 * Check if need to exclude this file/folder or not.
 */
func (f *Finder) NeedToExclude(file os.FileInfo) bool {
	if file.Name() == "." {
		return false
	}

	if file.IsDir() {
		i := sort.SearchStrings(f.config.ExcludeFiles, file.Name())
		return (len(f.config.ExcludeFiles) != i)
	}

	return false
}

/**
 * Check if file name is valid
 */
func (f *Finder) IsValidName(path string) bool {
	return strings.Contains(path, f.config.FilterFile)
}

/**
 * Search Files recursively and filter name that we need to exclude
 * and invalid name.
 */
func (f *Finder) FindFiles(searchDir string) []Element {
	filepath.Walk(searchDir, func(path string, file os.FileInfo, _ error) error {
		if f.NeedToExclude(file) {
			return filepath.SkipDir
		}

		if !f.IsValidName(path) {
			return nil
		}

		// if is a file valid add in list
		f.list = append(f.list, *NewElement(path, file, len(f.list)))

		return nil
	})

	return f.list
}

/**
 * Find word in line By string
 */
func (f *Finder) searchByText(elem *Element, numLine int, line string) {
	if f.config.Regex != nil {
		return
	}

	if strings.Contains(line, f.config.Text) {
		// Check if has the text in line
		elem.WriteComment(numLine, f.config.Text, line)    // write comment
	}
}

/**
 * Find word in line By regex
 */
func (f *Finder) searchByRegex(elem *Element, numLine int, line string) {
	if f.config.Regex == nil {
		return
	}

	words := f.config.Regex.FindAllString(line, -1)
	if len(words) > 0 {
		for _, v := range words {
			elem.WriteComment(numLine, v, line)
		}
	}
}

/**
 * Read line by line and send lines to check
 */
func (f *Finder) readAndFind(e *Element) {
	e.Enabled = false

	file, err := os.Open(e.Path)
	if err != nil {
		e.Error = err
		return
	}
	defer file.Close()

	if f.config.CaseInsensitive {
		f.config.Text = strings.ToLower(f.config.Text)
	}

	numLine := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numLine++

		if f.config.CaseInsensitive {
			line = strings.ToLower(line)
		}

		f.searchByText(e, numLine, line)
		f.searchByRegex(e, numLine, line)
	}
	e.NumLines = numLine
}

/**
 * Find Text in line
 */
func (f *Finder) FindText(text string) []Element {
	if len(f.config.Text) == 0 && f.config.Regex == nil {
		return f.list
	}

	f.config.Text = text
	for i, element := range f.list {
		f.readAndFind(&element)
		f.list[i] = element
	}

	return f.list
}

func (f *Finder) GetResult() (List []Element, err error) {
	if f.err != nil {
		return List, f.err
	}

	if len(f.list) == 0 {
		err = errors.New("There is no result to show!")
	}

	return f.list, err
}