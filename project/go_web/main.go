package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request")
	w.Write([]byte(go_hello_world.HelloWorld()))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)

	log.Println("GOing to listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
