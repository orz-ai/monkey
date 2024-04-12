package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENT Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1234567890

	// COMMA Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// ASSIGN Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// BANG Operators
	BANG = "!"

	// ASTERISK Operators
	ASTERISK = "*"

	// SLASH Operators
	SLASH = "/"

	// LT and GT Operators
	LT = "<" // less than
	GT = ">" // greater than

	// FUNCTION Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// EQ and NOT_EQ Operators
	EQ     = "=="
	NOT_EQ = "!="

	// STRING Literals
	STRING = "STRING"

	// COLON Delimiters
	COLON = ":"

	// WHILE Keywords
	WHILE = "WHILE"

	// FOR Keywords
	FOR = "FOR"

	// IN Keywords
	IN = "IN"

	// DOT Delimiters
	DOT = "."

	// BREAK Keywords
	BREAK = "BREAK"

	// CONTINUE Keywords
	CONTINUE = "CONTINUE"

	// AND Operators
	AND = "&&"

	// OR Operators
	OR = "||"

	// MOD Operators
	MOD = "%"

	// XOR Operators
	XOR = "^"

	// BIT_AND Operators
	BIT_AND = "&"

	// BIT_OR Operators
	BIT_OR = "|"

	// BIT_NOT Operators
	BIT_NOT = "~"

	// BIT_LSHIFT Operators
	BIT_LSHIFT = "<<"

	// BIT_RSHIFT Operators
	BIT_RSHIFT = ">>"

	// BIT_CLEAR Operators
	BIT_CLEAR = "&^"

	// BIT_XOR Operators
	BIT_XOR = "^"

	// BIT_AND_NOT Operators
	BIT_AND_NOT = "&^"

	// BIT_OR_NOT Operators
	BIT_OR_NOT = "|^"

	// BIT_LSHIFT_NOT Operators
	BIT_LSHIFT_NOT = "<<^"

	// BIT_RSHIFT_NOT Operators
	BIT_RSHIFT_NOT = ">>^"

	// BIT_CLEAR_NOT Operators
	BIT_CLEAR_NOT = "&^"

	// BIT_XOR_NOT Operators
	BIT_XOR_NOT = "^"
)

// LookupIdent takes a string identifier and returns the corresponding TokenType.
//
// Parameter: ident string
// Return type: TokenType
func LookupIdent(ident string) TokenType {
	switch ident {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	case "while":
		return WHILE
	case "for":
		return FOR
	case "in":
		return IN
	case "break":
		return BREAK
	case "continue":
		return CONTINUE
	default:
		return IDENT
	}
}

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"while":    WHILE,
	"for":      FOR,
	"in":       IN,
	"break":    BREAK,
	"continue": CONTINUE,

	"and":    AND,
	"or":     OR,
	"mod":    MOD,
	"xor":    XOR,
	"bitand": BIT_AND,

	"bitor":     BIT_OR,
	"bitnot":    BIT_NOT,
	"bitlshift": BIT_LSHIFT,
	"bitrshift": BIT_RSHIFT,
	"bitclear":  BIT_CLEAR,
	"bitxor":    BIT_XOR,

	"bitandnot": BIT_AND_NOT,
	"bitornot":  BIT_OR_NOT,

	"bitlshiftnot": BIT_LSHIFT_NOT,
	"bitrshiftnot": BIT_RSHIFT_NOT,
	"bitclearnot":  BIT_CLEAR_NOT,
	"bitxornot":    BIT_XOR_NOT,
}

type Token struct {
	Type    TokenType
	Literal string
}
