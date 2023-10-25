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
	case '(':
		scan.addToken(token.LEFT_PAREN)
	case ')':
		scan.addToken(token.RIGHT_PAREN)
	case '{':
		scan.addToken(token.LEFT_BRACE)
	case '}':
		scan.addToken(token.RIGHT_BRACE)
	case ',':
		scan.addToken(token.COMMA)
	case '.':
		scan.addToken(token.DOT)
	case '-':
		scan.addToken(token.MINUS)
	case '+':
		scan.addToken(token.PLUS)
	case ';':
		scan.addToken(token.SEMICOLON)
	case '*':
		scan.addToken(token.STAR)

	//could be 1 or 2 chars
	case '!':
		if scan.match('=') {
			scan.addToken(token.BANG_EQUAL)
		} else {
			scan.addToken(token.BANG)
		}
	case '=':
		if scan.match('=') {
			scan.addToken(token.EQUAL_EQUAL)
		} else {
			scan.addToken(token.EQUAL)
		}

	case '<':
		if scan.match('=') {
			scan.addToken(token.LESS_EQUAL)
		} else {
			scan.addToken(token.LESS)
		}

	case '>':
		if scan.match('=') {
			scan.addToken(token.GREATER_EQUAL)
		} else {
			scan.addToken(token.GREATER)
		}

	case '/':
		if scan.match('/') {
			//we hit a comment
			for scan.peek() != '\n' && !scan.isAtEnd() {
				scan.advance()
			}
		} else {
			scan.addToken(token.SLASH)
		}

	case ' ':
		fallthrough
	case '\r':
		fallthrough
	case '\t':
		//ignore spaces
		return
	case '\n':
		scan.Line += 1
		return
	case '"':
		scan.string()
	default:
		fmt.Printf("WHAT IS IT: %s\n", c)
		error.CoolError(scan.Line, "Unexpected Character")
	}
}

func HelloScanner() {
	fmt.Println("hello from scanner")
}

// we hit an open string character
func (scan *Scanner) string() {
	//while we dont hit the closing " or the end of the file
	for scan.peek() != '"' && !scan.isAtEnd() {
		//if we hit a new line, increment the counter
		//this will have the effect of letting us have multi-line strings
		if scan.peek() == '\n' {
			scan.Line += 1
		}
		//while were still in the loop, keep advancing
		scan.advance()
	}

	if scan.isAtEnd() {
		error.CoolError(scan.Line, "unterminated string")
	}

	//get that last " outta here
	scan.advance()

	//once we exit the loop get the string value
	//+1 and -1 to account for the "" chars
	strValue := scan.Source[scan.Start+1 : scan.Current-1]
	scan.addString(strValue)

}

func (scan *Scanner) advance() byte {
	char := scan.Source[scan.Current]
	scan.Current += 1
	return char
}

func (scan *Scanner) match(expected byte) bool {
	if scan.isAtEnd() {
		return false
	}
	if scan.Source[scan.Current] != expected {
		return false
	}
	scan.Current += 1
	return true
}

func (scan *Scanner) peek() byte {
	if scan.isAtEnd() {
		return ' '
	}
	return scan.Source[scan.Current]
}

func (scan *Scanner) addToken(tok token.TokenEnum) {
	scan.Tokens = append(scan.Tokens, token.Token{Type: tok, Lexeme: tok.String()})
}

func (scan *Scanner) addString(val string) {
	scan.Tokens = append(scan.Tokens, token.Token{Type: token.STRING, Lexeme: val})
}

func (scan *Scanner) isAtEnd() bool {
	return scan.Current >= len(scan.Source)
}
