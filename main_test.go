package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var myHTML = `<body> <a href="https://myhtml.com">My HTML</a></body>`

func TestCreateFileCopy(t *testing.T) {

	file, _ := os.Open(fileName)

	_, err := CreateFileCopy(file)
	if err != nil {
		t.Error(err)
	}

	source, err := ioutil.ReadFile("copy-wawandco.html")
	if err != nil {
		t.Error("Not equals")
	}

	copy, err := ioutil.ReadFile("copy-wawandco.html")
	if err != nil {
		t.Error("Not equals")
	}

	if !bytes.Equal(source, copy) {
		t.Error("Not equals")
	}
}

func TestExtractHtmlFromFile(t *testing.T) {
	_, err := ExtractHtmlFrom(myHTML)
	if err != nil {
		t.Error(err)
	}
}

func TestReplaceLinks(t *testing.T) {
	doc, _ := ExtractHtmlFrom(myHTML)

	ReplaceLinks(doc, "https://www.google.com")

	count := 0
	doc.Find("a").Each(func(i int, a *goquery.Selection) {
		href, _ := a.Attr("href")
		if href == "https://www.google.com" {
			count++
		}
	})

	if count != 1 {
		t.Error("Links didn't change")
	}
}

func TestGetHtmlInTextFormat(t *testing.T) {
	text := GetHtmlInTextFormatFromFile(fileName)

	if len(text) == 0 {
		t.Error("Html could not be extracted in String format.")
	}
}

func TestWriteHtmlIntoFile(t *testing.T) {
	doc, _ := ExtractHtmlFrom(myHTML)

	copy, _ := os.Create("test.html")
	defer copy.Close()

	WriteHtmlIntoFile(doc, copy)

	text := GetHtmlInTextFormatFromFile(copy.Name())

	if len(text) == 0 {
		t.Error("Html could not be extracted in String format.")
	}

	count := 0
	doc.Find("a").Each(func(i int, a *goquery.Selection) {
		href, _ := a.Attr("href")
		if href == "https://myhtml.com" {
			count++
		}
	})

	if count != 1 {
		t.Error("Links didn't change")
	}

	err := os.Remove(copy.Name())
	if err != nil {
		t.Error(err)
	}
}
