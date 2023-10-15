package adapter

import (
	"os"

	"github.com/okabe-yuya/makrdown-to-html/generator"
	"github.com/okabe-yuya/makrdown-to-html/lexer"
	"github.com/okabe-yuya/makrdown-to-html/parser"
)

func GenerateHtml(f *os.File) (string, error) {
	token, err := lexer.Tokenize(f)
	if err != nil {
		return "", err
	}

	ast, err := parser.Ast(token)
	if err != nil {
		return "", err
	}
	return generator.Html(ast), nil
}
