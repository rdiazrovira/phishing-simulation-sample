package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    "github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
    app.Name = "Phishing simulator CLI"
    app.Usage = "A CLI for simulating a Phishing"
    app.Author = "Rodolfo"
    app.Version = "1.0.0"
}

func commands() {
    app.Commands = []cli.Command{
        {
            Name:    "Replacement",
            Aliases: []string{"rpc"},
            Usage:   "Create a copy of the HTML page of any URL",

            Action: func(c *cli.Context) {
                if isValid(c) == "" {
                    err := createHTMLPageCopy(c.Args().Get(0), c.Args().Get(1))
                    if err != nil {
                        panic(err)
                    }

                    return
                }
                fmt.Println(isValid(c))
            },
        },
    }
}

func isValid(c *cli.Context) (err string) {
    if c.NArg() != 2 {
        err = "Err: bad input: two params are required: filepath (destination) and URL (source)"
    }
    return err
}

func createHTMLPageCopy(filename, url string) (err error) {
    fmt.Println("Downloading ", url, " to ", filename)

    resp, err := http.Get(url)
    if err != nil {
        return
    }
    defer resp.Body.Close()

    /*doc, err := goquery.NewDocumentFromResponse(resp)
    if err != nil {
        log.Fatal(err)
    }

    // Find the review items
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        s.SetAttr("href", "abcdefg")
    })*/

    file, err := os.Create(filename)
    if err != nil {
        return
    }
    defer file.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        return
    }

    fmt.Println(filename + " copied!")

    // create from a file
    /*one, err := os.Open("example.html")
    if err != nil {
        log.Fatal(err)
    }
    defer one.Close()
    doc, err := goquery.NewDocumentFromReader(one)
    if err != nil {
        log.Fatal(err)
    }

    // Find the review items
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        s.SetAttr("href", "abcdefg")
    })*/

    return
}

func main() {
    info()
    commands()

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

