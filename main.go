package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event


func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Println(string(`{"adrien": "12345", "adrien2": "abc", "time": "28/Jul/2006:10:22:04 -0300"}`))
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
