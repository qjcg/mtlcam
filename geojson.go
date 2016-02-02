package main

// TODO: Use GeoJSON data to obtain image URLs
type GeoJSONData struct {
	Features []struct {
		Geometry struct {
			Coordinates [2]float64
			Type        string
		}

		Properties struct {
			AxeRoutierEstOuest     string `json:"axe-routier-est-ouest"`
			AxeRoutierNordSud      string `json:"axe-routier-nord-sud"`
			Description            string
			IdArrondissement       int `json:"id-arrondissement"`
			IdCamera               int `json:"id-camera"`
			Nid                    int
			Titre                  string
			URL                    string
			URLImageDirectionEst   string `json:"url-image-direction-est"`
			URLImageDirectionNord  string `json:"url-image-direction-nord"`
			URLImageDirectionOuest string `json:"url-image-direction-ouest"`
			URLImageDirectionSud   string `json:"url-image-direction-sud"`
			URLImageEnDirect       string `json:"url-image-en-direct"`
		}

		Type string
	}
	Type string
}

// TODO
func DownloadGeoJSON(url string) {
}
