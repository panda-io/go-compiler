package compiler

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/panda-foundation/go-compiler/token"
)

const (
	bom = 0xFEFF // byte order mark, only permitted as very first character
	eof = -1

	//#if #else #elif #end
	preprocessorIf     = "if"
	preprocessorElse   = "else"
	preprocessorElseIf = "elif"
	preprocessorEnd    = "end"

	// TO-DO add logical check for preprocessor
	// () == != ! && ||
)

type preprocessor struct {
	currentBlock string
	satisfied    bool
}

// ErrorHandler is callback when error occurs
type ErrorHandler func(*token.Position, string)

// Scanner to scan token
type Scanner struct {
	file *token.File
	src  []byte

	errorHandler ErrorHandler

	flags              map[string]bool // flags for condition compiler
	preprocessorLevel  int             // for nested flag
	preprocessorStatus []*preprocessor

	char       rune
	offset     int
	readOffset int
}

// NewScanner create new scanner
func NewScanner(handler ErrorHandler, flags []string) *Scanner {
	s := &Scanner{
		errorHandler: handler,
	}
	for _, flag := range flags {
		s.flags[flag] = true
	}
	return s
}

// ScanFile start to scan new file
func (s *Scanner) ScanFile(file *token.File, src []byte) {
	s.file = file
	s.src = src

	s.preprocessorLevel = 0
	s.preprocessorStatus = s.preprocessorStatus[:0]
	s.char = ' '
	s.offset = 0
	s.readOffset = 0

	s.next()
	if s.char == bom {
		s.next()
	}
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

func (s *Scanner) error(offset int, message string) {
	if s.errorHandler != nil {
		s.errorHandler(s.file.Position(offset), message)
	}
}

func (s *Scanner) scanComment() string {
	offset := s.offset - 1
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

func (s *Scanner) scanNumber() (token.Token, string) {
	offset := s.offset
	t := token.INT

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
						t = token.ILLEGAL
					} else {
						return t, "0"
					}
				}
				if t != token.ILLEGAL {
					s.next()
					s.scanDigits(base)
					if s.offset-offset <= 2 {
						t = token.ILLEGAL
						s.error(offset, "illegal number")
					}
					if s.char == '.' {
						t = token.ILLEGAL
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
		t = token.FLOAT
		s.next()
		s.scanDigits(10)
		if offsetFraction == s.offset-1 {
			t = token.ILLEGAL
			s.error(offset, "float has no digits after .")
		}
	}

	return t, string(s.src[offset:s.offset])
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
		m := "unknown escape sequence"
		if s.char < 0 {
			m = "escape sequence not terminated"
		}
		s.error(offset, m)
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

func (s *Scanner) scanOperators() (t token.Token, literal string) {
	offset := s.offset - 1
	t, length := token.ReadOperator(s.src[offset:])
	if length > 0 {
		for i := 1; i < length; i++ {
			s.next()
		}
		literal = string(s.src[offset:s.offset])
	}
	return
}

func (s *Scanner) scanPreprossesor() (position int, t token.Token, literal string) {
	//#if #else #elif #end
	if !s.isLetter(s.char) {
		s.error(s.offset, "unexpected identifier")
	}
	literal = s.scanIdentifier()
	if literal == preprocessorIf {
		s.preprocessorLevel++
		s.preprocessorStatus = append(s.preprocessorStatus, &preprocessor{
			currentBlock: preprocessorIf,
			satisfied:    false,
		})

		result := s.scanPreprossesorExpression()
		if result {
			s.preprocessorStatus[s.preprocessorLevel-1].satisfied = true
		} else {
			s.skipPreprossesor()
		}
	} else if literal == preprocessorElseIf {
		if s.preprocessorLevel == 0 || s.preprocessorStatus[s.preprocessorLevel-1].currentBlock == preprocessorElse {
			s.error(s.offset, "unexpected #elif")
		} else if s.preprocessorStatus[s.preprocessorLevel-1].satisfied {
			s.skipPreprossesor()
		} else {
			if s.scanPreprossesorExpression() {
				s.preprocessorStatus[s.preprocessorLevel-1].satisfied = true
			} else {
				s.skipPreprossesor()
			}
		}
		s.preprocessorStatus[s.preprocessorLevel-1].currentBlock = preprocessorElseIf
	} else if literal == preprocessorElse {
		if s.preprocessorLevel == 0 || s.preprocessorStatus[s.preprocessorLevel-1].currentBlock == preprocessorElse {
			s.error(s.offset, "unexpected #else")
		} else if s.preprocessorStatus[s.preprocessorLevel-1].satisfied {
			s.skipPreprossesor()
		}
		s.preprocessorStatus[s.preprocessorLevel-1].currentBlock = preprocessorElse
	} else if literal == preprocessorEnd {
		if s.preprocessorLevel == 0 {
			s.error(s.offset, "unexpected #end")
		}
		s.preprocessorLevel--
		s.preprocessorStatus = s.preprocessorStatus[:len(s.preprocessorStatus)-1]
	} else {
		s.error(s.offset, "unexpected preprocessor: "+literal)
	}

	return s.Scan()
}

func (s *Scanner) scanPreprossesorExpression() bool {
	for s.char == ' ' || s.char == '\t' {
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
		return result
	}

	s.error(s.offset, "unexpected: "+string(s.char))
	return false
}

func (s *Scanner) skipPreprossesor() {
	level := s.preprocessorLevel
	for {
		for s.char != eof || s.char != '#' {
			s.next()
		}
		if s.char == eof {
			s.error(s.offset, "preprocessor not terminated, expecting #end")
		}
		s.next()
		if s.isLetter(s.char) {
			literal := s.scanIdentifier()

			if literal == preprocessorIf {
				s.preprocessorLevel++
				s.preprocessorStatus = append(s.preprocessorStatus, &preprocessor{
					currentBlock: preprocessorIf,
					satisfied:    false,
				})
			} else if literal == preprocessorElseIf {
				if s.preprocessorStatus[s.preprocessorLevel-1].currentBlock == preprocessorElse {
					s.error(s.offset, "unexpected #elif")
				}
				s.preprocessorStatus[s.preprocessorLevel-1].currentBlock = preprocessorElseIf
				if s.preprocessorLevel == level {
					break
				}
			} else if literal == preprocessorElse {
				if s.preprocessorStatus[s.preprocessorLevel-1].currentBlock == preprocessorElse {
					s.error(s.offset, "unexpected #else")
				}
				s.preprocessorStatus[s.preprocessorLevel-1].currentBlock = preprocessorElse
				if s.preprocessorLevel == level {
					break
				}
			} else if literal == preprocessorEnd {
				s.preprocessorLevel--
				s.preprocessorStatus = s.preprocessorStatus[:len(s.preprocessorStatus)-1]
				if s.preprocessorLevel == level-1 {
					break
				}
			} else {
				s.error(s.offset, "unexpected preprocessor: "+literal)
			}
		} else {
			s.error(s.offset, "expected identifier")
		}
	}
}

func (s *Scanner) isLetter(char rune) bool {
	return char == '_' || 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z'
}

func (s *Scanner) isDecimal(char rune) bool {
	return '0' <= char && char <= '9'
}

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
func (s *Scanner) Scan() (position int, t token.Token, literal string) {
	for s.char == ' ' || s.char == '\t' || s.char == '\n' || s.char == '\r' {
		s.next()
	}

	position = s.offset

	t = token.ILLEGAL
	if s.isLetter(s.char) {
		literal = s.scanIdentifier()
		t = token.ReadToken(literal)
	} else if s.isDecimal(s.char) || (s.char == '.' && s.isDecimal(rune(s.peek()))) {
		t, literal = s.scanNumber()
	} else {
		char := s.char
		s.next()
		switch char {
		case eof:
			t = token.EOF
			if s.preprocessorLevel > 0 {
				s.error(s.offset, "preprocessor not terminated, expecting #end")
			}
		case '"':
			t = token.STRING
			literal = s.scanString()
		case '`':
			t = token.STRING
			literal = s.scanRawString()
		case '\'':
			t = token.CHAR
			literal = s.scanChar()
		case '/':
			if s.char == '/' || s.char == '*' {
				s.scanComment()
				return s.Scan()
			}
			t, literal = s.scanOperators()
		case '@':
			t = token.META
			literal = "@"
		case ';':
			t = token.Semi
			literal = ";"
		case '#':
			return s.scanPreprossesor()

		default:
			t, literal = s.scanOperators()
			if t == token.ILLEGAL {
				s.error(s.offset, "invalid token")
				s.next()
			}
		}
	}
	return
}
