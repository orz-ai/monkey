// lexer/lexer.go

package lexer

import (
	"testing"

	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // initialize lexer
	l.readChar()              // read the first character
	return l
}

// readChar reads the next character in the input and advances the position pointers accordingly.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // if we reach the end of the input
		l.ch = 0 // set ch to 0, the ASCII code for the "NUL" character
	} else {
		l.ch = l.input[l.readPosition] // set ch to the current character
	}

	l.position = l.readPosition // set position to readPosition
	l.readPosition += 1         // increment readPosition
}

// NextToken gets the next token from the input and advances the lexer cursor.
func (l *Lexer) NextToken() token.Token {
	var tempToken token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=': // check for equality operator
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch) // ==
			tempToken = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tempToken = newToken(token.ASSIGN, l.ch)
		}

	case ';':
		tempToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		tempToken = newToken(token.LPAREN, l.ch)
	case ')':
		tempToken = newToken(token.RPAREN, l.ch)
	case ',':
		tempToken = newToken(token.COMMA, l.ch)
	case '+':
		tempToken = newToken(token.PLUS, l.ch)
	case '-':
		tempToken = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' { // check for inequality operator
			ch := l.ch
			l.readChar()                         // read the next character
			literal := string(ch) + string(l.ch) // !=
			tempToken = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tempToken = newToken(token.BANG, l.ch)
		}
	case '/':
		tempToken = newToken(token.SLASH, l.ch)
	case '*':
		tempToken = newToken(token.ASTERISK, l.ch)
	case '<':
		tempToken = newToken(token.LT, l.ch)
	case '>':
		tempToken = newToken(token.GT, l.ch)
	case '{':
		tempToken = newToken(token.LBRACE, l.ch)
	case '}':
		tempToken = newToken(token.RBRACE, l.ch)
	case 0:
		tempToken.Literal = ""
		tempToken.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tempToken.Literal = l.readIdentifier()
			tempToken.Type = token.LookupIdent(tempToken.Literal)
			return tempToken
		} else if isDigit(l.ch) {
			tempToken.Type = token.INT
			tempToken.Literal = l.readNumber()
			return tempToken
		} else {
			tempToken = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tempToken
}

// readIdentifier reads and returns an identifier from the input.
func (l *Lexer) readIdentifier() string {
	position := l.position

	l.skipWhitespace()

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position] // return the identifier
}

// readNumber reads a number from the input.
//
// No parameters.
// Returns a string.
func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position] // return the number
}

// isLetter checks if the input byte is a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' ||
		ch == '_' // allow underscores in identifiers
}

// isDigit checks if the given byte is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// skipWhitespace skips all whitespaces until it encounters a non-whitespace
// character.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' { // skip all whitespaces
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// peekChar returns the next character in the input without advancing the read position.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) { // if we reach the end of the input
		return 0 // set ch to 0, the ASCII code for the "NUL" character
	} else {
		return l.input[l.readPosition] // set ch to the current character
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
