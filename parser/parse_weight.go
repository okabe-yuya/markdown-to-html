package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseWeight(token *lexer.Token) (*Node, *lexer.Token) {
	curToken := token
	value := ""

	if expectNext(curToken, lexer.RESERVED, "*") {
		curToken = seek(curToken, 2)

		for curToken.Kind != lexer.SEPARATE {
			if exepct(curToken, lexer.RESERVED, "*") && expectNext(curToken, lexer.RESERVED, "*") {
				node := NewNode(ND_WEIGHT, value, 1, 0)
				curToken = seek(curToken, 2)
				return node, curToken
			}
			value += curToken.Value
			curToken = curToken.Next
		}
		// curTokenのvalueに相乗り
		curToken.Value = "**" + value
		return nil, curToken
	} else {
		return nil, curToken
	}
}
