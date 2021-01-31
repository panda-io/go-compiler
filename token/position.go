package token

import (
	"fmt"
)

type Position struct {
	file   *File
	offset int
}

func (p Position) String() string {
	path := p.file.Name
	line, column := p.file.location(p.offset)
	return fmt.Sprintf("%s:%d:%d", path, line, column)
}

func (p Position) Global() int {
	return p.file.Base + p.offset
}

type File struct {
	Name  string
	Size  int
	Base  int
	lines []int
}

func NewFile(filename string, size int) *File {
	return &File{
		Name:  filename,
		Size:  size,
		lines: []int{0},
	}
}

func (f *File) AddLine(offset int) {
	f.lines = append(f.lines, offset)
}

func (f *File) Lines() int {
	return len(f.lines)
}

func (f *File) Truncate(lines int) {
	f.lines = f.lines[:lines]
}

func (f *File) Position(offset int) *Position {
	return &Position{
		file:   f,
		offset: offset,
	}
}

func (f *File) Global(offset int) int {
	return f.Base + offset
}

func (f *File) location(offset int) (line, column int) {
	i, j := 0, len(f.lines)
	for i < j {
		h := i + (j-i)/2
		if f.lines[h] <= offset {
			i = h + 1
		} else {
			j = h
		}
	}
	i = i - 1
	if i >= 0 {
		line, column = i+1, offset-f.lines[i]+1
	}
	return
}

type FileSet struct {
	files []*File
	base  int
}

func (s *FileSet) AddFile(filename string, size int) *File {
	for _, f := range s.files {
		if f.Name == filename {
			panic(fmt.Sprintf("file %s already added \n", filename))
		}
	}
	f := NewFile(filename, size)
	f.Base = s.base
	s.base += size + 1
	s.files = append(s.files, f)
	return f
}

func (s *FileSet) UpdateFile(filename string, size int) {
	found := false
	for _, f := range s.files {
		if f.Name == filename {
			found = true
			f.Size = size
		}
	}
	if found {
		s.base = 0
		for _, f := range s.files {
			f.Base = s.base
			s.base += f.Size + 1
		}
	}
}

func (s *FileSet) File(position int) *File {
	for _, f := range s.files {
		if position <= f.Base+f.Size {
			return f
		}
	}
	return nil
}

func (s *FileSet) Position(position int) *Position {
	if f := s.File(position); f != nil {
		return f.Position(position - f.Base)
	}
	return nil
}
