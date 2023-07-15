package main

import (
	"database/sql"
	"net/http"
	dblayer "url-shortner/DBLayer"
)

func registerHandlers(db *sql.DB) map[string]http.Handler {

	getShortURL := getShortUrlHandler{
		db:          db,
		getShortUrl: dblayer.GetShortUrl,
	}
	getLongURL := getLongUrlHandler{
		db:         db,
		getLongUrl: dblayer.GetLongUrl,
	}
	removeURL := removeUrlHandler{
		db:        db,
		DeleteURL: dblayer.DeleteURL,
	}
	addURL := shortenUrlHandler{
		db:             db,
		InsertURL:      dblayer.InsertURL,
		CheckUrlExists: dblayer.CheckUrlExists,
	}

	m := make(map[string]http.Handler)
	m["get-short-url"] = &getShortURL
	m["remove-url"] = &removeURL
	m["add-url"] = &addURL
	m["get-long-url"] = &getLongURL

	return m
}
