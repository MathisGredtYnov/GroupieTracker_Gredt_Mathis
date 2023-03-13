package main

import (
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
	Followers  int    `json:"followers"`
	Country    string `json:"country"`
	Location   string `json:"location"`
	LastOnline int    `json:"last_online"`
	Joined     int    `json:"joined"`
	Status     string `json:"status"`
	IsStreamer bool   `json:"is_streamer"`
	Verified   bool   `json:"verified"`
	League     string `json:"league"`
}

type StatForGames struct {
	ChessDaily struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win            int `json:"win"`
			Loss           int `json:"loss"`
			Draw           int `json:"draw"`
			TimePerMove    int `json:"time_per_move"`
			TimeoutPercent int `json:"timeout_percent"`
		} `json:"record"`
		Tournament struct {
			Points        int `json:"points"`
			Withdraw      int `json:"withdraw"`
			Count         int `json:"count"`
			HighestFinish int `json:"highest_finish"`
		} `json:"tournament"`
	} `json:"chess_daily"`
	Chess960Daily struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win            int `json:"win"`
			Loss           int `json:"loss"`
			Draw           int `json:"draw"`
			TimePerMove    int `json:"time_per_move"`
			TimeoutPercent int `json:"timeout_percent"`
		} `json:"record"`
		Tournament struct {
			Points        int `json:"points"`
			Withdraw      int `json:"withdraw"`
			Count         int `json:"count"`
			HighestFinish int `json:"highest_finish"`
		} `json:"tournament"`
	} `json:"chess960_daily"`
	ChessRapid struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win  int `json:"win"`
			Loss int `json:"loss"`
			Draw int `json:"draw"`
		} `json:"record"`
	} `json:"chess_rapid"`
	ChessBullet struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win  int `json:"win"`
			Loss int `json:"loss"`
			Draw int `json:"draw"`
		} `json:"record"`
	} `json:"chess_bullet"`
	ChessBlitz struct {
		Last struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
			Rd     int `json:"rd"`
		} `json:"last"`
		Best struct {
			Rating int    `json:"rating"`
			Date   int    `json:"date"`
			Game   string `json:"game"`
		} `json:"best"`
		Record struct {
			Win  int `json:"win"`
			Loss int `json:"loss"`
			Draw int `json:"draw"`
		} `json:"record"`
	} `json:"chess_blitz"`
	Fide    int `json:"fide"`
	Tactics struct {
		Highest struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
		} `json:"highest"`
		Lowest struct {
			Rating int `json:"rating"`
			Date   int `json:"date"`
		} `json:"lowest"`
	} `json:"tactics"`
	PuzzleRush struct {
		Best struct {
			TotalAttempts int `json:"total_attempts"`
			Score         int `json:"score"`
		} `json:"best"`
	} `json:"puzzle_rush"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../WEB/index.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
	}
}

func requeteHandler(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("../WEB/requete.html"))
	if r.Method != http.MethodPost {
		tmpl1.Execute(w, nil)
	}
}

func main() {
	//démarage serveur
	fmt.Println("Server is running on port 80 http://localhost")
	//gestion css
	fs := http.FileServer(http.Dir("../WEB/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//gestion assets
	image := http.FileServer(http.Dir("../WEB/ASSETS"))
	http.Handle("/ASSETS/", http.StripPrefix("/ASSETS/", image))
	//gestion html
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/requete", requeteHandler)
	http.ListenAndServe(":80", nil)
}
