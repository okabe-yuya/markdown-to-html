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
		case lexer.PLAIN_TEXT:
			node, curToken = parseText(curToken)
			nodes = append(nodes, node)
		case lexer.SEPARATE:
			curToken = curToken.Next
			continue
		}
	}
	return nodes, nil
}

func reserved(token *lexer.Token) (*Node, *lexer.Token) {
	switch token.Value {
	case "#":
		return parseHeader(token)
	case "-":
		return parserList(token)
	default:
		panic(1)
	}
}
