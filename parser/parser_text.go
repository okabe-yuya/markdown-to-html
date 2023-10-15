package parser

import (
	"github.com/okabe-yuya/makrdown-to-html/lexer"
)

func parseText(token *lexer.Token) (*Node, *lexer.Token) {
	node, curToken := _parseText(token)
	if curToken != nil && curToken.Kind == lexer.SEPARATE {
		curToken = curToken.Next
	}
	return node, curToken
}

func _parseText(token *lexer.Token) (*Node, *lexer.Token) {
	var node_ *Node
	curToken := token
	node := NewNode(ND_VALUE, "", 0, 0)

L:
	for curToken != nil {
		switch curToken.Kind {
		case lexer.RESERVED:
			switch curToken.Value {
			case "*":
				node_, curToken = parseWeight(curToken)
				node.Nest = node_
			case "_":
				node_, curToken = parseItalic(curToken)
				node.Nest = node_
			case "`":
				node_, curToken = parseBackquote(curToken)
				node.Nest = node_
			case "[":
				node_, curToken = parseLink(curToken)
				node.Nest = node_
			default:
				break L
			}
		case lexer.PLAIN_TEXT, lexer.BLANK:
			if node.Nest == nil {
				node.Value += curToken.Value
				curToken = curToken.Next
			} else {
				curNode := node.Nest
				for curNode.Nest != nil {
					curNode = curNode.Nest
				}
				node_, curToken = _parseText(curToken)
				curNode.Nest = node_
			}
		default:
			break L
		}
	}
	return node, curToken
}
