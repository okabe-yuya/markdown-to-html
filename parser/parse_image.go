package parser

import (
	"github.com/okabe-yuya/makrdown-to-html/lexer"
)

func parseImage(token *lexer.Token) (*Node, *lexer.Token) {
	node, curToken := parseLink(token.Next)
	if node.Kind == ND_LINK {
		node.Kind = ND_IMAGE
	} else {
		node.Value = "!" + node.Value
	}
	return node, curToken
}
