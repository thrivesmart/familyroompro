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

func crawled(res http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "user=thrivesmart1 dbname=frp sslmode=disable")
	if err != nil {
		fmt.Fprintln(res, err)
	}

	age := 21
	_, err = db.Query("SELECT title FROM movies WHERE imdbid = $1", age)
	fmt.Fprintln(res, err)
}
