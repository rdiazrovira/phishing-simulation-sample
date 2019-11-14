package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var myHTML = `<body><a href="https://myhtml.com">My HTML</a></body>`

func TestCreateFileCopy(t *testing.T) {

	file, _ := os.Open(fileName)

	_, err := CreateFileCopy(file)
	if err != nil {
		t.Error(err)
	}

	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error(err)
	}

	copy, err := ioutil.ReadFile("copy-" + fileName)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(source, copy) {
		t.Error("The files are differents")
	}
}

func TestExtractHtmlFromFile(t *testing.T) {
	_, err := ExtractHtmlFrom(myHTML)
	if err != nil {
		t.Error(err)
	}
}

func TestReplaceLinks(t *testing.T) {
	cases := [][]string{
		{"<body><a href=\"https://myhtml.com\">My HTML</a></body>", "https://www.google.com", "1"},
		{"<body><a href=\"https://myhtml.com\">My HTML</a></br><a href=\"https://example.com\">Example</a></body>", "https://www.google.com", "2"},
		{"<body><a href=\"https://myhtml.com\">My HTML</a></body>", "https://www.anotherexample.com", "1"},
	}

	for _, c := range cases {
		doc, _ := ExtractHtmlFrom(c[0])

		ReplaceLinks(doc, c[1])

		count := 0
		doc.Find("a").Each(func(i int, a *goquery.Selection) {
			href, _ := a.Attr("href")
			if href == c[1] {
				count++
			}
		})

		r, _ := strconv.Atoi(c[2])
		if count != r {
			t.Error("Links didn't change")
		}
	}

}

func TestGetHtmlInTextFormat(t *testing.T) {
	text, _ := GetHtmlInTextFormatFromFile(fileName)

	if len(text) == 0 {
		t.Error("Html could not be extracted in String format.")
	}
}

func TestWriteHtmlIntoFile(t *testing.T) {

	cases := [][]string{
		{"test1.html", "<body><a href=\"https://myhtml.com\">My HTML</a></body>", "1"},
		{"test2.html", "<body><a href=\"https://myhtml.com\">My HTML</a></br><a href=\"https://myhtml.com\">My HTML</a></body>", "2"},
		{"test3.html", "<body><a href=\"https://myhtml.com\">My HTML</a></br><a href=\"https://www.google.com\">Google</a></br><a href=\"https://myhtml.com\">My HTML</a></body>", "2"},
	}

	for index, c := range cases {
		copy, _ := os.Create(c[0])
		defer copy.Close()

		doc, _ := ExtractHtmlFrom(c[1])

		WriteHtmlIntoFile(doc, copy)

		count := 0
		doc.Find("a").Each(func(i int, a *goquery.Selection) {
			href, _ := a.Attr("href")
			if href == "https://myhtml.com" {
				count++
			}
		})

		r, _ := strconv.Atoi(c[2])
		if count != r {
			t.Error("Links didn't change for the case: " + strconv.Itoa(index+1))
		}

		err := os.Remove(copy.Name())
		if err != nil {
			t.Error(err)
		}
	}
}
