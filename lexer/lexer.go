package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // Points to current char
	readPosition int  // Points to next char; helps with peeking
	ch           byte // Current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Puts the lexer in a "working state" before NextToken() is called
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "=="}
			l.readChar()
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
			l.readChar()
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readWord()
			tok.Type = token.LookupWord(tok.Literal)
			return tok // Early return neeeded since readChar() called in readWord()
		} else if isDigit(l.ch) {
			tok.Literal = l.readInteger()
			tok.Type = token.INT
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar() // Advance the lexer's pointers for next time this function is called
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII for "NUL"; end of file or we haven't read anything
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readWord() string {
	// A word is an identifier or a keyword
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readInteger() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
