package adapter

import (
	"bufio"
	"os"

	"github.com/okabe-yuya/makrdown-to-html/generator"
	"github.com/okabe-yuya/makrdown-to-html/lexer"
	"github.com/okabe-yuya/makrdown-to-html/parser"
)

func GenerateHtmlFromFile(f *os.File) (string, error) {
	br := bufio.NewReader(f)
	return GenerateHtml(br)
}

func GenerateHtml(br *bufio.Reader) (string, error) {
	token, err := lexer.Tokenize(br)
	if err != nil {
		return "", err
	}

	ast, err := parser.Ast(token)
	if err != nil {
		return "", err
	}
	return generator.Html(ast), nil
}
