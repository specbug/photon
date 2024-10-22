package token

import (
	"strconv"
)

const (
	// Keywords: Variable declarations, functions, and control flow
	KW_LET    = "LET"    // Variable declaration
	KW_CONST  = "CONST"  // Constant declaration
	KW_FUNC   = "FUNC"   // Function declaration
	KW_RETURN = "RETURN" // Return statement
	KW_EXIT   = "EXIT"   // Return statement

	// Keywords: Conditionals and loops
	KW_IF       = "IF"       // If statement
	KW_ELSE     = "ELSE"     // Else statement
	KW_ELIF     = "ELIF"     // Else if statement
	KW_FOR      = "FOR"      // For loop
	KW_WHILE    = "WHILE"    // While loop
	KW_BREAK    = "BREAK"    // Break statement for loops
	KW_CONTINUE = "CONTINUE" // Continue statement for loops
	KW_PRINT    = "PRINT"    // Print statement

	// Keywords: Object-oriented programming
	KW_CLASS = "CLASS" // Class definition
	KW_SELF  = "SELF"  // Self-reference within a class

	// Keywords: Boolean values
	KW_TRUE  = "TRUE"  // Boolean true
	KW_FALSE = "FALSE" // Boolean false
	KW_NIL   = "NIL"   // Null/nil value

	// Symbols: Arithmetic operators
	SYM_PLUS  = "PLUS"  // +
	SYM_MINUS = "MINUS" // -
	SYM_MUL   = "MUL"   // *
	SYM_DIV   = "DIV"   // /
	SYM_MOD   = "MOD"   // %

	// Symbols: Comparison operators
	SYM_EQ  = "EQ"  // ==
	SYM_NEQ = "NEQ" // !=
	SYM_GT  = "GT"  // >
	SYM_GTE = "GTE" // >=
	SYM_LT  = "LT"  // <
	SYM_LTE = "LTE" // <=

	// Symbols: Logical operators
	SYM_AND = "AND" // &&
	SYM_OR  = "OR"  // ||
	SYM_NOT = "NOT" // !

	// Symbols: Bitwise operators
	SYM_AMPERSAND = "AMPERSAND" // &
	SYM_PIPE      = "PIPE"      // |
	SYM_CARET     = "CARET"     // ^

	// Symbols: Parentheses, braces, brackets
	SYM_LPAREN   = "LPAREN"   // (
	SYM_RPAREN   = "RPAREN"   // )
	SYM_LBRACE   = "LBRACE"   // {
	SYM_RBRACE   = "RBRACE"   // }
	SYM_LBRACKET = "LBRACKET" // [
	SYM_RBRACKET = "RBRACKET" // ]

	// Symbols: Miscellaneous
	SYM_SEMI       = "SEMI"       // ;
	SYM_COMMA      = "COMMA"      // ,
	SYM_COLON      = "COLON"      // :
	SYM_DOT        = "DOT"        // .
	SYM_DQUOTE     = "DQUOTE"     // "
	SYM_SQUOTE     = "SQUOTE"     // '
	SYM_BACKTICK   = "BACKTICK"   // `
	SYM_UNDERSCORE = "UNDERSCORE" // _
	SYM_ASSIGN     = "ASSIGN"     // =

	// Literals: Data types
	LIT_INT    = "INT_LTRL"    // Integer literal
	LIT_FLOAT  = "FLOAT_LTRL"  // Float literal
	LIT_STRING = "STRING_LTRL" // String literal
	LIT_CHAR   = "CHAR_LTRL"   // Character literal
	LIT_BOOL   = "BOOL_LTRL"   // Boolean literal (true/false)

	// Types
	TYPE_INT    = "INT"    // Integer
	TYPE_FLOAT  = "FLOAT"  // Float
	TYPE_STRING = "STRING" // String
	TYPE_CHAR   = "CHAR"   // Character
	TYPE_BOOL   = "BOOL"   // Boolean

	// Miscellaneous: Comments, newlines, and end of file
	T_COMMENT      = "COMMENT" // Comment
	T_NEWLINE      = "NEWLINE" // Newline
	T_EOF          = "EOF"     // End of file
	T_ZERO_MEASURE = ""        // empty string
	T_SPACE        = "SPACE"   // Space
	T_IDENT        = "IDENT"   // Identifier (variable name, function name, etc.)
)

type TType string
type Token struct {
	Type    TType
	Literal string
}
type TokenizeError struct {
	Message string
}

// Keyword map
var kwMap = map[string]TType{
	"exit":     KW_EXIT,
	"let":      KW_LET,
	"const":    KW_CONST,
	"func":     KW_FUNC,
	"return":   KW_RETURN,
	"if":       KW_IF,
	"else":     KW_ELSE,
	"elif":     KW_ELIF,
	"for":      KW_FOR,
	"while":    KW_WHILE,
	"break":    KW_BREAK,
	"continue": KW_CONTINUE,
	"print":    KW_PRINT,
	"class":    KW_CLASS,
	"self":     KW_SELF,
	"true":     KW_TRUE,
	"false":    KW_FALSE,
	"nil":      KW_NIL,
}

// Symbol map
var symMap = map[string]TType{
	";":  SYM_SEMI,
	",":  SYM_COMMA,
	"+":  SYM_PLUS,
	"-":  SYM_MINUS,
	"*":  SYM_MUL,
	"/":  SYM_DIV,
	"%":  SYM_MOD,
	"==": SYM_EQ,
	"!=": SYM_NEQ,
	">":  SYM_GT,
	">=": SYM_GTE,
	"<":  SYM_LT,
	"<=": SYM_LTE,
	"&&": SYM_AND,
	"||": SYM_OR,
	"!":  SYM_NOT,
	"&":  SYM_AMPERSAND,
	"|":  SYM_PIPE,
	"^":  SYM_CARET,
	"(":  SYM_LPAREN,
	")":  SYM_RPAREN,
	"{":  SYM_LBRACE,
	"}":  SYM_RBRACE,
	"[":  SYM_LBRACKET,
	"]":  SYM_RBRACKET,
	"=":  SYM_ASSIGN,
	":":  SYM_COLON,
	".":  SYM_DOT,
	"\"": SYM_DQUOTE,
	"'":  SYM_SQUOTE,
	"`":  SYM_BACKTICK,
	"_":  SYM_UNDERSCORE,
}

// Misc map
var miscMap = map[string]TType{
	"//":  T_COMMENT, // Assuming "#" is used for comments
	"\n":  T_NEWLINE,
	"EOF": T_EOF,
	"":    T_ZERO_MEASURE,
	" ":   T_SPACE,
}

// Type map
var typeMap = map[string]TType{
	"int":    TYPE_INT,
	"float":  TYPE_FLOAT,
	"string": TYPE_STRING,
	"char":   TYPE_CHAR,
	"bool":   TYPE_BOOL,
}

// global map
var LookupMap = map[string]TType{}

func init() {
	for k, v := range kwMap {
		LookupMap[k] = v
	}
	for k, v := range symMap {
		LookupMap[k] = v
	}
	for k, v := range typeMap {
		LookupMap[k] = v
	}
	for k, v := range miscMap {
		LookupMap[k] = v
	}
}

func (t *Token) String() string {
	switch t.Literal {
	case "":
		return "{" + string(t.Type) + "}"
	default:
		return "{" + string(t.Type) + ", " + t.Literal + "}"
	}
}

func (t *Token) New(tType TType, literal string) {
	t.Type = tType
	switch tType {
	case LIT_INT, LIT_FLOAT, LIT_STRING, LIT_CHAR, LIT_BOOL, T_IDENT:
		t.Literal = literal
	default:
		t.Literal = ""
	}
}

func litLookup(s string) (TType, bool) {
	if _, err := strconv.Atoi(s); err == nil {
		return LIT_INT, true
	}
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return LIT_FLOAT, true
	}
	return T_ZERO_MEASURE, false
}

func Lookup(s string) TType {
	if t, ok := LookupMap[s]; ok {
		return t
	}
	if t, ok := litLookup(s); ok {
		return t
	}
	return T_IDENT
}

func (te *TokenizeError) Error() string {
	return te.Message
}
