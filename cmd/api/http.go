package main

import (
	"database/sql"
	"net/http"
	dblayer "url-shortner/DBLayer"

	"github.com/redis/go-redis/v9"
)

func registerHandlers(db *sql.DB, cache *redis.Client) map[string]http.Handler {

	getShortURL := getShortUrlHandler{
		db:          db,
		rdb:         cache,
		getShortUrl: dblayer.GetShortUrl,
	}
	getLongURL := getLongUrlHandler{
		db:         db,
		rdb:        cache,
		getLongUrl: dblayer.GetLongUrl,
	}
	removeURL := removeUrlHandler{
		db:        db,
		rdb:       cache,
		DeleteURL: dblayer.DeleteURL,
	}
	addURL := shortenUrlHandler{
		db:             db,
		rdb:            cache,
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
