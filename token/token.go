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
	Static
	Super
	Switch
	This
	Throw
	Try
	Var
	keywordEnd

	// scalars
	scalarBegin
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
	Mod
	Caret
	BitAnd
	BitOr
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
	And
	Or
	PlusPlus
	MinusMinus

	Comma
	Question
	Colon
	Semi
	Dot
	Ellipsis
	operatorEnd
)

var (
	tokenStrings = [...]string{
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
		Static:    "static",
		Super:     "super",
		Switch:    "switch",
		This:      "this",
		Throw:     "throw",
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
		Mul:              "*",
		Div:              "/",
		Mod:              "%",
		Caret:            "^",
		BitAnd:           "&",
		BitOr:            "|",
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
		And:              "&&",
		Or:               "||",
		PlusPlus:         "++",
		MinusMinus:       "--",
		Comma:            ",",
		Question:         "?",
		Colon:            ":",
		Semi:             ";",
		Dot:              ".",
		Ellipsis:         "...",
	}

	cppTypes = [...]string{
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

// TO-DO replace in generator
func (t Token) CppType() string {
	if t.IsScalar() {
		return cppTypes[t]
	}
	panic("invalid type:" + t.String())
}

func (t Token) Precedence() int {
	switch t {
	case Or:
		return 1
	case And:
		return 2
	case Equal, NotEqual, Less, LessEqual, Greater, GreaterEqual:
		return 3
	case Plus, Minus, BitOr, Caret:
		return 4
	case Mul, Div, Mod, LeftShift, RightShift, BitAnd:
		return 5
	}
	return 0
}
