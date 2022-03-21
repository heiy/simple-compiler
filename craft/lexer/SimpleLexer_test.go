package lexer

import (
	"log"
	"testing"
)

func TestLexer(t *testing.T) {
	str := "age >= 45"
	printLexer(str)

	str1 := "int age = 40"
	printLexer(str1)

	str2 := "2 + 3 * 5"
	printLexer(str2)
}

func printLexer(str string) {
	for _, token := range tokenize(str) {
		log.Printf("%d \t\t %s\n", token.TokenType, string(token.Value))
	}
}
