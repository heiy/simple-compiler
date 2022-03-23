package lexer

type TokenType int
type TokenState int

const (
	Plus TokenType = iota
	Minus
	Star
	Slash
	PPlus
	PlusEQ
	GT
	GE
	Less
	LessEQ
	MinusEQ
	MMinus
	StarEQ
	SlashEQ
	EQ
	Assign
	LeftParen
	RightParen

	IntLiteral
	DoubleLiteral

	Identifier

	INT
	DOUBLE
	BOOL
	NIL
	RUNE

	NULL
	EOF
	INIT
)

const (
	Initial TokenState = iota

	If
	Id_if1
	Id_if2
	Else
	Id_else1
	Id_else2
	Id_else3
	Id_else4
	Int
	Id_int1
	Id_int2
	Id_int3
	Id

	Assignment

	SemiColon
)

type Token struct {
	TokenType  TokenType
	TokenState TokenState
	Value      []rune
}

type TokenReader struct {
	Tokens []Token
	pos    int
}

func NewTokenReader(tokens []Token) *TokenReader {
	return &TokenReader{
		Tokens: tokens,
		pos:    0,
	}
}

func (r *TokenReader) Read() Token {
	if r.pos < len(r.Tokens) {
		r.pos++
		return r.Tokens[r.pos]
	}
	return Token{}
}

func (r *TokenReader) Peek() Token {
	if r.pos < len(r.Tokens) {
		return r.Tokens[r.pos]
	}
	return Token{}
}
