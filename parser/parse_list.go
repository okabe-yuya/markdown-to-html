package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseList(token *lexer.Token) (*Node, *lexer.Token) {
	var node *Node
	var curToken *lexer.Token

	if expectNext(token, lexer.BLANK, " ") {
		node, curToken = firstList(token.Next, token.Depth)
		for curToken != nil && curToken.Kind == lexer.BLANK {
			curToken = curToken.Next
		}
		if curToken != nil && curToken.Next != nil && curToken.Kind == lexer.RESERVED && curToken.Value == "-" {
			var node_ *Node
			node_, curToken = parseList(curToken)
			node.Sub = node_
		}
	} else {
		node, curToken = parseText(token.Next)
		node.Value = "-" + node.Value
	}
	return node, curToken
}

func firstList(token *lexer.Token, depth int) (*Node, *lexer.Token) {
	node := NewNode(ND_LIST, "", 0, depth)
	textNode, curToken := parseText(token.Next)
	node.Nest = textNode

	return node, curToken
}
