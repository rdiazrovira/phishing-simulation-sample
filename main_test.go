package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCreateFileCopy(t *testing.T) {
	filename := "wawandco.html"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = CreateFileCopy(file)
	if err != nil {
		log.Fatal(err)
	}

	f1, err := ioutil.ReadFile("wawandco.html")
	if err != nil {
		log.Fatal(err)
	}

	f2, err := ioutil.ReadFile("copy-wawandco.html")
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(f1, f2) {
		t.Error("Not equals")
	}
}
