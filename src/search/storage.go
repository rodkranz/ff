package search

type Storage struct {
	Files		[]File
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Add(newFile File) {
	s.Files = append(s.Files, newFile)
}

func (s *Storage) Remove(i int) {
	s.Files = s.Files[:i+copy(s.Files[i:], s.Files[i+1:])]
}