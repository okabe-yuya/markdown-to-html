package block

import "github.com/okabe-yuya/makrdown-to-html/token"

func Generate(t *token.Token) ([]*Block, error) {
	var block *Block
	var blocks []*Block
	curToken := t

	// Avoid empty markdown file
	for {
		if curToken == nil || curToken.Next == nil {
			break
		}
		switch curToken.Kind {
		case token.RESERVED:
			block, curToken = reserved(curToken)
			blocks = append(blocks, block)
		case token.PLAIN_TEXT:
			block, curToken = plaintext(curToken)
			blocks = append(blocks, block)
		case token.SEPARATE:
			curToken = curToken.Next
			continue
		}
	}
	return blocks, nil
}

func reserved(token *token.Token) (*Block, *token.Token) {
	switch token.Value {
	case "#":
		return header(token)
	case "-":
		return list(token)
	default:
		panic(1)
	}
}

func header(t *token.Token) (*Block, *token.Token) {
	hdrLvl := 0
	value := ""
	curToken := t

	for curToken.Kind == token.RESERVED && curToken.Value == "#" {
		hdrLvl++
		curToken = curToken.Next
	}

	if curToken.Kind == token.PLAIN_TEXT {
		for {
			if curToken.Kind == token.SEPARATE {
				curToken = curToken.Next
				break
			}
			value += curToken.Value
			curToken = curToken.Next
		}
	}

	res := NewBlock(ND_HEADER, value, hdrLvl, 0, nil)
	return res, curToken
}

func list(t *token.Token) (*Block, *token.Token) {
	curToken := t.Next
	value := ""
	for {
		if curToken.Kind == token.SEPARATE {
			curToken = curToken.Next
			break
		}
		value += curToken.Value
		curToken = curToken.Next
	}

	res := NewBlock(ND_LIST, value, 0, t.Depth, nil)
	if curToken != nil && curToken.Next != nil && curToken.Kind == token.RESERVED && curToken.Value == "-" {
		var block *Block
		block, curToken = list(curToken)
		res.Nest = block
	}
	return res, curToken
}

func plaintext(t *token.Token) (*Block, *token.Token) {
	res := NewBlock(ND_VALUE, t.Value, 0, 0, nil)
	nextToken := t.Next

	for nextToken != nil && nextToken.Kind == token.PLAIN_TEXT {

		res.Value += nextToken.Value
		nextToken = nextToken.Next
	}
	return res, nextToken
}
