package main

import (
	"fmt"
)

type Position struct {
	FileName string // filename, if any
	Offset   int    // offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (byte count)
}

func (position *Position) IsValid() bool {
	return position.Line > 0
}

func (position Position) String() string {
	str := position.FileName
	if str == "" {
		str = "<input>"
	}
	if position.IsValid() {
		str += fmt.Sprintf(":%d", position.Line)
		if position.Column != 0 {
			str += fmt.Sprintf(":%d", position.Column)
		}
	}
	return str
}

type File struct {
	name  string
	lines []int
}

func NewFile(fileName string) *File {
	return &File{name: fileName, lines: []int{0}}
}

func (f *File) LineCount() int {
	return len(f.lines)
}

func (f *File) AddLine(offset int) {
	f.lines = append(f.lines, offset)
}

func (f *File) Line(p int) int {
	return f.Position(p).Line
}

func (f *File) Position(p int) (pos Position) {
	pos.Offset = p
	pos.FileName, pos.Line, pos.Column = f.unpack(p)
	return
}

func (f *File) unpack(offset int) (fileName string, line, column int) {
	fileName = f.name
	i, j := 0, len(f.lines)
	for i < j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
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
