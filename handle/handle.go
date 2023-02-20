package handle

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Word struct {
	Words string
}

type DataToArtistPage struct {
	Concerts Coords
	Artist   Band
}

var (
	index, indParse      = template.ParseFiles("web/template/index.html")
	errPage, errParse    = template.ParseFiles("web/template/error.html")
	artistPage, artParse = template.ParseFiles("web/template/artist.html")
)

func (groupie *Groupie) MainHandler(w http.ResponseWriter, r *http.Request) {
	if groupie.err1 != nil {
		ErrorPageExecute(w, http.StatusInternalServerError)
		return
	}
	log.Println(r.Method, "/")
	if r.URL.Path != "/" {
		ErrorPageExecute(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPageExecute(w, http.StatusMethodNotAllowed)
		return
	}
	IndexPageExecute(w, *groupie)
}

func (groupie *Groupie) ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if groupie.err1 != nil {
		ErrorPageExecute(w, http.StatusInternalServerError)
		return
	}
	log.Println(r.Method, "/artist")
	link := r.URL.Path
	linkList := strings.Split(link, "/")
	id, err := strconv.Atoi(linkList[len(linkList)-1])
	// fmt.Println(id)
	if err != nil || id > 52 || id < 1 {
		ErrorPageExecute(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPageExecute(w, http.StatusMethodNotAllowed)
		return
	}

	points := GetLocationPoints(groupie.Artists[id-1].LocationsList)
	res := DataToArtistPage{points, groupie.Artists[id-1]}

	ArtistPageExecute(w, res)
}

func ErrorPageExecute(writer http.ResponseWriter, status int) {
	writer.WriteHeader(status)
	if errParse == nil {
		if err := errPage.Execute(writer, status); err == nil {
			return
		}
	}
	http.Error(writer, errParse.Error(), http.StatusInternalServerError)
}

func IndexPageExecute(writer http.ResponseWriter, groupie Groupie) {
	if indParse == nil {
		if err := index.Execute(writer, groupie); err == nil {
			return
		}
	}
	ErrorPageExecute(writer, http.StatusInternalServerError)
}

func ArtistPageExecute(writer http.ResponseWriter, data DataToArtistPage) {
	if artParse == nil {
		if err := artistPage.Execute(writer, data); err == nil {
			return
		}
	}
	ErrorPageExecute(writer, http.StatusInternalServerError)
}
