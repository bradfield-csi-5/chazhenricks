package com.craftinginterpreters.lox;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static com.craftinginterpreters.lox.TokenType.*;

class Scanner {
    private final String source; 
    private final List<Token> tokens = new ArrayList<>();
    private int start = 0; //first char in the current lexeme 
    private int current = 0; //index of where we are in the line
    private int line = 1;

    Scanner(String source){
        this.source = source;
    }

     private static final Map<String, TokenType> keywords;
     static {
        keywords = new HashMap<>();
        keywords.put("and", AND);
        keywords.put("class", CLASS);
        keywords.put("else", ELSE);
        keywords.put("false", FALSE);
        keywords.put("for", FOR);
        keywords.put("fun", FUN);
        keywords.put("if", IF);
        keywords.put("nil", NIL);
        keywords.put("or", OR);
        keywords.put("print", PRINT);
        keywords.put("return", RETURN);
        keywords.put("super", SUPER);
        keywords.put("this", THIS);
        keywords.put("true", TRUE);
        keywords.put("var", VAR);
        keywords.put("while", WHILE);
     }





     List<Token> scanTokens() {
        while (!isAtEnd()){
            start = current; 
            scanToken();
        }
        tokens.add(new Token(EOF, "", null, line));
        return tokens;
    }

    private void scanToken() {
        char c = advance();
        switch(c){
            //single character lexemes
            case '(' : addToken(LEFT_PAREN); break;
            case ')' : addToken(RIGHT_PAREN); break;
            case '{' : addToken(LEFT_BRACE); break;
            case '}' : addToken(RIGHT_BRACE); break;
            case ',' : addToken(COMMA); break;
            case '.' : addToken(DOT); break;
            case '-' : addToken(MINUS); break;
            case '+' : addToken(PLUS); break;
            case ';' : addToken(SEMICOLON); break;
            case '*' : addToken(STAR); break;

            //single/double char lexemes
            case '!': addToken(match('=') ? BANG_EQUAL : BANG); break;
            case '=': addToken(match('=') ? EQUAL_EQUAL : EQUAL);break;
            case '<': addToken(match('=') ? LESS_EQUAL : LESS);break;
            case '>': addToken(match('=') ? GREATER_EQUAL : GREATER);break;

            //division or comment 
            case '/': 
                if(match('/')){
                    //if we dont fine EOF or a newline, keep movin forward
                    while(peek() != '\n' && !isAtEnd()) advance();
                }else{
                    addToken(SLASH);
                }break;

                //shit we dont care about 
                case ' ':
                case '\r':
                case '\t': 
                    //ignore whitespaces
                    break;
                case '\n':
                    //ignore newline, but advance line counter
                    line++;
                    break;

                //string literals 
                case '"': string(); break;

            default: 
                if(isDigit(c)){
                    number();
                }else if(isAlpha(c)){
                    identifier();
                }else{
                //if we dont recognize some shit, call error
                //we keep going to check if there are any more errors, but the hasError flag is set so none of this will be executed
                Lox.error(line, "Unxepected Character");
                }
        }
    }

   

    private void identifier() {
        while(isAlphaNumeric(peek())) advance();
        String text = source.substring(start, current);
        TokenType type = keywords.get(text);
        if(type == null) type = IDENTIFIER;
        addToken(type);
    }

    private boolean isAlphaNumeric(char c){
        return isAlpha(c) || isDigit(c);
    }

    private boolean isAlpha(char c) {
        return (c >= 'a' && c <= 'z') || 
        (c >= 'A' && c <= 'Z') ||
        c == '_';
    }

    private void number() {
        //if next value is a digit, keep going
        while(isDigit(peek())) advance();
        //if next thing is a . and the value after that is a number, then we have a decimal 
        if(peek() == '.' && isDigit(peekNext())){
            advance();
            while(isDigit(peek())) advance();
        }
        addToken(NUMBER, Double.parseDouble(source.substring(start, current)));
    }

    private char peekNext() {
        //if next value is past end of file, 
        if (current + 1 >= source.length()) return '\0';
        //if not return the next value
        return source.charAt(current + 1);
    }

    private boolean isDigit(char c) {
        return c >= '0' && c <= '9';
    }

    private void string() {
        while(peek() != '"' && !isAtEnd()){
            if (peek() == '\n') line++;
            advance();
        }

        if(isAtEnd()){
            Lox.error(line, "Unterminated string");
        }
        advance(); //at the closing "

        //start will be the starting " and curtrent will be the ending ", so get string inside those
        String value = source.substring(start + 1, current -1);
        addToken(STRING, value);
    }

    private boolean match(char c) {
        if(isAtEnd()) return false;
        //advance() moves current ahead by one. 
        if(source.charAt(current) != c) return false;
        current++;
        return true;
    }

    //lookahead
    private char peek() {
        //if were at the end, just return null character
        if(isAtEnd()) return '\0';
        //otherwise, return the next char
        return source.charAt(current);
    }

    //adds token with null as object literal
    private void addToken(TokenType type) {
        addToken(type, null);
    }

    //overrides the addToken but allows to pass in object literal
    //adds token to the class list of tokens
    private void addToken(TokenType type, Object literal) {
        String text = source.substring(start, current);
        tokens.add(new Token(type, text, literal, line));
    }

    private char advance() {
        return source.charAt(current++);
    }

    private boolean isAtEnd(){
        return current >= source.length();
    }
}
