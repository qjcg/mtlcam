package main

import (
	"io/ioutil"
	"net/http"
	"path"

	log "github.com/sirupsen/logrus"
)

// GET and return contents from URL.
func download(URL string) []byte {
	resp, err := http.Get(URL)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("Couldn't GET file.")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Couldn't read reesponse body.")
	}

	return contents
}

// Write slice of bytes to disk.
func saveFile(data []byte, filename, dir string) {
	err := ioutil.WriteFile(path.Join(dir, filename), data, 0644)
	if err != nil {
		log.Fatal("Couldn't create file -- ", err)
	}
}
