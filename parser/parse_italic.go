package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseItalic(token *lexer.Token) (*Node, *lexer.Token) {
	curToken := token
	value := ""

	if expectNext(curToken, lexer.RESERVED, "_") {
		curToken = seek(curToken, 2)
		for curToken.Kind == lexer.PLAIN_TEXT {
			value += curToken.Value
			curToken = curToken.Next
		}
		// **分だけ進めておく
		curToken = seek(curToken, 2)
	} else {
		// ここにplain textのパースが必要
	}
	node := NewNode(ND_ITALIC, value, 1, 0)
	return node, curToken
}
