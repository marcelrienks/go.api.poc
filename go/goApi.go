package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func customResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom response!")
	fmt.Println("Endpoint Hit: customResponse")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/custom", customResponse)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
