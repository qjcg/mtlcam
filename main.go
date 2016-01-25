// Download Montreal traffic camera images
// Uses open data from the city of Montreal at:
//   - http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere
// See in action via city map application at @ http://ville.montreal.qc.ca/circulation/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/fatih/color"
)

const (
	// See http://donnees.ville.montreal.qc.ca/dataset/cameras-observation-routiere
	URLGeoJSON string = "http://ville.montreal.qc.ca/circulation/sites/ville.montreal.qc.ca.circulation/files/cameras-de-circulation.json"
	URLBase    string = "http://www1.ville.montreal.qc.ca/Circulation-Cameras/GEN%03d.jpeg"
)

var (
	// functions to colorize strings for use in sprintf-style functions
	bluef  = color.New(color.FgBlue, color.Bold).SprintFunc()
	greenf = color.New(color.FgGreen, color.Bold).SprintFunc()
)

// TODO: Use GeoJSON data for downloading below rather than for-loop
type TrafficCam struct {
	Geometry struct {
		Coordinates [2]float64
		Type        string
	}

	Properties struct {
		AxeRoutierEstOuest     string
		AxeRoutierNordSud      string
		Description            string
		IdArrondissement       int
		IdCamera               int
		Nid                    int
		Titre                  string
		URL                    string
		URLImageDirectionEst   string
		URLImageDirectionNord  string
		URLImageDirectionOuest string
		URLImageDirectionSud   string
		URLImageEnDirect       string
	}

	Type string
}

// Download image at URL to specified directory.
func DownloadImage(URL string, dir string) {
	resp, err := http.Get(URL)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("Couldn't GET image.")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Couldn't read reesponse body.")
	}

	filename := path.Base(URL)
	if filename == "" {
		log.Fatalf("Couldn't derive filename for %s", URL)
	}

	err = ioutil.WriteFile(path.Join(dir, filename), contents, 0644)
	if err != nil {
		log.Fatal("Couldn't create file -- ", err)
	}
}

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
	//concurrency := flag.Int("c", 10, "max download concurrency")
	parentDir := flag.String("d", "mtlcam", "parent directory for downloaded files")
	flag.Parse()

	// timestamped directory to store images in
	tsDir := MakeTimeStampDir(*parentDir)

	// FIXME: use a pool of workers based on concurrency flag value
	var wg sync.WaitGroup
	for i := 1; i <= 500; i++ {
		wg.Add(1)
		url := fmt.Sprintf(URLBase, i)
		go func(URL, dir string) {
			defer wg.Done()
			defer log.Printf("%s %s\n", greenf("DONE"), URL)

			log.Printf("%s %s\n", bluef("GET"), URL)
			DownloadImage(URL, dir)
		}(url, tsDir)
	}
	wg.Wait()
}
