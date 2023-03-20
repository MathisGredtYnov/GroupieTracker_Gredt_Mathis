package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type PlayerInformation struct {
	Avatar     string `json:"avatar"`
	PlayerID   int    `json:"player_id"`
	ID         string `json:"@id"`
	URL        string `json:"url"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Title      string `json:"title"`
	Followers  int    `json:"followers"`
	Country    string `json:"country"`
	Location   string `json:"location"`
	LastOnline int    `json:"last_online"`
	Joined     int    `json:"joined"`
	Status     string `json:"status"`
	IsStreamer bool   `json:"is_streamer"`
	Verified   bool   `json:"verified"`
}

type PlayerData struct {
	Success   bool
	Entree    string
	Data      *PlayerInformation
	DataTitle *PlayerTitle
}

type PlayerTitle struct {
	Players []string `json:"players"`
}

func main() {
	var p PlayerData
	p.init()
	p.InitTitle()
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("WEB/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion assets
	image := http.FileServer(http.Dir("WEB/ASSETS"))
	http.Handle("/ASSETS/", http.StripPrefix("/ASSETS/", image))
	//gestion html
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/requete", p.requeteHandler)
	http.ListenAndServe(":80", nil)
}

func (p *PlayerData) InitTitle() {
	url := "https://api.chess.com/pub/titled/GM"

	// Envoyer une requête GET à l'API
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	// Parser la réponse JSON
	var playerInfo PlayerTitle
	err = json.NewDecoder(res.Body).Decode(&playerInfo)
	if err != nil {
		panic(err.Error())
	}
}

func (p *PlayerData) RechercheTitle(title string) *PlayerTitle {
	url := "https://api.chess.com/pub/titled/" + title
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	// Parser la réponse JSON
	var playerInfo PlayerTitle
	err = json.NewDecoder(res.Body).Decode(&playerInfo)
	if err != nil {
		panic(err.Error())
	}

	return &playerInfo
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("WEB/index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
	}
}

func (p *PlayerData) requeteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("WEB/requete.html"))
	details := PlayerData{
		Success:   true,
		Entree:    r.PostFormValue("title"),
		Data:      p.recherche(r.FormValue("search")),
		DataTitle: p.RechercheTitle(r.PostFormValue("title")),
	}
	tmpl.Execute(w, details)
}

// Fonction qui récupère le profile d'un joueur
func (p *PlayerData) init() {

	url := "https://api.chess.com/pub/player/erik"

	// Envoyer une requête GET à l'API
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	// Parser la réponse JSON
	var playerInfo PlayerInformation
	err = json.NewDecoder(res.Body).Decode(&playerInfo)
	if err != nil {
		panic(err.Error())
	}
}

func (p *PlayerData) recherche(pseudo string) *PlayerInformation {
	url := "https://api.chess.com/pub/player/" + pseudo

	// Envoyer une requête GET à l'API
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	// Parser la réponse JSON
	var playerInfo *PlayerInformation
	err = json.NewDecoder(res.Body).Decode(&playerInfo)
	if err != nil {
		panic(err.Error())
	}
	if playerInfo.Avatar == "" {
		playerInfo.Avatar = "ASSETS/NoAvatar.gif"
	}

	return playerInfo

}
