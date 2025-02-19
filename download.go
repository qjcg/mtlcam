package main

import (
	"log"
	"os"
	"path"

	"github.com/gocolly/colly/v2"
)

// GET and return contents from URL.
func download(URL string) {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		log.Println(e.Attr("href"))
	})

	err := c.Visit(URL)
	if err != nil {
		log.Fatal(err)
	}
}

// Write slice of bytes to disk.
func saveFile(data []byte, filename, dir string) {
	err := os.WriteFile(path.Join(dir, filename), data, 0o644)
	if err != nil {
		log.Fatal("Couldn't create file -- ", err)
	}
}
