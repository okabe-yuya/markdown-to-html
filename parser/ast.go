package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func Generate(token *lexer.Token) ([]*Node, error) {
	var node *Node
	var nodes []*Node
	curToken := token

	for {
		if curToken == nil || curToken.Next == nil {
			break
		}
		switch curToken.Kind {
		case lexer.RESERVED:
			node, curToken = reserved(curToken)
			nodes = append(nodes, node)
		case lexer.PLAIN_TEXT:
			node, curToken = plaintext(curToken)
			nodes = append(nodes, node)
		case lexer.SEPARATE:
			curToken = curToken.Next
			continue
		}
	}
	return nodes, nil
}

func reserved(token *lexer.Token) (*Node, *lexer.Token) {
	switch token.Value {
	case "#":
		return header(token)
	case "-":
		return list(token)
	default:
		panic(1)
	}
}

func header(token *lexer.Token) (*Node, *lexer.Token) {
	hdrLvl := 0
	value := ""
	curToken := token

	for curToken.Kind == lexer.RESERVED && curToken.Value == "#" {
		hdrLvl++
		curToken = curToken.Next
	}

	if curToken.Kind == lexer.PLAIN_TEXT {
		for {
			if curToken.Kind == lexer.SEPARATE {
				curToken = curToken.Next
				break
			}
			value += curToken.Value
			curToken = curToken.Next
		}
	}

	res := NewNode(ND_HEADER, value, hdrLvl, 0, nil)
	return res, curToken
}

func list(t *lexer.Token) (*Node, *lexer.Token) {
	curToken := t.Next
	value := ""
	for {
		if curToken.Kind == lexer.SEPARATE {
			curToken = curToken.Next
			break
		}
		value += curToken.Value
		curToken = curToken.Next
	}

	res := NewNode(ND_LIST, value, 0, t.Depth, nil)
	if curToken != nil && curToken.Next != nil && curToken.Kind == lexer.RESERVED && curToken.Value == "-" {
		var block *Node
		block, curToken = list(curToken)
		res.Nest = block
	}
	return res, curToken
}

func plaintext(t *lexer.Token) (*Node, *lexer.Token) {
	res := NewNode(ND_VALUE, t.Value, 0, 0, nil)
	nextToken := t.Next

	for nextToken != nil && nextToken.Kind == lexer.PLAIN_TEXT {

		res.Value += nextToken.Value
		nextToken = nextToken.Next
	}
	return res, nextToken
}
