package main

import (
	"log"
	"net/http"

	"github.com/OzBank/oz-echo-api/api"
)

func main() {
	server := api.NewServer()
	r := http.NewServeMux()
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
