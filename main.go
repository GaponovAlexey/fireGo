package main

import (
	"log"
	"net/http"
	"server/firebase/routes"

	"github.com/gorilla/mux"

)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.GetPosts).Methods("GET")

	log.Println("start port 3000")
	http.ListenAndServe(":3000", r)
}
