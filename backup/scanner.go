package main

/**
Meta programming

// stmt
@emit string

// include
@include "<vector>"

// doc and auto generation related
@doc string
@serializer(name:"name", omit_empty, index:1)
@meta(name:"name", data:"data") // runtime meta

// cpp related // later
@ref string
@call "$.insert($.begin() + ${position}, ${val})"
@return "$.back()"
@macro

**/

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const (
	bom = 0xFEFF // byte order mark, only permitted as very first character
	eof = -1
)

type ErrorHandler func(position Position, msg string)

type Scanner struct {
	file *File
	src  []byte

	errorHandler ErrorHandler
	errorCount   int // total errors

	flags    map[string]bool // flags for condition compiler
	flagOpen bool            // if #if is true

	char       rune
	offset     int
	readOffset int
}

func NewScanner(file *File, src []byte, errorHandler ErrorHandler, scanComment bool, flags []string) *Scanner {
	s := &Scanner{}

	s.file = file
	s.src = src
	s.errorHandler = errorHandler
	s.flags = make(map[string]bool)
	s.char = ' '

	s.next()
	if s.char == bom {
		s.next()
	}
	for _, flag := range flags {
		s.flags[flag] = true
	}

	return s
}

func (s *Scanner) next() {
	if s.readOffset < len(s.src) {
		s.offset = s.readOffset
		if s.char == '\n' {
			s.file.AddLine(s.offset)
		}
		r, w := rune(s.src[s.readOffset]), 1
		switch {
		case r == 0:
			s.error(s.offset, "illegal character NUL")
		case r >= utf8.RuneSelf:
			// not ASCII
			r, w = utf8.DecodeRune(s.src[s.readOffset:])
			if r == utf8.RuneError && w == 1 {
				s.error(s.offset, "illegal UTF-8 encoding")
			} else if r == bom && s.offset > 0 {
				s.error(s.offset, "illegal byte order mark")
			}
		}
		s.readOffset += w
		s.char = r
	} else {
		s.offset = len(s.src)
		if s.char == '\n' {
			s.file.AddLine(s.offset)
		}
		s.char = eof
	}
}

func (s *Scanner) peek() byte {
	if s.readOffset < len(s.src) {
		return s.src[s.readOffset]
	}
	return 0
}

func (s *Scanner) error(offset int, msg string) {
	if s.errorHandler != nil {
		s.errorHandler(s.file.Position(offset), msg)
	}
	s.errorCount++
}

func (s *Scanner) scanComment() string {
	// initial '/' already consumed
	offset := s.offset - 1 // position of initial '/'

	if s.char == '/' {
		//-style comment
		s.next()
		for s.char != '\n' && s.char >= 0 {
			s.next()
		}
	} else {
		/*-style comment */
		terminated := false
		s.next()
		for s.char >= 0 {
			char := s.char
			s.next()
			if char == '*' && s.char == '/' {
				s.next()
				terminated = true
				break
			}
		}
		if !terminated {
			s.error(offset, "comment not terminated")
		}
	}
	return string(s.src[offset:s.offset])
}

func (s *Scanner) scanIdentifier() string {
	offset := s.offset
	for s.isLetter(s.char) || s.isDecimal(s.char) {
		s.next()
	}
	return string(s.src[offset:s.offset])
}

func (s *Scanner) scanDigits(base int) {
	for s.digitVal(s.char) < base {
		s.next()
	}
}

func (s *Scanner) scanNumber() (Token, string) {
	offset := s.offset
	token := INT

	if s.char != '.' {
		if s.char == '0' {
			s.next()
			if s.char != '.' {
				base := 10
				switch s.lower(s.char) {
				case 'x':
					base = 16
				case 'b':
					base = 2
				case 'o':
					base = 8
				default:
					if s.isDecimal(s.char) {
						s.error(offset, "invalid integer")
						token = ILLEGAL
					} else {
						return token, "0"
					}
				}
				if token != ILLEGAL {
					s.next()
					s.scanDigits(base)
					if s.offset-offset <= 2 {
						// only scanned "0x" or "0X"
						token = ILLEGAL
						s.error(offset, "illegal number")
					}
					if s.char == '.' {
						token = ILLEGAL
						s.error(offset, "invalid radix point")
					}
				}
			}
		} else {
			s.scanDigits(10)
		}
	}

	if s.char == '.' {
		offsetFraction := s.offset
		token = FLOAT
		s.next()
		s.scanDigits(10)
		if offsetFraction == s.offset-1 {
			token = ILLEGAL
			s.error(offset, "float has no digits after .")
		}
	}

	return token, string(s.src[offset:s.offset])
}

func (s *Scanner) scanEscape(quote rune) bool {
	offset := s.offset

	var n int
	var base, max uint32
	switch s.char {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		s.next()
		return true
	case '0', '1', '2', '3', '4', '5', '6', '7':
		n, base, max = 3, 8, 255
	case 'x':
		s.next()
		n, base, max = 2, 16, 255
	case 'u':
		s.next()
		n, base, max = 4, 16, unicode.MaxRune
	case 'U':
		s.next()
		n, base, max = 8, 16, unicode.MaxRune
	default:
		msg := "unknown escape sequence"
		if s.char < 0 {
			msg = "escape sequence not terminated"
		}
		s.error(offset, msg)
		return false
	}

	var x uint32
	for n > 0 {
		d := uint32(s.digitVal(s.char))
		if d >= base {
			msg := fmt.Sprintf("illegal character %#U in escape sequence", s.char)
			if s.char < 0 {
				msg = "escape sequence not terminated"
			}
			s.error(s.offset, msg)
			return false
		}
		x = x*base + d
		s.next()
		n--
	}

	if x > max || 0xD800 <= x && x < 0xE000 {
		s.error(offset, "escape sequence is invalid Unicode code point")
		return false
	}

	return true
}

func (s *Scanner) scanString() string {
	offset := s.offset - 1

	for {
		char := s.char
		if char == '\n' || char < 0 {
			s.error(s.offset, "string literal not terminated")
			break
		}
		s.next()
		if char == '"' {
			break
		}
		if char == '\\' {
			s.scanEscape('"')
		}
	}

	return string(s.src[offset:s.offset])
}

func (s *Scanner) scanChar() string {
	// '\'' opening already consumed
	//TO-DO conver to uint32
	offset := s.offset - 1

	valid := true
	n := 0
	for {
		char := s.char
		if char == '\n' || char < 0 {
			if valid {
				s.error(offset, "rune literal not terminated")
				valid = false
			}
			break
		}
		s.next()
		if char == '\'' {
			break
		}
		n++
		if char == '\\' {
			switch s.char {
			case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', '\'':
				s.next()
			default:
				s.error(offset, "illegal char literal")
				valid = false
			}
		}
	}

	if valid && n != 1 {
		s.error(offset, "illegal char literal")
	}

	return string(s.src[offset:s.offset])
}

func (s *Scanner) scanRawString() string {
	// '`' opening already consumed
	offset := s.offset - 1
	for {
		char := s.char
		if char < 0 {
			s.error(s.offset, "raw string literal not terminated")
			break
		}
		s.next()
		if char == '`' {
			break
		}
	}
	return string(s.src[offset:s.offset])
}

func (s *Scanner) scanOperators() (token Token, literal string) {
	offset := s.offset - 1
	length := 0
	token, length = ReadOperator(s.src[offset:])
	if length > 0 {
		for i := 1; i < length; i++ {
			s.next()
		}
		literal = string(s.src[offset:s.offset])
	}
	return
}

func (s *Scanner) scanPreprossesor() bool {
	// skip whitespace but not newline
	for s.char == ' ' || s.char == '\t' {
		s.next()
	}

	notOp := false
	if s.char == '!' {
		notOp = true
		s.next()
	}

	if s.isLetter(s.char) {
		flag := s.scanIdentifier()

		for s.char == ' ' || s.char == '\t' || s.char == '\r' {
			s.next()
		}
		if s.char != '\n' {
			s.error(s.offset, "unexpected: "+string(s.char))
		}

		result := false
		if _, ok := s.flags[flag]; ok {
			result = true
		}
		if notOp {
			result = !result
		}
		return result
	}

	s.error(s.offset, "#if not terminated")
	return false
}

func (s *Scanner) skipPreprossesor() {
	for s.char != '#' {
		s.next()
	}
	s.next()
	if s.isLetter(s.char) {
		literal := s.scanIdentifier()
		if literal != "end" {
			s.error(s.offset, "unexpected: "+literal)
		}
	} else {
		s.error(s.offset, "unexpected: "+string(s.char))
	}
}

func (s *Scanner) isLetter(char rune) bool {
	return char == '_' || 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z'
}

func (s *Scanner) isDecimal(char rune) bool {
	return '0' <= char && char <= '9'
}

// returns lower-case char if char is ASCII letter
// use 0b00100000 instead 'a' - 'A' later in panda own compiler
func (s *Scanner) lower(char rune) rune {
	return ('a' - 'A') | char
}

func (s *Scanner) digitVal(char rune) int {
	switch {
	case '0' <= char && char <= '9':
		return int(char - '0')
	case 'a' <= s.lower(char) && s.lower(char) <= 'f':
		return int(s.lower(char) - 'a' + 10)
	}
	return 16 // larger than any legal digit val
}

// Scan next token
func (s *Scanner) Scan() (pos int, token Token, literal string) {
	for s.char == ' ' || s.char == '\t' || s.char == '\n' || s.char == '\r' {
		s.next()
	}

	pos = s.offset

	token = ILLEGAL
	if s.isLetter(s.char) {
		literal = s.scanIdentifier()
		token = GetToken(literal)
	} else if s.isDecimal(s.char) || (s.char == '.' && s.isDecimal(rune(s.peek()))) {
		token, literal = s.scanNumber()
	} else {
		char := s.char
		s.next()
		switch char {
		case eof:
			token = EOF
			if s.flagOpen {
				s.error(s.offset, "#if not terminated, expecting #end")
			}
		case '"':
			token = STRING
			literal = s.scanString()
		case '`':
			token = STRING
			literal = s.scanRawString()
		case '\'':
			token = CHAR
			literal = s.scanChar()
		case '/': // alse maybe operator /
			if s.char == '/' || s.char == '*' {
				s.scanComment()
				return s.Scan()
			}
			token, literal = s.scanOperators()
		case '@':
			token = META
			literal = "@"
		case ';':
			token = Semi
			literal = ";"
		case '#':
			//#if #end, flag can add '!' for not operation
			//nested # is not supported
			//#else is not supported
			if s.char == 'i' || s.char == 'e' {
				literal = s.scanIdentifier()

				if literal == "if" {
					if s.flagOpen {
						s.error(s.offset, "unexpected #if")
					}
					s.flagOpen = true

					result := s.scanPreprossesor()

					if !result {
						s.skipPreprossesor()
						s.flagOpen = false
					}
				} else if literal == "end" {
					if !s.flagOpen {
						s.error(s.offset, "unexpected #end")
					}
					s.flagOpen = false
				} else {
					s.error(s.offset, "unexpected: "+literal)
				}
				return s.Scan()
			} else {
				s.error(s.offset, "unexpected: "+string(s.char))
			}
		default:
			token, literal = s.scanOperators()
			if token == ILLEGAL {
				s.error(s.offset, "invalid token")
				s.next()
			}
		}
	}
	return
}
