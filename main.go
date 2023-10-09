package main

import (
	"fmt"
	"os"

	"github.com/okabe-yuya/makrdown-to-html/block"
	htmlgen "github.com/okabe-yuya/makrdown-to-html/html-gen"
	"github.com/okabe-yuya/makrdown-to-html/token"
)

func main() {
	md, err := os.Open("static/md/empty.md")
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
	WriteHtml(html, "parsed.html")
}

func GenrateHtml(f *os.File) (string, error) {
	token, err := token.Generate(f)
	if err != nil {
		return "", err
	}

	blocks, err := block.Generate(token)
	if err != nil {
		return "", err
	}
	return htmlgen.Exec(blocks), nil
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
