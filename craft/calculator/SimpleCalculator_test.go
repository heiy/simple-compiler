package calculator

import (
	"code/PlayWithCompiler/craft/lexer"
	"testing"
)

func TestCalcultor(t *testing.T) {
	tokenReader := lexer.NewTokenReader(lexer.Tokenize("int age = b + 40"))
	node := intDeclare(tokenReader)
	printAST(node, "")

	evaluateAll("2 + 3 * 5")
	evaluateAll("2 + ")
	evaluateAll("2 + 3 + 4")
}
