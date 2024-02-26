package token

type TokenType string


type Token struct {
    Type    TokenType
    Literal string
}


const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT   = "INT"   // 1343456
    STRING = "STRING" // "foobar"

    // Operators
    ASSIGN   = "="
    PLUS     = "+"

    BANG     = "!"
    MINUS    = "-"
    SLASH    = "/"
    ASTERISK = "*"

    LT = "<"
    GT = ">"

    // Delimiters
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"
    LBRACKET = "["
    RBRACKET = "]"
    COLON = ":"

    // Keywords
    FUNCTION = "FUNCTION"
    LET      = "LET"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"

    // comparison Operators
    EQ = "=="
    NOT_EQ = "!="

)
