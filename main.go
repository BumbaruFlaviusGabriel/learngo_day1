package main

import (
	"encoding/json"
	"exercise4/entity"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var person entity.Person

func main() {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		router.Get("/person", GetPerson)
		router.Post("/person", CreatePerson)
	})

	router.Route("/v2", func(r chi.Router) {
		router.Get("/person", GetPerson)
		router.Post("/person", CreatePerson)
	})

	http.ListenAndServe(":8080", router)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &person)

	person = entity.Person{
		Name:   "Flavius",
		ID:     "1",
		Gender: "Male",
	}

	fmt.Fprint(w, person)
}

func GetPerson2(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func CreatePerson2(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)

	person = entity.Person{
		Name:   "Gabriel",
		ID:     "2",
		Gender: "Male",
	}

	fmt.Fprint(w, person)
}
