package search

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"github.com/rodkranz/ff/src/storage"
	"github.com/rodkranz/ff/src/file"
	"github.com/fatih/color"
	"sync"
)

type Search struct {
	Text          string
	File          string
	Path          string
	Reach         int
	WithRegex     bool
	Regex         *regexp.Regexp
	CaseSensitive bool
}

var (
	localStorage *storage.Storage
	ColorSearchText = color.New(color.FgRed).SprintFunc()
)


func (s *Search) SetStorage(storage *storage.Storage) {
	localStorage = storage
}

func (s *Search) GetStorage() *storage.Storage {
	return localStorage
}

func (s *Search) FindFiles() {
	filepath.Walk(s.Path, s.visitor)
}

func (s *Search) IsValidName(path string) bool {
	searchMe := s.File

	if !s.CaseSensitive {
		searchMe = strings.ToLower(searchMe)
		path	 = strings.ToLower(path)
	}

	return strings.Contains(path, searchMe)
}

func (s *Search) visitor(path string, f os.FileInfo, _ error) error {
	if s.IsValidName(path) {
		localStorage.Add(*file.NewFile(path, f))
	}

	return nil
}

func (s *Search) FindRegex(f *file.File, line string) {
	words := s.Regex.FindAllString(line, -1)
	if len(words) > 0 {
		for _, v := range words {
			f.WriteComment(s.Range(line, v))
		}
	}
}

func (s *Search) FindText(f *file.File, line string) {
	if strings.Contains(line, s.Text) {
		f.WriteComment(s.Range(line, s.Text))
	}
}

func (s *Search) HasText(f *file.File) bool {
	file, err := os.Open(f.Path)
	if err != nil {
		f.Comment[-1] = err.Error()
		return true
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.LineNum++
		line := scanner.Text()

		if s.WithRegex {
			s.FindRegex(f, line)
		}

		if !s.WithRegex {
			s.FindText(f, line)
		}
	}

	return len(f.Comment) > 0
}

func (s *Search) FindInGroup() {
	var ctrl sync.WaitGroup

	for i := 0; i < len(localStorage.Files); i++ {
		ctrl.Add(1)
		go func (file *file.File){
			file.Enabled = s.HasText(file)
			ctrl.Done()
		}(localStorage.GetById(i))
	}

	ctrl.Wait()
}

func (s *Search) SearchByText() {
	if len(s.Text) == 0 {
		return
	}

	s.FindInGroup()
}

func (s *Search) Range(line, text string) string {
	var ii, ie int
	index := strings.Index(line, text)

	word := line[index : index+len(text)]

	ii = index - s.Reach
	ie = len(text) + index + s.Reach

	if ii < 0 {
		ii = 0
	}
	if ie > len(line) {
		ie = len(line)
	}

	fontWord := line[ii:index]
	endWord := line[index+len(text) : ie]

	return fmt.Sprintf("%s%s%s", fontWord, ColorSearchText(word), endWord)
	return fmt.Sprintf("%s%s%s", fontWord, word, endWord)
}
