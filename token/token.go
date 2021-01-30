package token

type Token int

const (
	ILLEGAL Token = iota
	EOF
	META

	// literals
	literalBegin
	IDENT
	BOOL
	INT
	FLOAT
	CHAR
	STRING
	NULL
	literalEnd

	// keywords
	keywordBegin
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
	Finally
	For
	Function
	If
	Import
	Interface
	New
	Namespace
	Public
	Return
	Switch
	This
	Throw
	Try
	Var
	Weak
	keywordEnd

	// scalars
	scalarBegin
	Any
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
	Void
	Pointer
	scalarEnd

	// operators
	operatorBegin
	LeftParen
	RightParen
	LeftBracket
	RightBracket
	LeftBrace
	RightBrace

	Plus
	Minus
	Mul
	Div
	Rem
	BitAnd
	BitOr
	BitXor
	Complement
	Not
	Assign
	Less
	Greater

	PlusAssign
	MinusAssign
	MulAssign
	DivAssign
	RemAssign
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
	And
	Or
	PlusPlus
	MinusMinus

	Comma
	Colon
	Semi
	Dot
	Ellipsis
	operatorEnd
)

var (
	tokenStrings = [...]string{
		IDENT:  "identifier",
		BOOL:   "bool_literal",
		INT:    "int_literal",
		FLOAT:  "float_literal",
		CHAR:   "char_literal",
		STRING: "string_literal",
		NULL:   "null",

		Break:     "break",
		Case:      "case",
		Catch:     "catch",
		Class:     "class",
		Const:     "const",
		Continue:  "continue",
		Default:   "default",
		Else:      "else",
		Enum:      "enum",
		Finally:   "finally",
		For:       "for",
		Function:  "function",
		If:        "if",
		Import:    "import",
		Interface: "interface",
		New:       "new",
		Namespace: "namespace",
		Public:    "public",
		Return:    "return",
		Switch:    "switch",
		This:      "this",
		Throw:     "throw",
		Try:       "try",
		Var:       "var",
		Weak:      "weak",

		Any:     "any",
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
		Void:    "void",
		Pointer: "pointer",

		LeftParen:        "(",
		RightParen:       ")",
		LeftBracket:      "[",
		RightBracket:     "]",
		LeftBrace:        "{",
		RightBrace:       "}",
		Plus:             "+",
		Minus:            "-",
		Mul:              "*",
		Div:              "/",
		Rem:              "%",
		BitAnd:           "&",
		BitOr:            "|",
		BitXor:           "^",
		Complement:       "~",
		Not:              "!",
		Assign:           "=",
		Less:             "<",
		Greater:          ">",
		PlusAssign:       "+=",
		MinusAssign:      "-=",
		MulAssign:        "*=",
		DivAssign:        "/=",
		RemAssign:        "%=",
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
		And:              "&&",
		Or:               "||",
		PlusPlus:         "++",
		MinusMinus:       "--",
		Comma:            ",",
		Colon:            ":",
		Semi:             ";",
		Dot:              ".",
		Ellipsis:         "...",
	}

	tokens map[string]Token
)

func init() {
	tokens = make(map[string]Token)

	for i := keywordBegin + 1; i < keywordEnd; i++ {
		tokens[tokenStrings[i]] = i
	}

	for i := scalarBegin + 1; i < scalarEnd; i++ {
		tokens[tokenStrings[i]] = i
	}

	for i := operatorBegin + 1; i < operatorEnd; i++ {
		tokens[tokenStrings[i]] = i
	}

	operatorRoot = &operatorNode{
		children: make(map[byte]*operatorNode),
		token:    ILLEGAL,
	}

	for i := operatorBegin + 1; i < operatorEnd; i++ {
		operatorRoot.insert(tokenStrings[i])
	}
}

func ReadToken(literal string) Token {
	if token, ok := tokens[literal]; ok {
		return token
	}
	if literal == "true" || literal == "false" {
		return BOOL
	}
	if literal == "null" {
		return NULL
	}
	return IDENT
}

func (t Token) String() string {
	if 0 <= t && t < Token(len(tokenStrings)) {
		return tokenStrings[t]
	}
	return ""
}

func (t Token) IsLiteral() bool {
	return literalBegin < t && t < literalEnd
}

func (t Token) IsOperator() bool {
	return operatorBegin < t && t < operatorEnd
}

func (t Token) IsKeyword() bool {
	return keywordBegin < t && t < keywordEnd
}

func (t Token) IsScalar() bool {
	return scalarBegin < t && t < scalarEnd
}

func (t Token) Precedence() int {
	switch t {
	case Assign, MulAssign, DivAssign, RemAssign, PlusAssign, MinusAssign,
		LeftShiftAssign, RightShiftAssign, AndAssign, OrAssign, XorAssign:
		return 1

	case Or:
		return 2

	case And:
		return 3

	case BitOr:
		return 4

	case BitXor:
		return 5

	case BitAnd:
		return 6

	case Equal, NotEqual:
		return 7

	case Less, LessEqual, Greater, GreaterEqual:
		return 8

	case LeftShift, RightShift:
		return 9

	case Plus, Minus:
		return 10

	case Mul, Div, Rem:
		return 11
	}
	return 0
}
