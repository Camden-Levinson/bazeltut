package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	var list = []string{"add", "subtract", "multiple", "divide"}
	randIndex := rand.Intn(len(list))
	randone := rand.Intn(100)
	randtwo := rand.Intn(100)
	switch list[randIndex] {
	case "add":
		ini := string(randone) + " + " + string(randtwo) + " = " + string(go_hello_world.add(randone, randtwo))
		w.Write([]byte(ini))
	case "subtract":
		ini := string(randone) + " - " + string(randtwo) + " = " + string(go_hello_world.subtract(randone, randtwo))
		w.Write([]byte(ini))
	case "multiple":
		ini := string(randone) + " * " + string(randtwo) + " = " + string(go_hello_world.multiple(randone, randtwo))
		w.Write([]byte(ini))
	case "divide":
		ini := string(randone) + " / " + string(randtwo) + " = " + string(go_hello_world.divide(randone, randtwo))
		w.Write([]byte(ini))
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)

	log.Println("GOing to listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
