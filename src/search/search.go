package search

import (
	"os"
	"path/filepath"
	"strings"
	"bufio"
	"regexp"
	"fmt"
)

type Search struct {
	Text 			string
	File 			string
	Path 			string
	Reach 			int
	WithRegex		bool
	Regex			*regexp.Regexp
	CaseSensitive 	bool
}

var localStorage *Storage

func (s *Search) SetStorage(storage *Storage) {
	localStorage = storage
}

func (s *Search) GetStorage() *Storage {
	return localStorage
}

func (s *Search) FindFiles() {
	filepath.Walk(s.Path, s.visitor)
}

func (s *Search) IsValidName(path string) bool {
	searchMe := s.File

	if !s.CaseSensitive {
		searchMe = strings.ToLower(searchMe)
		path 	 = strings.ToLower(path)
	}

	return strings.Contains(path, searchMe)
}

func (s *Search) visitor(path string, file os.FileInfo, _ error) error {
	if s.IsValidName(path) {
		localStorage.Add(*NewFile(path, file))
	}

	return nil
}


func (s *Search) FindRegex(f *File, line string) {
	words := s.Regex.FindAllString(line, -1)
	if len(words) > 0 {
		for _, v := range words {
			f.WriteComment(s.Range(line, v))
		}
	}
}

func (s *Search) FindText(f *File, line string) {
	if strings.Contains(line, s.Text) {
		f.WriteComment(s.Range(line, s.Text))
	}
}

func (s *Search) HasText(f *File) bool {
	file, err := os.Open(f.Path);
	if err != nil {
		f.Comment[-1] = err.Error()
		return true
	}
	defer file.Close()

	scanner  := bufio.NewScanner(file)
	for scanner.Scan() {
		f.NLine++
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

func (s *Search) SearchByText() {
	if len(s.Text) == 0 {
		return;
	}

	for i := len(localStorage.Files) -1; i >= 0; i-- {
		file := localStorage.Files[i]
		if !s.HasText(&file) {
			localStorage.Remove(i)
		}
	}
}

func (s *Search)  Range(line, text string) string {
	var ii, ie int
	index := strings.Index(line, text)

	word := line[index:index+len(text)]

	ii = index - s.Reach
	ie = len(text) + index + s.Reach

	if ii < 0 { ii = 0 }
	if ie > len(line) { ie = len(line) }

	fontWord := line[ii:index]
	endWord  := line[index+len(text):ie]

	return fmt.Sprintf("%s%s%s", fontWord, ColorSearchText(word), endWord)
}