package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// var groupie = handle.Groupie{}

type Input struct {
	Response struct {
		GeoObjectCollection struct {
			MetaDataProperty struct {
				GeocoderResponseMetaData struct {
					BoundedBy struct {
						Envelope struct {
							LowerCorner string `json:"lowerCorner"`
							UpperCorner string `json:"upperCorner"`
						} `json:"Envelope"`
					} `json:"boundedBy"`
					Request string `json:"request"`
					Results string `json:"results"`
					Found   string `json:"found"`
				} `json:"GeocoderResponseMetaData"`
			} `json:"metaDataProperty"`
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							Precision string `json:"precision"`
							Text      string `json:"text"`
							Kind      string `json:"kind"`
							Address   struct {
								CountryCode string `json:"country_code"`
								Formatted   string `json:"formatted"`
								Components  []struct {
									Kind string `json:"kind"`
									Name string `json:"name"`
								} `json:"Components"`
							} `json:"Address"`
							AddressDetails struct {
								Country struct {
									AddressLine        string `json:"AddressLine"`
									CountryNameCode    string `json:"CountryNameCode"`
									CountryName        string `json:"CountryName"`
									AdministrativeArea struct {
										AdministrativeAreaName string `json:"AdministrativeAreaName"`
										Locality               struct {
											LocalityName string `json:"LocalityName"`
										} `json:"Locality"`
									} `json:"AdministrativeArea"`
								} `json:"Country"`
							} `json:"AddressDetails"`
						} `json:"GeocoderMetaData"`
					} `json:"metaDataProperty"`
					Name        string `json:"name"`
					Description string `json:"description"`
					BoundedBy   struct {
						Envelope struct {
							LowerCorner string `json:"lowerCorner"`
							UpperCorner string `json:"upperCorner"`
						} `json:"Envelope"`
					} `json:"boundedBy"`
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

type Point struct {
	Latt float64
	Long float64
}

type Coords struct {
	Points []Point
	Size   int
}

func GetLocationPoints(locations []string) Coords {
	var result Coords

	for _, loc := range locations {
		pos := strings.Split(loc, ", ")

		request := fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?apikey=053d5266-ba49-4f1e-9159-b96a012464fa&geocode=%s+%s&format=json", pos[0], pos[1])

		response, _ := http.Get(request)
		content, err := ioutil.ReadAll(response.Body)

		var data Input
		json.Unmarshal(content, &data)
		if err != nil {
			log.Fatal("could not marshal json: %s\n", err)
		}
		coord := strings.Split(data.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos, " ")
		Latt, _ := strconv.ParseFloat(coord[0], 32)
		Long, _ := strconv.ParseFloat(coord[1], 32)
		point := Point{Latt, Long}

		result.Points = append(result.Points, point)
		result.Size += 1
	}

	return result
}
