package main

import (
	"log"
	"net/http"
	"math/rand"
	"time"
	"strconv"
	"flag"
)

var (
	users Users
	tokens TokenCollection
)

func setup() {
	rand.Seed(time.Now().UnixNano())
	users = make(Users, 0)
	tokens = TokenCollection{}
	router = NewRouter()
}
var (
	port = flag.Int("port", 8080, "Port")
)

func main() {
	setup()
	log.Fatal(http.ListenAndServe("0.0.0.0:" + strconv.Itoa(*port), router))
}
