package storage

import "github.com/rodkranz/ff/src/file"

type Storage struct {
	Files		[]file.File
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Add(newFile file.File) {
	s.Files = append(s.Files, newFile)
}

func (s *Storage) Remove(i int) {
	s.Files = s.Files[:i+copy(s.Files[i:], s.Files[i+1:])]
}