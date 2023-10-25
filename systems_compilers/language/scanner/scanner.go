package scanner

import (
	"fmt"
	"language/error"
	"language/token"
)

type Scanner struct {
	Source  string
	Tokens  []token.Token
	Start   int
	Current int
	Line    int
}

func NewScanner(source string) Scanner {
	return Scanner{
		Source:  source,
		Tokens:  make([]token.Token, 0),
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

func (scan *Scanner) ScanTokens() []token.Token {
	for !scan.isAtEnd() {
		scan.Start = scan.Current
		scan.scanToken()
	}

	scan.Tokens = append(scan.Tokens, token.Token{Type: token.EOF, Lexeme: ""})
	return scan.Tokens
}

func (scan *Scanner) scanToken() {
	c := scan.advance()
	switch c {
	case "(":
		scan.addToken(token.LEFT_PAREN)
		break
	case ")":
		scan.addToken(token.RIGHT_PAREN)
		break
	case "{":
		scan.addToken(token.LEFT_BRACE)
		break
	case "}":
		scan.addToken(token.RIGHT_BRACE)
		break
	case ",":
		scan.addToken(token.COMMA)
		break
	case ".":
		scan.addToken(token.DOT)
		break
	case "-":
		scan.addToken(token.MINUS)
		break
	case "+":
		scan.addToken(token.PLUS)
		break
	case ";":
		scan.addToken(token.SEMICOLON)
		break
	case "*":
		scan.addToken(token.STAR)
		break

	//could be 1 or 2 chars
	case "!":
		if scan.match("=") {
			scan.addToken(token.BANG_EQUAL)
		} else {
			scan.addToken(token.BANG)
		}
	default:
		error.CoolError(scan.Line, "Unexpected Character")
		break
	}
}

func HelloScanner() {
	fmt.Println("hello from scanner")
}

func (scan *Scanner) advance() string {
	char := string(scan.Source[scan.Current])
	scan.Current += 1
	return char
}

func (scan *Scanner) match(expected string) bool {
	if scan.isAtEnd() {
		return false
	}
	if string(scan.Source[scan.Current]) != expected {
		return false
	}
	scan.Current += 1
	return true
}

func (scan *Scanner) peek() byte {
	if scan.isAtEnd() {
		return 0
	}
	return scan.Source[scan.Current]
}

func (scan *Scanner) addToken(tok token.TokenEnum) {
	scan.Tokens = append(scan.Tokens, token.Token{Type: tok, Lexeme: tok.String()})
}

func (scan *Scanner) isAtEnd() bool {
	return scan.Current >= len(scan.Source)
}
