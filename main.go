package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getRedirect(abbreviation string) string {

	switch abbreviation {
	case "go":
		return "https://gobyexample.com/"
	case "tr":
		return "https://play.typeracer.com/"
	case "yt":
		return "http://youtube.com"
	}

	return "" // TODO : Query DB against "abbreviation" and return mapped domain
}

func getRedirectHandler(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	mapping := request["id"]

	mapped := getRedirect(mapping)

	fmt.Println("Redirecting", mapping, "to", mapped)
	http.Redirect(w, r, mapped, http.StatusFound)
}

func main() {

	fmt.Println("Starting listening on port 8080")

	r := mux.NewRouter()

	r.HandleFunc("/{id}", getRedirectHandler).Methods("GET")
	http.ListenAndServe(":8080", r)

}
