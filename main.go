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
	r.HandleFunc("/posts", routes.AddPost).Methods("POST")

	log.Println("start port 3000")
	http.ListenAndServe(":3000", r)
}
