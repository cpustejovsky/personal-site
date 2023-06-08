package main

import (
	"github.com/cpustejovsky/personal-site/handlers"
	"log"
	"net/http"
)

// TODO: configure address for deployment
const addr = ":8080"

func main() {
	log.Printf("listening on %s", addr)
	handler, err := handlers.New()
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal(err)
	}
}
