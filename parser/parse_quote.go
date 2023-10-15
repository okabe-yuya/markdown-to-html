package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseQuote(token *lexer.Token) (*Node, *lexer.Token) {
	var node *Node

	depth, curToken := quoteDepth(token)
	if exepct(curToken, lexer.BLANK, " ") {
		node, curToken = firstQuote(curToken, depth)
		for curToken != nil && curToken.Kind == lexer.BLANK {
			curToken = curToken.Next
		}
		if exepct(curToken, lexer.RESERVED, ">") && curToken.Next != nil {
			var node_ *Node
			node_, curToken = parseQuote(curToken)
			node.Sub = node_
		}
	} else {
		node, curToken = parseText(curToken.Next)
		prefix := ""
		for i := 0; i < depth; i++ {
			prefix += ">"
		}
		node.Value = prefix + node.Value
	}

	return node, curToken
}

func firstQuote(token *lexer.Token, depth int) (*Node, *lexer.Token) {
	node := NewNode(ND_QUOTE, "", 0, depth)
	textNode, curToken := parseText(token.Next)
	node.Nest = textNode

	return node, curToken
}

func quoteDepth(token *lexer.Token) (int, *lexer.Token) {
	depth := 0
	curToken := token
	for exepct(curToken, lexer.RESERVED, ">") {
		depth++
		curToken = curToken.Next
	}
	return depth, curToken
}
