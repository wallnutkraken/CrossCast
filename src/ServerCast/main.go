package main

import (
	"log"
	"net/http"
	"math/rand"
	"time"
)

var (
	users Users
	tokens TokenCollection
)

func main() {
	rand.Seed(time.Now().UnixNano())
	users = make(Users, 0)
	tokens = TokenCollection{}

	router = NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
