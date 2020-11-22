package main

import "bytes"

type Token int

//TO-DO inline, async, yield
const (
	// special tokens
	ILLEGAL Token = iota
	EOF

	// literals
	literal_begin
	IDENT
	INT
	FLOAT
	CHAR
	STRING
	META
	literal_end

	// keywords
	keyword_begin
	Async
	Await
	Base
	Break
	Case
	Catch
	Class
	Const
	Continue
	Default
	Else
	Enum
	False
	For
	Function
	If
	Import
	Inline
	Interface
	New
	Null
	Namespace
	Public
	Return
	Static
	Switch
	This
	Throw
	True
	Try
	Var
	Yield
	keyword_end

	// scalars
	scalar_begin
	Bool
	Char //16 bit unicode char
	Int8
	Int16
	Int32
	Int64
	Uint8
	Uint16
	Uint32
	Uint64
	SByte //int8
	Short
	Int
	Long
	Byte //uint8
	Ushort
	Uint
	Ulong
	Float32
	Float64
	Float
	Double
	String
	Void
	scalar_end

	// operators
	operator_begin
	LeftParen
	RightParen
	LeftBracket
	RightBracket
	LeftBrace
	RightBrace

	Plus
	Minus
	Star
	Div
	Mod
	Caret
	And
	Or
	Tilde
	Not
	Assign
	Less
	Greater

	PlusAssign
	MinusAssign
	MulAssign
	DivAssign
	ModAssign
	XorAssign
	AndAssign
	OrAssign
	LeftShift
	RightShift
	LeftShiftAssign
	RightShiftAssign

	Equal
	NotEqual
	LessEqual
	GreaterEqual
	AndAnd
	OrOr
	PlusPlus
	MinusMinus

	Comma
	Question
	Colon
	Semi
	Dot
	Ellipsis
	operator_end
)

var (
	tokens = [...]string{
		Async:     "async",
		Await:     "await",
		Base:      "base",
		Break:     "break",
		Case:      "case",
		Catch:     "catch",
		Class:     "class",
		Const:     "const",
		Continue:  "continue",
		Default:   "default",
		Else:      "else",
		Enum:      "enum",
		False:     "false",
		For:       "for",
		Function:  "function",
		If:        "if",
		Import:    "import",
		Inline:    "inline",
		Interface: "interface",
		New:       "new",
		Null:      "null",
		Namespace: "namespace",
		Public:    "public",
		Return:    "return",
		Static:    "static",
		Switch:    "switch",
		This:      "this",
		Throw:     "throw",
		True:      "true",
		Try:       "try",
		Var:       "var",

		Bool:    "bool",
		Char:    "char",
		Int8:    "i8",
		Int16:   "i16",
		Int32:   "i32",
		Int64:   "i64",
		Uint8:   "u8",
		Uint16:  "u16",
		Uint32:  "u32",
		Uint64:  "u64",
		SByte:   "sbyte",
		Short:   "short",
		Int:     "int",
		Long:    "long",
		Byte:    "byte",
		Ushort:  "ushort",
		Uint:    "uint",
		Ulong:   "ulong",
		Float32: "f32",
		Float64: "f64",
		Float:   "float",
		Double:  "double",
		String:  "string",
		Void:    "void",

		LeftParen:        "(",
		RightParen:       ")",
		LeftBracket:      "[",
		RightBracket:     "]",
		LeftBrace:        "{",
		RightBrace:       "}",
		Plus:             "+",
		Minus:            "-",
		Star:             "*",
		Div:              "/",
		Mod:              "%",
		Caret:            "^",
		And:              "&",
		Or:               "|",
		Tilde:            "~",
		Not:              "!",
		Assign:           "=",
		Less:             "<",
		Greater:          ">",
		PlusAssign:       "+=",
		MinusAssign:      "-=",
		MulAssign:        "*=",
		DivAssign:        "/=",
		ModAssign:        "%=",
		XorAssign:        "^=",
		AndAssign:        "&=",
		OrAssign:         "|=",
		LeftShift:        "<<",
		RightShift:       ">>",
		LeftShiftAssign:  "<<=",
		RightShiftAssign: ">>=",
		Equal:            "==",
		NotEqual:         "!=",
		LessEqual:        "<=",
		GreaterEqual:     ">=",
		AndAnd:           "&&",
		OrOr:             "||",
		PlusPlus:         "++",
		MinusMinus:       "--",
		Comma:            ",",
		Question:         "?",
		Colon:            ":",
		Semi:             ";",
		Dot:              ".",
		Ellipsis:         "...",
	}

	cppTokens = [...]string{
		Break:     "break",
		Base:      "base",
		Case:      "case",
		Catch:     "catch",
		Class:     "class",
		Const:     "const",
		Continue:  "continue",
		Default:   "default",
		Else:      "else",
		Enum:      "enum class",
		False:     "false",
		For:       "for",
		Function:  "",
		Import:    "#include",
		Interface: "class",
		If:        "if",
		Inline:    "inline",
		New:       "new",
		Null:      "nullptr",
		Namespace: "namespace",
		Public:    "",
		Return:    "return",
		Static:    "static",
		Switch:    "switch",
		This:      "this",
		Throw:     "throw",
		True:      "true",
		Try:       "try",
		Var:       "",

		Bool:    "bool",
		Char:    "uint32_t",
		Int8:    "int8_t",
		Int16:   "int16_t",
		Int32:   "int32_t",
		Int64:   "int64_t",
		Uint8:   "uint8_t",
		Uint16:  "uint16_t",
		Uint32:  "uint32_t",
		Uint64:  "uint64_t",
		SByte:   "int8_t",
		Short:   "int16_t",
		Int:     "int32_t",
		Long:    "int64_t",
		Byte:    "uint8_t",
		Ushort:  "uint16_t",
		Uint:    "uint32_t",
		Ulong:   "uint64_t",
		Float32: "float",
		Float64: "double",
		Float:   "float",
		Double:  "double",
		String:  "std::string",
		Void:    "void",

		LeftParen:        "(",
		RightParen:       ")",
		LeftBracket:      "[",
		RightBracket:     "]",
		LeftBrace:        "{",
		RightBrace:       "}",
		Plus:             "+",
		Minus:            "-",
		Star:             "*",
		Div:              "/",
		Mod:              "%",
		Caret:            "^",
		And:              "&",
		Or:               "|",
		Tilde:            "~",
		Not:              "!",
		Assign:           "=",
		Less:             "<",
		Greater:          ">",
		PlusAssign:       "+=",
		MinusAssign:      "-=",
		MulAssign:        "*=",
		DivAssign:        "/=",
		ModAssign:        "%=",
		XorAssign:        "^=",
		AndAssign:        "&=",
		OrAssign:         "|=",
		LeftShift:        "<<",
		RightShift:       ">>",
		LeftShiftAssign:  "<<=",
		RightShiftAssign: ">>=",
		Equal:            "==",
		NotEqual:         "!=",
		LessEqual:        "<=",
		GreaterEqual:     ">=",
		AndAnd:           "&&",
		OrOr:             "||",
		PlusPlus:         "++",
		MinusMinus:       "--",
		Comma:            ",",
		Question:         "?",
		Colon:            ":",
		Semi:             ";",
		Dot:              ".",
		Ellipsis:         "...",
	}

	allTokens    map[string]Token
	operatorRoot *OperatorNode
)

func init() {
	allTokens = make(map[string]Token)

	for i := keyword_begin + 1; i < keyword_end; i++ {
		allTokens[tokens[i]] = i
	}

	for i := scalar_begin + 1; i < scalar_end; i++ {
		allTokens[tokens[i]] = i
	}

	operatorRoot = NewOperatorNode()
	for i := operator_begin + 1; i < operator_end; i++ {
		allTokens[tokens[i]] = i
		operatorRoot.Insert(tokens[i])
	}
}

func (token Token) String() string {
	if 0 <= token && token < Token(len(tokens)) {
		return tokens[token]
	}
	return ""
}
func (token Token) IsLiteral() bool {
	return literal_begin < token && token < literal_end
}

func (token Token) IsOperator() bool {
	return operator_begin < token && token < operator_end
}

func (token Token) IsKeyword() bool {
	return keyword_begin < token && token < keyword_end
}

func (token Token) IsScalar() bool {
	return scalar_begin < token && token < scalar_end
}

func (token Token) Print(buffer *bytes.Buffer) {
	if int(token) < len(cppTokens) {
		buffer.WriteString(cppTokens[token])
	}
}

func GetToken(identifier string) Token {
	if token, ok := allTokens[identifier]; ok {
		return token
	}
	return IDENT
}

func ReadOperator(bytes []byte) (Token, int) {
	return operatorRoot.Find(bytes)
}

type OperatorNode struct {
	children map[byte]*OperatorNode
	token    Token
}

func NewOperatorNode() *OperatorNode {
	return &OperatorNode{
		children: make(map[byte]*OperatorNode),
		token:    ILLEGAL,
	}
}

func (node *OperatorNode) Insert(operator string) {
	node.insertOperator(operator, 0)
}

func (node *OperatorNode) Find(bytes []byte) (Token, int) {
	return node.findOperator(bytes, 0)
}

func (node *OperatorNode) findOperator(bytes []byte, offset int) (Token, int) {
	if child, ok := node.children[bytes[offset]]; ok {
		offset++
		if offset < len(bytes) {
			return child.findOperator(bytes, offset)
		}
		return child.token, offset
	} else if offset > 0 {
		return GetToken(string(bytes[:offset])), offset
	}
	return ILLEGAL, 1
}

func (node *OperatorNode) insertOperator(operator string, position int) {
	if position < len(operator) {
		char := operator[position]
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewOperatorNode()
		}
		position++
		node.children[char].insertOperator(operator, position)
	} else {
		node.token = GetToken(operator)
	}
}

// ==========================================================
// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
//
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
//
func (op Token) Precedence() int {
	switch op {
	case OrOr:
		return 1
	case AndAnd:
		return 2
	case Equal, NotEqual, Less, LessEqual, Greater, GreaterEqual, Question:
		return 3
	case Plus, Minus, Or, Caret:
		return 4
	case Star, Div, Mod, LeftShift, RightShift, And:
		return 5
	}
	return LowestPrec
}
