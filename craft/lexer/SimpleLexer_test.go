package lexer

import (
	"log"
	"testing"
)

func TestLexer(t *testing.T) {
	log.Print(1)
	str := "age >= 45"
	for _, token := range tokenize(str) {
		log.Printf("%d \t\t %s\n", token.TokenType, string(token.Value))
	}
}
