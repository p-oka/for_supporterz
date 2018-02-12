package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyHandler struct{}

func (t *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	person := Person{
		Name: "p-oka",
		Age:  25,
	}

	response, err := json.Marshal(person)

	if err != nil {
		log.Fatal("error occurred:", err)
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := "500 - Something bad happened!"
		w.Write([]byte(errorResponse))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func main() {
	http.Handle("/", &MyHandler{})

	log.Println("Golang application starting on http://localhost:8880")
	log.Println("Ctrl-C to shutdown server")

	if err := http.ListenAndServe(":8880", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
