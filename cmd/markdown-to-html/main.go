package main

import (
	"fmt"
	"os"

	"github.com/okabe-yuya/makrdown-to-html/adapter"
)

func main() {
	md, err := os.Open("static/md/link/single.md")
	defer md.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	html, err := adapter.GenerateHtml(md)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	WriteHtml(html, "verification.html")
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
