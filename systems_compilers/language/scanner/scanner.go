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
	case ")":
		scan.addToken(token.RIGHT_PAREN)
	case "{":
		scan.addToken(token.LEFT_BRACE)
	case "}":
		scan.addToken(token.RIGHT_BRACE)
	case ",":
		scan.addToken(token.COMMA)
	case ".":
		scan.addToken(token.DOT)
	case "-":
		scan.addToken(token.MINUS)
	case "+":
		scan.addToken(token.PLUS)
	case ";":
		scan.addToken(token.SEMICOLON)
	case "*":
		scan.addToken(token.STAR)

	//could be 1 or 2 chars
	case "!":
		if scan.match("=") {
			scan.addToken(token.BANG_EQUAL)
		} else {
			scan.addToken(token.BANG)
		}
  case "=":
    if scan.match("="){
      scan.addToken(token.EQUAL_EQUAL)
    }else {
      scan.addToken(token.EQUAL)
    }

  case "<":
    if scan.match("="){
      scan.addToken(token.LESS_EQUAL)
    }else{
      scan.addToken(token.LESS)
    }

  case ">":
    if scan.match("="){
      scan.addToken(token.GREATER_EQUAL)
    }else{
      scan.addToken(token.GREATER)
    }

  case "/":
  if scan.match("/"){
      //we hit a comment
      for scan.peek() != "\n" && !scan.isAtEnd(){
        scan.advance()
      }
    }else{
      scan.addToken(token.SLASH)
    }


	default:
    fmt.Printf("WHAT IS IT: %s\n", c)
		error.CoolError(scan.Line, "Unexpected Character")
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

func (scan *Scanner) peek() string {
	if scan.isAtEnd() {
		return ""
	}
	return string(scan.Source[scan.Current])
}

func (scan *Scanner) addToken(tok token.TokenEnum) {
	scan.Tokens = append(scan.Tokens, token.Token{Type: tok, Lexeme: tok.String()})
}

func (scan *Scanner) isAtEnd() bool {
	return scan.Current >= len(scan.Source)
}
