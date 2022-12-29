package main

import (
	"bytes"
	"log"
	"net/http"
	"testProject/internal/handlers/randomizer"
)

func main() {

	for i := 0; i < 20; i++ {
		_, err := http.Post("http://localhost:8000/create-announcement", "application/json", bytes.NewBuffer(randomizer.SenderRandomString()))
		if err != nil {
			log.Fatal(err)
		}
	}
}
