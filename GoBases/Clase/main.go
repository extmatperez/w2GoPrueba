package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Users struct {
	Name     string `json:"name"`
	LastName string
	email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	jsonData := `{"name":"Pepe", "LastName":"Lopez", "email":"pepe@gmail.com","Age":15}`
	pepe := Users{}
	if err := json.Unmarshal([]byte(jsonData), &pepe); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", pepe)
}
