package main

import (
	"database/sql"
	"net/http"
	dblayer "url-shortner/DBLayer"
)

func registerHandlers(db *sql.DB) map[string]http.Handler {

	getURL := getUrlHandler{
		db:          db,
		getShortUrl: dblayer.GetShortUrl,
	}
	removeURL := removeUrlHandler{
		db: db,
	}
	updateUrl := updateUrlHandler{
		db: db,
	}
	addURL := shortenUrlHandler{
		db:             db,
		InsertURL:      dblayer.InsertURL,
		CheckUrlExists: dblayer.CheckUrlExists,
	}

	m := make(map[string]http.Handler)
	m["get-url"] = &getURL
	m["remove-url"] = &removeURL
	m["update-url"] = &updateUrl
	m["add-url"] = &addURL

	return m
}
