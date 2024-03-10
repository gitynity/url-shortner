package main

import (
	"log"
	"net/http"
	"os"
	cachelayer "url-shortner/CacheLayer"
	dblayer "url-shortner/DBLayer"

	"github.com/gorilla/mux"
)

func main() {
	db, err := dblayer.DBconfig()
	if err != nil {
		log.Println("Getting DB Connection", err)
		os.Exit(3)
	}
	rdb, err := cachelayer.CacheConfig()
	if err != nil {
		log.Println("Getting redis Connection", err)
		os.Exit(3)
	}
	r := mux.NewRouter()
	handlers := registerHandlers(db, rdb)

	r.Handle("/get-short-url", handlers["get-short-url"]).Methods(http.MethodGet)

	r.Handle("/get-long-url", handlers["get-long-url"]).Methods(http.MethodGet)

	r.Handle("/remove-url", handlers["remove-url"]).Methods(http.MethodDelete)

	r.Handle("/add-url", handlers["add-url"]).Methods(http.MethodPost)

	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		os.Exit(4)
	}

	log.Println("Server Started")
}
