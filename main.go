package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
        "github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title divya Reddy", Desc: "Test Descritpion  studying go language", Content: "Hello world"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func testPostArticles(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "Test POST endpoint worked") 
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Divya Reddy")
}
func handleRequests() {
        myRouter :=mux.NewRouter().StrictSlash(true)
        myRouter.HandleFunc("/", homepage)
        myRouter.HandleFunc("/articles", allArticles).Methods("GET")
        myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	handleRequests()
}