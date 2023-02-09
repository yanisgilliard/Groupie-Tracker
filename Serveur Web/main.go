// Serveur
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const port = ":8080"  

type rawArtist struct {
	Id           int      `json:"id"` 
	Image        string   `json:"image"` 
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func getArtists() []rawArtist {
	artists := []rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists
}
func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", GroupieTracker)
	http.HandleFunc("/artist", Artist)
	http.HandleFunc("/search", Search)
	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)

}

func getData(url string, data interface{}) {
	rawData := getRawData(url)
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		log.Panic("Problème dans la fonction getData lors du déclassement des données :", err)
	}
}

func getRawData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Panic("Problème dans la fonction getRawData lors de l'obtention de la réponse :", err)
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic("Problème dans la fonction getRawData lors de la lecture de la réponse :", err)
		return nil
	}
	return responseData
}

func GroupieTracker(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	data := getArtists()
	fmt.Println(data)
	tmpl.Execute(w, data)
}

func Artist(w http.ResponseWriter, r *http.Request) {
	// récupérer l'id de l'artiste passé dans la requête GET
	artistId := r.URL.Query().Get("id")
	fmt.Println(artistId)
	// récupérer les données de l'artiste
	artist := getArtist(artistId)
	fmt.Println(artist)

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, artist)
}

func getArtist(artistId string) rawArtist {
	artist := rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists/"+artistId, &artist)
	return artist
}
func Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Println("Recherche pour:", query)
	// Code pour faire la recherche dans les données de l'API
}
   