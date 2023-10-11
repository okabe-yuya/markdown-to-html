package parser

import (
	"github.com/okabe-yuya/makrdown-to-html/lexer"
)

func _parseText(token *lexer.Token) (*Node, *lexer.Token) {
	node := NewNode(ND_VALUE, token.Value, 0, 0, nil)
	curToken := token.Next

	for curToken != nil && curToken.Kind == lexer.PLAIN_TEXT {
		node.Value += curToken.Value
		curToken = curToken.Next
	}
	return node, curToken
}

func parseText(token *lexer.Token) (*Node, *lexer.Token) {
	var node_ *Node
	curToken := token
	node := NewNode(ND_VALUE, "", 0, 0, nil)

L:
	for curToken != nil {
		switch curToken.Kind {
		case lexer.RESERVED:
			if curToken.Value == "*" {
				node_, curToken = parseWeight(curToken)
				node.Nest = node_
			} else {
				break L
			}
		case lexer.PLAIN_TEXT:
			if node.Nest == nil {
				node.Value += curToken.Value
				curToken = curToken.Next
			} else {
				curNode := node.Nest
				for curNode.Nest != nil {
					curNode = curNode.Nest
				}
				node_, curToken = parseText(curToken)
				curNode.Nest = node_
			}
		default:
			break L
		}
	}
	return node, curToken
}
