package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const filepath = "wawandco.html"
const newFilename = "copy.html"

func CreateFileCopy(file *os.File) (*os.File, error) {
	copy, err := os.Create(newFilename)
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

func getParamValueByName(r *http.Request, name string) string {
	arguments, ok := r.URL.Query()[name]

	if !ok || len(arguments[0]) < 1 {
		return ""
	}

	return string(arguments[0])
}

func replaceLinksFromFile(filepath, url string) (string, error) {

	if filepath == "" || url == "" {
		return "", errors.New("Empty params")
	}

	text, err := GetHtmlInTextFormatFromFile(filepath)
	if err != nil {
		return "", err
	}

	doc, err := ExtractHtmlFrom(text)
	if err != nil {
		return "", err
	}

	ReplaceLinks(doc, url)

	copy, _ := os.Create(newFilename)
	defer copy.Close()

	_, err = WriteHtmlIntoFile(doc, copy)
	return copy.Name(), err
}

func runServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := getParamValueByName(r, "filepath")
		url := getParamValueByName(r, "url")

		_, err := replaceLinksFromFile(filepath, url)
		if err == nil {
			fmt.Println("Links replaced.")
			http.ServeFile(w, r, newFilename)
			return
		}

		fmt.Println(err)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	if len(os.Args) <= 1 {
		runServer()
		return
	}

	if len(os.Args) != 3 {
		fmt.Println(message)
		return
	}

	_, err := replaceLinksFromFile(os.Args[1], os.Args[2])
	if err == nil {
		fmt.Println("Links replaced.")
	}
}
