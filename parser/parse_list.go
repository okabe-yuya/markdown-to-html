package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseList(token *lexer.Token) (*Node, *lexer.Token) {
	node, curToken := firstList(token)
	if curToken != nil && curToken.Next != nil && curToken.Kind == lexer.RESERVED && curToken.Value == "-" {
		var node_ *Node
		node_, curToken = parseList(curToken)
		node.Nest = node_
	}
	return node, curToken
}

func firstList(token *lexer.Token) (*Node, *lexer.Token) {
	txtNode, curToken := parseText(token.Next)
	node := NewNode(ND_LIST, "", 0, token.Depth, txtNode)

	return node, curToken
}
