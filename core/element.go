package ff

import (
	"os"
	"sort"
)

type Element struct {
	File     os.FileInfo       `json: "file"`
	Path     string            `json: "path"`
	Data     []byte            `json: "data"`
	Line     map[int]string    `json: "line"`
	Text     map[int]string    `json: "text "`
	NumLines int               `json: "numlines"`
	Index    int               `json: "index"`
	Error    error             `json: "error"`
	Enabled  bool              `json: "enabled"`
}

/**
 * Create a new instance of Element
 */
func NewElement(path string, file os.FileInfo, index int) *Element {
	return &Element{Path:path,
		File:       file,
		NumLines    : 0,
		Index:      index,
		Text:       make(map[int]string),
		Line:       make(map[int]string),
		Enabled:    true,
	}
}

/**
 * Write a comment of element
 */
func (e *Element) WriteComment(numLine int, text, line string) {
	e.Enabled		= true
	e.Line[numLine] = line
	e.Text[numLine] = text
}

/**
 * Get comments of element
 */
func (e *Element) GetComment() []int {
	keys := make([]int, 0, len(e.Text))
	for i, _ := range e.Text {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	return keys
}