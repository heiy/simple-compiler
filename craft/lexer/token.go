package lexer

type TokenType int
type TokenState int

const (
	INIT TokenType = iota
	Plus
	Minus
	Star
	Slash
	PlusEQ
	GT
	GE
	Less
	LessEQ
	MinusEQ
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

	SemiColon
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
)

type Token struct {
	TokenType  TokenType
	TokenState TokenState
	Value      []rune
}

type TokenReader struct {
	Tokens []Token
	Pos    int
}

func NewTokenReader(tokens []Token) *TokenReader {
	return &TokenReader{
		Tokens: tokens,
		Pos:    0,
	}
}

func (r *TokenReader) Read() Token {
	if r.Pos < len(r.Tokens)-1 {
		r.Pos++
		return r.Tokens[r.Pos]
	}
	return Token{}
}

func (r *TokenReader) Peek() Token {
	if r.Pos < len(r.Tokens) {
		return r.Tokens[r.Pos]
	}
	return Token{}
}
