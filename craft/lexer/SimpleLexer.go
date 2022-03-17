package lexer

import (
	"log"
	"regexp"
)

var tokens []Token

func initToken(ch rune) Token {
	token := Token{
		TokenType: Initial,
		Value:     []rune{ch},
	}

	switch {
	case isAlpha(ch):
		token.TokenType = Id
		break
	case isDigit(ch):
		token.TokenType = IntLiteral
		break
	case ch == '>':
		token.TokenType = GT
		break
	case ch == '<':
		token.TokenType = LT
		break
	default:
	}

	return token
}

func isAlpha(ch rune) bool {
	ok, err := regexp.MatchString(`[a-zA-Z]`, string(ch))

	if err != nil {
		log.Print(err)
	}

	return ok
}

func isDigit(ch rune) bool {
	ok, err := regexp.MatchString(`[0-9]`, string(ch))

	if err != nil {
		log.Print(err)
	}

	return ok
}

func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func tokenize(str string) []Token {
	var token Token

	for index, ch := range str {
		switch token.TokenType {
		case Initial:
			token = initToken(ch)
			break
		case Id:
			if isAlpha(ch) || (isDigit(ch)) {
				token.Value = append(token.Value, ch)
			} else {
				tokens = append(tokens, token)
				token = initToken(ch)
			}
			break
		case GT:
			if ch == '=' {
				token.TokenType = GE
				token.Value = append(token.Value, ch)
			} else {
				tokens = append(tokens, token)
				token = initToken(ch)
			}
			break
		case GE:
			tokens = append(tokens, token)
			token = initToken(ch)
			break
		case IntLiteral:
			if isDigit(ch) {
				token.Value = append(token.Value, ch)
			} else {
				tokens = append(tokens, token)
				token = initToken(ch)
			}
			break
		default:
		}

		if index == len(str)-1 {
			tokens = append(tokens, token)
		}
	}

	return tokens
}
