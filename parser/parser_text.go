package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseText(token *lexer.Token) (*Node, *lexer.Token) {
	node := NewNode(ND_VALUE, token.Value, 0, 0, nil)
	curToken := token.Next

	for curToken != nil && curToken.Kind == lexer.PLAIN_TEXT {
		node.Value += curToken.Value
		curToken = curToken.Next
	}
	return node, curToken
}
