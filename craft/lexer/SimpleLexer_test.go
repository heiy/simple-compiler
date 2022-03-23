package lexer

import (
	"log"
	"testing"
)

func TestLexer(t *testing.T) {
	str := "age >= 45"
	printLexer(str)

	str = "int age = 40"
	printLexer(str)

	str = "inta = 40"
	printLexer(str)

	str = "ina = 40"
	printLexer(str)

	str = "in = 40"
	printLexer(str)

	str = "i = 40"
	printLexer(str)

	str = "2 + 3 * 5"
	printLexer(str)
}

func printLexer(str string) {
	for _, token := range tokenize(str) {
		log.Printf("%d \t\t %d \t\t %s\n", token.TokenType, token.TokenState, string(token.Value))
	}
}
