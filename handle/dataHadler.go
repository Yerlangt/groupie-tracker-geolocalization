package handle

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const ApiURL = "https://groupietrackers.herokuapp.com/api"

type Band struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Dates        string   `json:"concertDates"`
	// Relations    Relation
	// Relationship Relations
	Relationships map[string][]string
	LocationsList []string
	DatesList     []string
}

//	type Relations struct {
//		DatesLocations map[string][]string
//	}
type Location struct {
	Index []struct {
		Id           int      `json: id`
		Locations    []string `json:"locations`
		Locationship []string
	}
}
type Dates struct {
	Index []struct {
		Id       int      `json: id`
		Dates    []string `json:"dates`
		Dateship []string `json:"dates`
	}
}
type Relation struct {
	Index []struct {
		Id             int                 `json:id`
		DatesLocations map[string][]string `json:"datesLocations`
		DateLocation   map[string][]string
	}
}

type URLs struct {
	ArtistsURL   string `json:"artists"`
	LocationsURL string `json:"locations"`
	DatesURL     string `json:"dates"`
	RelationURL  string `json:"relation"`
}

type Groupie struct {
	Artists      []Band
	Size         int
	Relationship Relation
	Locations    Location
	Dates        Dates
	err1         error
}

func (url *URLs) GetURL() error {
	res, err1 := GetPageData(ApiURL)
	if err1 != nil {
		return err1
	}
	if err := json.Unmarshal(res, &url); err != nil {
		return err
	}
	return nil
}

func (groupie *Groupie) GetData(url *URLs) {
	artists, err1 := GetPageData(url.ArtistsURL)
	if err1 != nil {
		groupie.err1 = err1
	}
	if err := json.Unmarshal(artists, &groupie.Artists); err != nil {
		groupie.err1 = err
	}
	groupie.Size = len(groupie.Artists)
	rel, err1 := GetPageData(url.RelationURL)
	if err1 != nil {
		groupie.err1 = err1
	}
	if err := json.Unmarshal(rel, &groupie.Relationship); err != nil {
		groupie.err1 = err
	}
	locn, err1 := GetPageData(url.LocationsURL)
	if err1 != nil {
		groupie.err1 = err1
	}
	if err := json.Unmarshal(locn, &groupie.Locations); err != nil {
		groupie.err1 = err
	}
	date, err1 := GetPageData(url.DatesURL)
	if err1 != nil {
		groupie.err1 = err1
	}
	if err := json.Unmarshal(date, &groupie.Dates); err != nil {
		groupie.err1 = err
	}
	for i := range groupie.Artists {
		groupie.Relationship.Index[i].DateLocation = make(map[string][]string)
		for key := range groupie.Relationship.Index[i].DatesLocations {
			value := groupie.Relationship.Index[i].DatesLocations[key]
			key = fixedLocation(key)
			groupie.Relationship.Index[i].DateLocation[key] = value
		}
		groupie.Artists[i].Relationships = groupie.Relationship.Index[i].DateLocation
		for j := range groupie.Locations.Index[i].Locations {
			groupie.Locations.Index[i].Locationship = append(groupie.Locations.Index[i].Locationship, fixedLocation(groupie.Locations.Index[i].Locations[j]))
		}
		groupie.Artists[i].LocationsList = groupie.Locations.Index[i].Locationship
		for j := range groupie.Dates.Index[i].Dates {
			groupie.Dates.Index[i].Dateship = append(groupie.Dates.Index[i].Dateship, fixedDate(groupie.Dates.Index[i].Dates[j]))
		}
		groupie.Artists[i].DatesList = groupie.Dates.Index[i].Dateship
	}
}

func fixedDate(key string) string {
	// fmt.Println(key)
	keyLists := strings.Split(key[1:], "-")
	key = strings.Join(keyLists, ".")
	return key
}

func fixedLocation(key string) string {
	// fmt.Println(key)
	keyLists := strings.Split(key, "-")
	city := keyLists[0]
	country := keyLists[1]
	city = fixName(city)
	country = fixName(country)
	key = city + ", " + country
	return key
}

func fixName(name string) string {
	if name == "usa" {
		return "USA"
	}
	nameParts := strings.Split(name, "_")
	for i := range nameParts {
		nameParts[i] = strings.Title(nameParts[i])
	}
	name = strings.Join(nameParts, "-")
	return name
}

func GetPageData(link string) ([]byte, error) {
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return content, nil
}

func GroupieCreator() Groupie {
	url := URLs{}
	err := url.GetURL()
	groupie := Groupie{}
	if err != nil {
		groupie.err1 = err
		return groupie
	}
	groupie.GetData(&url)
	return groupie
}
