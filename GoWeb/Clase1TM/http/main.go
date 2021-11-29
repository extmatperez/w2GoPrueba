package main

import (
	"fmt"
	"net/http"
)

func funcionSaludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo!")
}

func main() {

	http.HandleFunc("/hola", funcionSaludar)
	http.HandleFunc("/adios", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Adios!")
	})
	http.ListenAndServe(":8080", nil)
}
