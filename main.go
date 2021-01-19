// Download Montreal traffic camera images
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/gosuri/uiprogress"
	log "github.com/sirupsen/logrus"
)

const (
	URLGeoJSON string = "https://ville.montreal.qc.ca/circulation/sites/ville.montreal.qc.ca.circulation/files/cameras-de-circulation.json"
	URLAbout   string = "https://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere"
)

var reDigits = regexp.MustCompile(`\d+`)

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

	concurrency := flag.Int("c", 20, "max concurrent downloads")
	debug := flag.Bool("d", false, "print debug messages")
	parentDir := flag.String("p", "images", "parent directory for downloaded files")
	version := flag.Bool("v", false, "print version")
	flag.Parse()

	if *version {
		fmt.Printf("%s v%s\n", os.Args[0], Version)
		os.Exit(0)
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	// timestamped directory to store images in
	tsDir := MakeTimeStampDir(*parentDir)

	// max number of workers (counting semaphore)
	workers := make(chan int, *concurrency)

	var fc FeatureCollection
	geoJSON := download(URLGeoJSON)
	if err := json.Unmarshal(geoJSON, &fc); err != nil {
		log.Errorf("Error unmarshalling GeoJSON data: %v\n", err)
	}

	// progress bar
	nImgs := len(fc.Features)
	uiprogress.Start()
	bar := uiprogress.AddBar(nImgs - 1)
	bar.AppendCompleted()
	bar.PrependElapsed()

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

		imgFileServer := path.Base(imgURL)
		if imgFileServer == "" {
			log.Fatalf("Couldn't derive filename: %s\n", imgURL)
		}
		imgNum, err := strconv.Atoi(reDigits.FindString(imgFileServer))
		if err != nil {
			log.Errorf("Error deriving image number: %s\n", err)
		}
		imgFile := fmt.Sprintf("%03d.jpg", imgNum)

		go func(URL, file, dir string) {
			defer wg.Done()
			img := download(URL)
			saveFile(img, file, dir)
			bar.Incr()

			// when download finished, free one slot in workers
			<-workers
		}(imgURL, imgFile, tsDir)
	}

	wg.Wait()
}
