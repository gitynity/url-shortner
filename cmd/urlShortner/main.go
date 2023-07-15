package main

import (
	"fmt"
	"net/http"
	"os"
	dblayer "url-shortner/DBLayer"

	"github.com/gorilla/mux"
)

func main() {
	db, err := dblayer.DBconfig()
	if err != nil {
		fmt.Println("Getting DB Connection", err)
		os.Exit(3)
	}
	r := mux.NewRouter()
	handlers := registerHandlers(db)

	r.Handle("/get-url", handlers["get-url"]).Methods(http.MethodGet)

	r.Handle("/remove-url", handlers["remove-url"]).Methods(http.MethodDelete)

	r.Handle("/add-url", handlers["add-url"]).Methods(http.MethodPost)

	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		os.Exit(4)
	}

	fmt.Println("Server Started")
}
