package lexer

type TokenType int

const (
	Initial TokenType = iota

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

	GT
	LT
	GE
	LE

	Assignment

	Plus
	Minus
	Star
	Slash

	SemiColon
	LeftParen
	RightParen

	IntLiteral
)

type Token struct {
	TokenType TokenType
	Value     []rune
}
