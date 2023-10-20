package parser

import (
	"fmt"

	"github.com/okabe-yuya/makrdown-to-html/lexer"
)

func parseLink(token *lexer.Token) (*Node, *lexer.Token) {
	var node *Node
	display, link := "", ""
	curToken := token.Next

	for curToken != nil && !exepct(curToken, lexer.RESERVED, "]") {
		display += curToken.Value
		curToken = curToken.Next
		if exepct(curToken, lexer.SEPARATE, "\n") {
			node = NewNode(ND_VALUE, "["+display, 0, 0)
			return node, curToken
		}
	}
	if expectNext(curToken, lexer.RESERVED, "(") {
		curToken = seek(curToken, 2) // ]と(の2つ分
		for curToken != nil && !exepct(curToken, lexer.RESERVED, ")") {
			link += curToken.Value
			curToken = curToken.Next
			if exepct(curToken, lexer.SEPARATE, "\n") {
				value := fmt.Sprintf("[%s](%s", display, link)
				node = NewNode(ND_VALUE, value, 0, 0)
				return node, curToken
			}
		}
		node = NewNode(ND_LINK, link, 0, 0)
		node.Nest = NewNode(ND_VALUE, display, 0, 0)
		curToken = curToken.Next // )分
	} else {
		node, curToken = parseText(curToken.Next)
		node.Value = fmt.Sprintf("[%s]%s", display, node.Value)
	}
	return node, curToken
}
