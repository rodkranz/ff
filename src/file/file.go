package file

import (
	"os"
	"sort"
)

type File struct {
	File    os.FileInfo
	Path    string
	Data    []byte
	Comment map[int]string
	LineNum int
}

func NewFile(path string, file os.FileInfo) *File {
	return &File{Path:path, File:file, Comment: make(map[int]string), LineNum: 0}
}

func (f *File) WriteComment(line string) {
	f.Comment[f.LineNum] = line
}

func (f *File) GetCommentSorted() []int {
	keys := make([]int, 0, len(f.Comment))
	for i, _ := range f.Comment {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	return keys
}