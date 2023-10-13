package main

import (
	"fmt"
	"os"

	"github.com/okabe-yuya/makrdown-to-html/generator"
	"github.com/okabe-yuya/makrdown-to-html/lexer"
	"github.com/okabe-yuya/makrdown-to-html/parser"
)

func main() {
	md, err := os.Open("static/md/list/weight.md")
	defer md.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	html, err := GenrateHtml(md)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(html)
	WriteHtml(html, "verification.html")
}

func GenrateHtml(f *os.File) (string, error) {
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

func WriteHtml(html, filename string) error {
	hf, err := os.Create(filename)
	if err != nil {
		return err
	}
	data := []byte(html)
	if _, err := hf.Write(data); err != nil {
		return err
	}
	return nil
}
