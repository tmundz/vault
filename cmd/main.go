package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//get rid of later
type Article struct {
  Title string `json:"Title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r * http.Request) {
  articles := Articles {
    Article{
      Title:"Test Title",
      Desc: "test Desc",
      Content: "Look at this cool test",
    },
  }
  fmt.Println("allArticles")
  json.NewEncoder(w).Encode(articles);
}

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Homepage endpoint")
}


func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/articles", allArticles)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
  handleRequests();
}
