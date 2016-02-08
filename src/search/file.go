package search

import "os"

type File struct {
	File		os.FileInfo
	Path		string
	Data 		[]byte
	Comment 	map[int]string
}

func NewFile(path string, file os.FileInfo) *File {
	return &File{Path:path, File:file}
}


func (f *File) IsValid(s Search) {



}