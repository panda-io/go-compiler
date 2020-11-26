package token

// Token present language token
type Token int

// all available token list
const (
	ILLEGAL Token = iota
	EOF
	META

	// literals
	literalBegin
	IDENT
	INT
	FLOAT
	CHAR
	STRING
	literalEnd

	// keywords
	keywordBegin
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
		Yield:     "yield",

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

// ReadToken read literal to token, or identifier
func ReadToken(literal string) Token {
	if token, ok := tokens[literal]; ok {
		return token
	}
	return IDENT
}

func (t Token) String() string {
	if 0 <= t && t < Token(len(tokenStrings)) {
		return tokenStrings[t]
	}
	return ""
}

// IsLiteral check if token is literal
func (t Token) IsLiteral() bool {
	return literalBegin < t && t < literalEnd
}

// IsOperator check if token is operator
func (t Token) IsOperator() bool {
	return operatorBegin < t && t < operatorEnd
}

// IsKeyword check if token is keyword
func (t Token) IsKeyword() bool {
	return keywordBegin < t && t < keywordEnd
}

// IsScalar check if token is built-in type
func (t Token) IsScalar() bool {
	return scalarBegin < t && t < scalarEnd
}

// CppType convert to cpp type (when export)
func (t Token) CppType() string {
	if t.IsScalar() {
		return cppTypes[t]
	}
	panic("invalid type:" + t.String())
}

// Precedence returns the operator precedence of the binary operator op
func (t Token) Precedence() int {
	switch t {
	case Or:
		return 1
	case And:
		return 2
	case Equal, NotEqual, Less, LessEqual, Greater, GreaterEqual, Question:
		return 3
	case Plus, Minus, BitOr, Caret:
		return 4
	case Mul, Div, Mod, LeftShift, RightShift, BitAnd:
		return 5
	}
	return 0
}
