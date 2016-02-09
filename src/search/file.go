package search

import (
	"os"
)

type File struct {
	File		os.FileInfo
	Path		string
	Data 		[]byte
	Comment 	map[int]string
	NLine		int
}

func NewFile(path string, file os.FileInfo) *File {
	return &File{Path:path, File:file, Comment: make(map[int]string), NLine: 0}
}

func (f *File) WriteComment(line string) {
	f.Comment[f.NLine] = line
}