package main

import (
	"fmt"
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

func GetHtmlInTextFormatFromFile(src string) (string, error) {
	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func WriteHtmlIntoFile(doc *goquery.Document, file *os.File) (int, error) {
	result, _ := doc.Html()
	return file.WriteString(result)
}

const message = `
My Cli is a tool for replacing the links <a> of a page html.

Usage:

	./main [arguments]
	
The arguments should be: 
	[path]	Path to a HTML file (1st).
	[URL]	Replacement URL (2nd).

With those arguments it will create a copy of that HTML file (myfile.html) 
with the same content that the original except for all the links <a>
replaced by [URL].`

func main() {
	if len(os.Args) != 3 {
		fmt.Println(message)
		return
	}

	text, err := GetHtmlInTextFormatFromFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := ExtractHtmlFrom(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	ReplaceLinks(doc, os.Args[2])

	copy, _ := os.Create("copy-" + os.Args[1])
	defer copy.Close()

	_, err = WriteHtmlIntoFile(doc, copy)
	if err == nil {
		fmt.Println("Links replaced.")
	}
}
