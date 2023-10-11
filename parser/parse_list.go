package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parserList(token *lexer.Token) (*Node, *lexer.Token) {
	node, curToken := firstList(token)
	if curToken != nil && curToken.Next != nil && curToken.Kind == lexer.RESERVED && curToken.Value == "-" {
		var node_ *Node
		node_, curToken = parserList(curToken)
		node.Nest = node_
	}
	return node, curToken
}

func firstList(token *lexer.Token) (*Node, *lexer.Token) {
	curToken := token.Next
	value := ""
	for {
		if curToken.Kind == lexer.SEPARATE {
			curToken = curToken.Next
			break
		}
		value += curToken.Value
		curToken = curToken.Next
	}

	node := NewNode(ND_LIST, value, 0, token.Depth, nil)
	return node, curToken
}
