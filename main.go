package main

import (
	"database/sql"
  "encoding/json"
	"fmt"
  "strings"
	_ "github.com/lib/pq"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"os"
	"time"
)

var DB_USERNAME string
// var DB_PASSWORD string
var DB_DBNAME string
var DB_SSLMODE string

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "3000" }
	DB_USERNAME = os.Getenv("FRP_DB_USERNAME")
	if DB_USERNAME == "" { DB_USERNAME = "thrivesmart1" }
  // DB_PASSWORD = os.Getenv("FRP_DB_PASSWORD")
	DB_DBNAME = os.Getenv("FRP_DB_DBNAME")
	if DB_DBNAME == "" { DB_DBNAME = "frp" }
	DB_SSLMODE = os.Getenv("FRP_DB_SSLMODE")
	if DB_SSLMODE == "" { DB_SSLMODE = "disable" }
  
	http.HandleFunc("/", hello)
	http.HandleFunc("/yahoo", yhoo)
	http.HandleFunc("/update-library", updates)
	fmt.Println("listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

type Episode struct {
	Id                    int64     `json:"id"`
	SeriesId              int64     `json:"series_id"`
	Title                 string    `json:"title"`
	Year                  int64     `json:"year"`
	Rated                 string    `json:"rated"`
	Released              time.Time `json:"released"`
	RuntimeMinutes        int64     `json:"runtime_minutes"`
	Plot                  string    `json:"plot"`
	Language              string    `json:"language"`
	PosterUrl             string    `json:"poster_url"`
	TomameterRating       int64     `json:"tomameter_rating"`
	TomameterVoteCount    int64     `json:"tomameter_vote_count"`
	Imdbid                string    `json:"imdbid"`
	ImdbRating            int64     `json:"imdb_rating"`
	ImdbVoteCount         int64     `json:"imdb_vote_count"`
	GenresJson            string    `json:"genres_json"`
	DirectorsJson         string    `json:"directors_json"`
	WritersJson           string    `json:"writers_json"`
	ActorsJson            string    `json:"actors_json"`
	CountriesJson         string    `json:"countries_json"`
	SubtitleLanguagesJson string    `json:"subtitle_languages_json"`
	StreamingUrlsJson     string    `json:"streaming_urls_json"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type Movie struct {
	Id                    int64     `json:"id"`
	Title                 string    `json:"title"`
	Year                  int64     `json:"year"`
	Rated                 string    `json:"rated"`
	Released              time.Time `json:"released"`
	RuntimeMinutes        int64     `json:"runtime_minutes"`
	Plot                  string    `json:"plot"`
	Language              string    `json:"language"`
	PosterUrl             string    `json:"poster_url"`
	TomameterRating       int64     `json:"tomameter_rating"`
	TomameterVoteCount    int64     `json:"tomameter_vote_count"`
	Imdbid                string    `json:"imdbid"`
	ImdbRating            int64     `json:"imdb_rating"`
	ImdbVoteCount         int64     `json:"imdb_vote_count"`
	GenresJson            string    `json:"genres_json"`
	DirectorsJson         string    `json:"directors_json"`
	WritersJson           string    `json:"writers_json"`
	ActorsJson            string    `json:"actors_json"`
	CountriesJson         string    `json:"countries_json"`
	SubtitleLanguagesJson string    `json:"subtitle_languages_json"`
	StreamingUrlsJson     string    `json:"streaming_urls_json"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type Series struct {
	Id                    int64     `json:"id"`
	Title                 string    `json:"title"`
	StartYear             int64     `json:"start_year"`
	EndYear               int64     `json:"end_year"`
	Rated                 string    `json:"rated"`
	Released              time.Time `json:"released"`
	Plot                  string    `json:"plot"`
	Language              string    `json:"language"`
	PosterUrl             string    `json:"poster_url"`
	TomameterRating       int64     `json:"tomameter_rating"`
	TomameterVoteCount    int64     `json:"tomameter_vote_count"`
	Imdbid                string    `json:"imdbid"`
	ImdbRating            int64     `json:"imdb_rating"`
	ImdbVoteCount         int64     `json:"imdb_vote_count"`
	GenresJson            string    `json:"genres_json"`
	DirectorsJson         string    `json:"directors_json"`
	WritersJson           string    `json:"writers_json"`
	ActorsJson            string    `json:"actors_json"`
	CountriesJson         string    `json:"countries_json"`
	SubtitleLanguagesJson string    `json:"subtitle_languages_json"`
	StreamingUrlsJson     string    `json:"streaming_urls_json"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello, world!")
}

func yhoo(res http.ResponseWriter, req *http.Request) {
	open.RunWith("http://www.yahoo.com/", "safari")
}

func updates(res http.ResponseWriter, req *http.Request) {
  var err error
  
  connector := []string{"user=", DB_USERNAME, " dbname=", DB_DBNAME, " sslmode=", DB_SSLMODE};
  
	db, err := sql.Open("postgres", strings.Join(connector, ""))
	if err != nil { http.Error(res, err.Error(), 500); return }
  
  kind := req.FormValue("kind")
  js := req.FormValue("json")

  var episode Episode
  var movie Movie
  var series Series

  if kind == "episode" {
    err = json.Unmarshal([]byte(js), &episode)
  	if err != nil { http.Error(res, err.Error(), 500); return }
  } else if kind == "movie" {
    err = json.Unmarshal([]byte(js), &movie)
  	if err != nil { http.Error(res, err.Error(), 500); return }
  } else if kind == "series" {
    err = json.Unmarshal([]byte(js), &series)
  	if err != nil { http.Error(res, err.Error(), 500); return }
  } else {
    http.Error(res, "Missing `kind` parameter.", 422)
    return
  }
  
	age := 21
	_, err = db.Query("SELECT title FROM movies WHERE imdbid = $1", age)
	fmt.Fprintln(res, err)
}
