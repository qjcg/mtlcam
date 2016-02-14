package main

type FeatureCollection struct {
	Features []struct {
		Geometry struct {
			Coordinates [2]float64 `json:"coordinates,Number"`
			Type        string
		}

		// NOTE: Always `"description": null` for city dataset
		Properties struct {
			AxeRoutierEstOuest     string `json:"axe-routier-est-ouest"`
			AxeRoutierNordSud      string `json:"axe-routier-nord-sud"`
			Description            string `json:"description,omitempty"`
			IdArrondissement       int    `json:"id-arrondissement,Number"`
			IdCamera               int    `json:"id-camera,Number"`
			Nid                    int    `json:"nid,Number"`
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
