package main

import (
  "fmt"
  "net/http"
  "os"
  "github.com/skratchdot/open-golang/open"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" { 
    port = "3000"
  }
  http.HandleFunc("/", hello)
  http.HandleFunc("/yahoo", yhoo)
  fmt.Println("listening on port " + port + "...")
  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    panic(err)
  }
}

func hello(res http.ResponseWriter, req *http.Request) {
  fmt.Fprintln(res, "Hello, world!")
}

func yhoo(res http.ResponseWriter, req *http.Request) {
  open.RunWith("http://www.yahoo.com/", "safari") 
}