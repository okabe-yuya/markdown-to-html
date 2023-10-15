package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseBackquote(token *lexer.Token) (*Node, *lexer.Token) {
	count, curToken := backquoteCount(token)
	content := ""
	for !exepct(curToken, lexer.RESERVED, "`") {
		content += curToken.Value
		curToken = curToken.Next
	}

	curToken = seek(curToken, count)
	node := NewNode(ND_BACKQUOTE, content, count, 0)
	return node, curToken
}

func backquoteCount(token *lexer.Token) (int, *lexer.Token) {
	freq := 0
	curToken := token
	for curToken.Kind == lexer.RESERVED && curToken.Value == "`" {
		freq++
		curToken = curToken.Next
	}
	return freq, curToken
}
