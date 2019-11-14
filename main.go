package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CreateFileCopy(file *os.File) (*os.File, error) {
	copy, err := os.Create("copy-" + file.Name())
	if err != nil {
		return file, err
	}

	_, err = io.Copy(copy, file)
	if err != nil {
		return file, err
	}

	return os.Open(copy.Name())
}

func ExtractHtmlFrom(text string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(text))
}

func ReplaceLinks(doc *goquery.Document, newLink string) {
	doc.Find("a").Each(func(i int, a *goquery.Selection) {
		a.SetAttr("href", newLink)
	})
}

func main() {
	text := `<body><a href="https://myhtml.com">My HTML</a></body>`

	doc, _ := ExtractHtmlFrom(text)

	ReplaceLinks(doc, "https://www.google.com")

	fmt.Println(doc.Html())
}
