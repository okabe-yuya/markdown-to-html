package parser

import (
	"github.com/okabe-yuya/makrdown-to-html/lexer"
)

func parseHeader(token *lexer.Token) (*Node, *lexer.Token) {
	var value string
	var node *Node

	level, curToken := headerLevel(token)
	if exepct(curToken, lexer.BLANK, " ") {
		value, curToken = headerValue(curToken.Next)
		node := NewNode(ND_HEADER, value, level, 0)
		return node, curToken
	} else {
		node, curToken = parseText(curToken)
		prefix := ""
		for i := 0; i < level; i++ {
			prefix += "#"
		}
		node.Value = prefix + node.Value
		return node, curToken
	}
}

func headerLevel(token *lexer.Token) (int, *lexer.Token) {
	level := 0
	curToken := token
	for curToken.Kind == lexer.RESERVED && curToken.Value == "#" {
		level++
		curToken = curToken.Next
	}
	return level, curToken
}

func headerValue(token *lexer.Token) (string, *lexer.Token) {
	value := ""
	curToken := token
	for {
		if curToken.Kind == lexer.SEPARATE {
			curToken = curToken.Next
			break
		}
		value += curToken.Value
		curToken = curToken.Next
	}
	return value, curToken
}
