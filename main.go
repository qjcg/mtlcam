// Download Montreal traffic camera images
// Uses open data from the city of Montreal at:
//   - http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere
// See in action via city map application at @ http://ville.montreal.qc.ca/circulation/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/fatih/color"
)

const (
	// See http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere
	URLGeoJSON string = "http://ville.montreal.qc.ca/circulation/sites/ville.montreal.qc.ca.circulation/files/cameras-de-circulation.json"
	URLAbout   string = "http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere"
	URLBase    string = "http://www1.ville.montreal.qc.ca/Circulation-Cameras/GEN%03d.jpeg"
)

var (
	// functions to colorize strings for use in sprintf-style functions
	bluef  = color.New(color.FgBlue, color.Bold).SprintFunc()
	greenf = color.New(color.FgGreen, color.Bold).SprintFunc()
)

// Create date/timestampped subdirectories for saving images
func MakeTimeStampDir(parentDir string) string {
	timeStampDir := time.Now().Format(path.Join(parentDir, "060102/150405"))
	err := os.MkdirAll(timeStampDir, 0755)
	if err != nil {
		log.Fatalf("Couldn't create image directory: %s\n", err)
	}

	return timeStampDir
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\n%s: Download Montreal traffic camera images\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Data source: %s\n\n", URLAbout)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}

	concurrency := flag.Int("c", 90, "max concurrent downloads")
	parentDir := flag.String("d", "images", "parent directory for downloaded files")
	flag.Parse()

	// timestamped directory to store images in
	tsDir := MakeTimeStampDir(*parentDir)

	// max number of workers (counting semaphore)
	workers := make(chan int, *concurrency)

	var fc FeatureCollection
	geoJSON := download(URLGeoJSON)
	if err := json.Unmarshal(geoJSON, &fc); err != nil {
		log.Printf("Error unmarshalling GeoJSON data: %v\n", err)
	}

	// wg waits for all downloads to complete
	var wg sync.WaitGroup
	for _, f := range fc.Features {
		imgURL := f.Properties.URLImageEnDirect
		if imgURL == "" {
			continue
		}

		// use one counting semaphore slot (when full, will block until slot free)
		workers <- 1
		wg.Add(1)

		imgFile := path.Base(imgURL)
		if imgFile == "" {
			log.Fatalf("Couldn't derive filename: %s\n", imgURL)
		}

		go func(URL, file, dir string) {
			defer wg.Done()
			img := download(URL)
			saveFile(img, file, dir)

			// when download finished, free one slot in workers
			<-workers
		}(imgURL, imgFile, tsDir)
	}

	wg.Wait()
}
