package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func parseItalic(token *lexer.Token) (*Node, *lexer.Token) {
	curToken := token
	value := ""

	if expectNext(curToken, lexer.RESERVED, "_") {
		curToken = seek(curToken, 2)

		for curToken.Kind != lexer.SEPARATE {
			if exepct(curToken, lexer.RESERVED, "_") && expectNext(curToken, lexer.RESERVED, "_") {
				node := NewNode(ND_ITALIC, value, 1, 0)
				curToken = seek(curToken, 2)
				return node, curToken
			}
			value += curToken.Value
			curToken = curToken.Next
		}

		// curTokenのvalueに相乗り
		curToken.Value = "__" + value
		return nil, curToken
	} else {
		return nil, curToken
	}
}
