package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var fileName = "wawandco.html"

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

func GetHtmlInTextFormatFromFile(src string) string {
	bytes, _ := ioutil.ReadFile(src)
	return string(bytes)
}

func WriteHtmlIntoFile(doc *goquery.Document, file *os.File) {
	result, _ := doc.Html()
	file.WriteString(result)
}

func main() {
	text := GetHtmlInTextFormatFromFile(fileName)

	doc, _ := ExtractHtmlFrom(text)

	ReplaceLinks(doc, "https://www.google.com")

	copy, _ := os.Create("copy-" + fileName)
	defer copy.Close()

	WriteHtmlIntoFile(doc, copy)
}
