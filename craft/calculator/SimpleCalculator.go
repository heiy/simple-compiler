package calculator

import (
	"code/PlayWithCompiler/craft/lexer"
	"log"
	"strconv"
)

func intDeclare(tokenReader *lexer.TokenReader) (node *Node) {
	token := tokenReader.Peek()
	node = NewNode(NUll, nil)

	if NodeType(token.TokenType) == NUll {
		return
	}

	if token.TokenType == lexer.INT {
		token = tokenReader.Read()

		if tokenReader.Peek().TokenType == lexer.Identifier {
			node.NodeType = IntDeclaration
			node.Value = token.Value

			token = tokenReader.Read()

			if NodeType(token.TokenType) == NUll {
				return
			}

			if token.TokenType == lexer.Assign {
				token = tokenReader.Read()

				child := additive(tokenReader)

				if child.GetType() != NUll {
					node.AddChild(child)
				}
			}
		}

		token = tokenReader.Peek()
		if token.TokenType == lexer.SemiColon {
			tokenReader.Read()
		}
	}

	return
}

func additive(tokenReader *lexer.TokenReader) (astNode ASTNode) {
	multiplicativeChild := multiplicative(tokenReader)
	astNode = multiplicativeChild
	token := tokenReader.Peek()

	if multiplicativeChild.GetType() == NUll {
		return
	}

	if NodeType(token.TokenType) == NUll {
		return
	}

	if token.TokenType == lexer.Plus || token.TokenType == lexer.Minus {
		operator := token.Value
		token = tokenReader.Read()

		additiveChild := additive(tokenReader)

		if additiveChild.GetType() == NUll {
			return
		}

		astNode = NewNode(Additive, operator)
		astNode.AddChild(multiplicativeChild)
		astNode.AddChild(additiveChild)
	}

	return
}

func multiplicative(tokenReader *lexer.TokenReader) (astNode ASTNode) {
	primaryChild := primary(tokenReader)
	astNode = primaryChild
	token := tokenReader.Peek()

	if primaryChild.GetType() == NUll {
		return
	}

	if NodeType(token.TokenType) == NUll {
		return
	}

	if token.TokenType == lexer.Star || token.TokenType == lexer.Slash {
		operator := token.Value
		token = tokenReader.Read()

		multiplicativeChild := multiplicative(tokenReader)

		if multiplicativeChild.GetType() == NUll {
			return
		}

		astNode = NewNode(Multiplicative, operator)
		astNode.AddChild(primaryChild)
		astNode.AddChild(multiplicativeChild)
	}

	return
}

func primary(tokenReader *lexer.TokenReader) (astNode ASTNode) {
	astNode = NewNode(NUll, nil)
	token := tokenReader.Peek()

	if NodeType(token.TokenType) == NUll {
		return
	}
	switch token.TokenType {
	case lexer.IntLiteral:
		astNode = NewNode(IntLiteral, token.Value)
		tokenReader.Read()
		break
	case lexer.Identifier:
		astNode = NewNode(Identifier, token.Value)
		tokenReader.Read()
		break
	case lexer.LeftParen:
		astNode = additive(tokenReader)
		tokenReader.Read()
		if astNode.GetType() == NUll {
			return
		}
		token = tokenReader.Peek()

		if NodeType(token.TokenType) == NUll {
			return
		}

		if token.TokenType == lexer.RightParen {
			tokenReader.Read()
		}
		break
	default:
	}

	return
}

func printAST(astNode ASTNode, indent string) {
	log.Printf("%s %s %s", indent, NodeType2Str[astNode.GetType()], string(astNode.GetValue()))
	for _, child := range astNode.GetChildren() {
		printAST(child, indent+"\t")
	}
}

func evaluate(astNode ASTNode, indent string) (ret int) {
	switch astNode.GetType() {
	case Program:
		for _, child := range astNode.GetChildren() {
			ret = evaluate(child, indent+"\t")
		}
		break
	case Additive:
		child1 := astNode.GetChildren()[0]
		value1 := evaluate(child1, indent+"\t")
		child2 := astNode.GetChildren()[1]
		value2 := evaluate(child2, indent+"\t")
		if string(astNode.GetValue()) == "+" {
			ret = value1 + value2
		} else {
			ret = value1 - value2
		}
		break
	case Multiplicative:
		child1 := astNode.GetChildren()[0]
		value1 := evaluate(child1, indent+"\t")
		child2 := astNode.GetChildren()[1]
		value2 := evaluate(child2, indent+"\t")
		if string(astNode.GetValue()) == "*" {
			ret = value1 * value2
		} else {
			ret = value1 / value2
		}
		break
	case IntLiteral:
		ret, _ = strconv.Atoi(string(astNode.GetValue()))
		break
	default:
	}
	return
}

func evaluateAll(str string) int {
	root := parse(str)
	printAST(root, "")
	return evaluate(root, "")

}

func parse(str string) ASTNode {
	tokens := lexer.Tokenize(str)
	tokenReader := lexer.NewTokenReader(tokens)
	return prog(tokenReader)

}

func prog(tokenReader *lexer.TokenReader) *Node {
	node := NewNode(Program, []rune("Calculator"))

	child := additive(tokenReader)
	if child.GetType() != NUll {
		node.AddChild(child)
	}
	return node
}
