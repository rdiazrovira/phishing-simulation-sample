package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func CreateFileCopy(file *os.File) (*os.File, error) {
	copy, err := os.Create("copy-" + file.Name())
	if err != nil {
		return file, err
	}
	defer copy.Close()

	_, err = io.Copy(copy, file)
	if err != nil {
		return file, err
	}

	return os.Open(copy.Name())
}

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(0)
	}
}

func modifyLinks(file *os.File) {
	doc, err := html.Parse(file)
	checkError(err)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(".")
					a.Val = "https://www.google.com"
				}
			}
		}

		for at := n.FirstChild; at != nil; at = at.NextSibling {
			f(at)
		}
	}

	f(doc)

	err = html.Render(file, doc)
	checkError(err)
}

func main() {
	filename := "wawandco.html"

	file, err := os.Open(filename)
	checkError(err)

	copy, err := CreateFileCopy(file)
	checkError(err)

	modifyLinks(copy)
}
