package parser

import (
	"github.com/okabe-yuya/makrdown-to-html/lexer"
)

func parseImage(token *lexer.Token) (*Node, *lexer.Token) {
	node, curToken := parseLink(token.Next)
	node.Kind = ND_IMAGE

	return node, curToken
}
