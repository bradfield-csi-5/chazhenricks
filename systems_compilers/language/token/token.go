package token

import (
	"fmt"
	"strconv"
)


type TokenEnum int

const (
  //single char tokens 
	LEFT_PAREN TokenEnum = iota
  
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

  //one or two char tokens 
  BANG
  BANG_EQUAL
  EQUAL 
  EQUAL_EQUAL 
  GREATER
  GREATER_EQUAL 
  LESS 
  LESS_EQUAL 

  //literals 
  IDENTIFIER 
  STRING 
  NUMBER 

  //keywords 
  AND 
  CLASS 
  ELSE 
  FALSE 
  FUN 
  FOR 
  IF 
  NIL 
  OR 
  PRINT 
  RETURN 
  SUPER 
  THIS 
  TRUE 
  VAR 
  WHILE 


  EOF 
)

var tokens = [...]string{
  LEFT_PAREN: "(" ,
  RIGHT_PAREN:")",
  LEFT_BRACE:"{",
  RIGHT_BRACE:"}",
  COMMA:",",
  DOT:".",
  MINUS:"-",
  PLUS:"+",
  SEMICOLON:";",
  SLASH:"/",
  STAR:"*",

  //one or two char tokens 
  BANG:"!",
  BANG_EQUAL:"!!",
  EQUAL :"=",
  EQUAL_EQUAL :"==",
  GREATER:">",
  LESS :"<",
  GREATER_EQUAL :">=",
  LESS_EQUAL :"<=",
  //literals 
  IDENTIFIER :"IDENTIFIER",
  STRING :"STRING",
  NUMBER :"NUMBER",
  //keywords 
  AND :"and",
  CLASS :"class",
  ELSE :"else",
  FALSE :"false",
  FUN :"fun",
  FOR :"for",
  IF :"if",
  NIL :"nil",
  OR :"or",
  PRINT :"print",
  RETURN :"return",
  SUPER :"super",
  THIS :"this",
  TRUE :"true",
  VAR :"var",
  WHILE :"while",
  EOF :"eof",
}


func (tok TokenEnum) String() string {
  s:= ""
  if 0 <= tok && tok < TokenEnum(len(tokens)){
    s = tokens[tok]
  }
  if s == ""{
    s = "token(" + strconv.Itoa(int(tok)) + ")"
  }
  return s
}


type Token struct {
  Type TokenEnum
  Lexeme string
  Literal any 
  Line int
}

func (tok Token) String() string {
  return fmt.Sprintf("type: %s lexeme: %s",tok.Type.String() ,tok.Lexeme)
}
