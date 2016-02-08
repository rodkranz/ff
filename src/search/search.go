package search

import (
	"os"
	"path/filepath"
	"strings"
)

type Search struct {
	Text 			string
	File 			string
	Path 			string
	Reach 			int
	Regex 			bool
	CaseSensitive 	bool
}

var localStorage *Storage

func (s *Search) FindFiles(storage *Storage) {
	localStorage = storage
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

func (s *Search) HasText(file File) bool {
	return strings.Contains(file.File.Name(), "toastr")
}

func (s *Search) FindText(storage *Storage) {
	if len(s.Text) == 0 {
		return;
	}

	for i := len(storage.Files) -1; i >= 0; i-- {
		file := storage.Files[i]
		if !s.HasText(file) {
			storage.Remove(i)
		}
	}

}
