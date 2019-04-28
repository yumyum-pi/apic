package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = "8080"
const ip = ""

func welcomeToAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the api")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", welcomeToAPI).Methods("GET")
	//--->route
	serverURL := fmt.Sprintf("%v:%v", ip, port)
	log.Fatal(http.ListenAndServe(serverURL, router))
}
