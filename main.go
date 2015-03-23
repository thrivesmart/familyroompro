package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", hello)
	http.HandleFunc("/yahoo", yhoo)
	http.HandleFunc("/crawl-result", crawled)
	fmt.Println("listening on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

type Episode struct {
	Id                    int64
	SeriesId              int64
	Title                 string
	Year                  int64
	Rated                 string
	Released              time.Time
	RuntimeMinutes        int64
	Plot                  string
	Language              string
	PosterUrl             string
	TomameterRating       int64
	TomameterVoteCount    int64
	Imdbid                string
	ImdbRating            int64
	ImdbVoteCount         int64
	GenresJson            string
	DirectorsJson         string
	WritersJson           string
	ActorsJson            string
	CountriesJson         string
	SubtitleLanguagesJson string
	StreamingUrlsJson     string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type Movie struct {
	Id                    int64
	Title                 string
	Year                  int64
	Rated                 string
	Released              time.Time
	RuntimeMinutes        int64
	Plot                  string
	Language              string
	PosterUrl             string
	TomameterRating       int64
	TomameterVoteCount    int64
	Imdbid                string
	ImdbRating            int64
	ImdbVoteCount         int64
	GenresJson            string
	DirectorsJson         string
	WritersJson           string
	ActorsJson            string
	CountriesJson         string
	SubtitleLanguagesJson string
	StreamingUrlsJson     string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type Series struct {
	Id                    int64
	Title                 string
	StartYear             int64
	EndYear               int64
	Rated                 string
	Released              time.Time
	Plot                  string
	Language              string
	PosterUrl             string
	TomameterRating       int64
	TomameterVoteCount    int64
	Imdbid                string
	ImdbRating            int64
	ImdbVoteCount         int64
	GenresJson            string
	DirectorsJson         string
	WritersJson           string
	ActorsJson            string
	CountriesJson         string
	SubtitleLanguagesJson string
	StreamingUrlsJson     string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello, world!")
}

func yhoo(res http.ResponseWriter, req *http.Request) {
	open.RunWith("http://www.yahoo.com/", "safari")
}

func crawled(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "user=thrivesmart1 dbname=frp sslmode=disable")
	if err != nil {
		fmt.Fprintln(res, err)
	}

	age := 21
	_, err = db.Query("SELECT title FROM movies WHERE imdbid = $1", age)
	fmt.Fprintln(res, err)
}
