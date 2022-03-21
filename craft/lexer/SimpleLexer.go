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
		if string(ch) == "i" {
			token.TokenType = Id_int1
		} else {
			token.TokenType = Id
		}
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
	case ch == '=':
		token.TokenType = Assignment
		break
	case ch == '+':
		token.TokenType = Plus
		break
	case ch == '-':
		token.TokenType = Minus
		break
	case ch == '*':
		token.TokenType = Star
		break
	case ch == '/':
		token.TokenType = Slash
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
	tokens = []Token{}
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
		case Id_int1:
			if string(ch) == "n" {
				token.Value = append(token.Value, ch)
				token.TokenType = Id_int2
			} else if isAlpha(ch) || (isDigit(ch)) {
				token.TokenType = Id
			} else {
				tokens = append(tokens, token)
				token = initToken(ch)
			}
			break
		case Id_int2:
			if string(ch) == "t" {
				token.Value = append(token.Value, ch)
				token.TokenType = Id_int3
			} else if isAlpha(ch) || (isDigit(ch)) {
				token.TokenType = Id
			} else {
				tokens = append(tokens, token)
				token = initToken(ch)
			}
			break
		case Id_int3:
			if isBlank(ch) {
				tokens = append(tokens, token)
				token.TokenType = Int
				token = initToken(ch)
			} else {
				token.TokenType = Id
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
		case Assignment:
			tokens = append(tokens, token)
			token = initToken(ch)
		case Plus:
			tokens = append(tokens, token)
			token = initToken(ch)
		case Minus:
			tokens = append(tokens, token)
			token = initToken(ch)
		case Star:
			tokens = append(tokens, token)
			token = initToken(ch)
		case Slash:
			tokens = append(tokens, token)
			token = initToken(ch)
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
