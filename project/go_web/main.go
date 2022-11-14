package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/Camden-Levinson/bazeltut/project/go_hello_world"
	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	var list = []string{"add", "subtract", "multiple", "divide"}
	randIndex := rand.Intn(len(list))
	randone := rand.Float64() + float64(rand.Intn(100))
	randtwo := rand.Float64() + float64(rand.Intn(100))
	switch list[randIndex] {
	case "add":
		w.Write([]byte(fmt.Sprintf("%f", randone) + " + " + fmt.Sprintf("%f", randtwo) + " = " + fmt.Sprintf("%f", go_hello_world.Add(randone, randtwo))))
	case "subtract":
		w.Write([]byte(fmt.Sprintf("%f", randone) + " - " + fmt.Sprintf("%f", randtwo) + " = " + fmt.Sprintf("%f", go_hello_world.Subtract(randone, randtwo))))
	case "multiple":
		w.Write([]byte(fmt.Sprintf("%f", randone) + " * " + fmt.Sprintf("%f", randtwo) + " = " + fmt.Sprintf("%f", go_hello_world.Multiple(randone, randtwo))))
	case "divide":
		w.Write([]byte(fmt.Sprintf("%f", randone) + " / " + fmt.Sprintf("%f", randtwo) + " = " + fmt.Sprintf("%f", go_hello_world.Divide(randone, randtwo))))
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)

	log.Println("Going to listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
