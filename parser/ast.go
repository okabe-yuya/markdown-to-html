package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func Ast(token *lexer.Token) ([]*Node, error) {
	var node *Node
	var nodes []*Node
	curToken := token

	for {
		if curToken == nil || curToken.Next == nil {
			break
		}
		switch curToken.Kind {
		case lexer.RESERVED:
			node, curToken = reserved(curToken)
			nodes = append(nodes, node)
		case lexer.PLAIN_TEXT, lexer.BLANK:
			node, curToken = parseText(curToken)
			nodes = append(nodes, node)
		case lexer.SEPARATE:
			node = NewNode(ND_NEW_LINE, "\n", 0, 0)
			curToken = curToken.Next
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
}

func reserved(token *lexer.Token) (*Node, *lexer.Token) {
	switch token.Value {
	case "#":
		return parseHeader(token)
	case "-":
		return parseList(token)
	case "*", "_", "`", "[":
		return parseText(token)
	case ">":
		return parseQuote(token)
	default:
		panic(1)
	}
}
