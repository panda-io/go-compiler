package token

import (
	"fmt"
)

// Position present position of token
type Position struct {
	file   *File
	offset int
}

// String return the location of current position
func (p Position) String() string {
	path := p.file.Name
	line, column := p.file.location(p.offset)
	return fmt.Sprintf("%s:%d:%d", path, line, column)
}

// File store info of source file
type File struct {
	Name  string
	lines []int
	size  int
	base  int
}

// AddLine add line to soure file
func (f *File) AddLine(offset int) {
	f.lines = append(f.lines, offset)
}

// Position create position info
func (f *File) Position(offset int) *Position {
	return &Position{
		file:   f,
		offset: offset,
	}
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

// FileSet represents a set of source files.
type FileSet struct {
	files []*File
	base  int
	// TO-DO use it later for language server
	active *File
}

// AddFile add new file
func (s *FileSet) AddFile(filename string, size int) *File {
	for _, f := range s.files {
		if f.Name == filename {
			return f
		}
	}
	f := &File{
		Name:  filename,
		lines: []int{0},
		size:  size,
		base:  s.base,
	}
	s.base += size + 1
	s.files = append(s.files, f)
	return f
}

// File returns the file that contains the position (global position)
func (s *FileSet) File(position int) *File {
	for _, f := range s.files {
		if position <= f.base+f.size {
			return f
		}
	}
	return nil
}

// Position converts global position to a local positon in file
func (s *FileSet) Position(position int) *Position {
	if f := s.File(position); f != nil {
		return f.Position(position - f.base)
	}
	return nil
}
