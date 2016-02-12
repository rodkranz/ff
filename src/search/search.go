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
	"sort"
	"github.com/golang/glog"
)

type Search struct {
	Text          string
	File          string
	Path          string
	Exclude       []string
	Reach         int
	WithRegex     bool
	Regex         *regexp.Regexp
	CaseSensitive bool
}

var (
	localStorage *storage.Storage
	ColorSearchText = color.New(color.FgRed).SprintFunc()
)

func (s *Search) NeedToExclude(f os.FileInfo) bool {
	if f.Name() == "." {
		return false
	}
	if f.IsDir() {
		i := sort.SearchStrings(s.Exclude, f.Name())
		return (len(s.Exclude) != i)
	}
	return false
}

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
	if s.NeedToExclude(f) {
		return filepath.SkipDir
	}

	if s.IsValidName(path) {
		localStorage.Add(*file.NewFile(path, f))
	}

	return nil
}

func (s *Search) FindRegex(f *file.File, line string) {
	if s.Regex == nil {
		glog.Error("E.R. has error!")
	}

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
	totalFiles 	:= len(localStorage.Files)
	maxThread 	:= 4
	indexFile 	:= 0

	for indexFile < totalFiles {
		/** Make num of next threads **/
		numThreads := 0;
		for thread := 0; thread < maxThread; thread++ {
			numThreads = numThreads+1
			if (indexFile+numThreads) >= totalFiles {
				break
			}
		}

		semaphore := make(chan bool, numThreads)
		for thread := 0; thread < numThreads; thread++  {
			go func (i int) {
				file := localStorage.GetById(i)
				file.Enabled = s.HasText(file)
				semaphore <- true
			}(indexFile)
			indexFile = indexFile + 1
		}

		for thread := 0; thread < numThreads; thread++  {
			<-semaphore
		}
	}
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