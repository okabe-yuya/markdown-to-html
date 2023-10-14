package parser

import "github.com/okabe-yuya/makrdown-to-html/lexer"

func exepct(token *lexer.Token, kind lexer.TokenKind, ept string) bool {
	if token == nil {
		return false
	}
	return token.Kind == kind && token.Value == ept
}

func expectNext(token *lexer.Token, kind lexer.TokenKind, expect string) bool {
	if token.Next == nil {
		return false
	}

	next := token.Next
	return next.Kind == kind && next.Value == expect
}

func seek(token *lexer.Token, n int) *lexer.Token {
	res := token
	for i := 0; i < n; i++ {
		res = res.Next
		if res == nil {
			break
		}
	}
	return res
}
