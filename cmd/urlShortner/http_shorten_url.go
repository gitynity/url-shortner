package main

import (
	"database/sql"
	"net/http"
)

type shortenUrlHandler struct {
	db *sql.DB
}

func (h *shortenUrlHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req)
}

func (h *shortenUrlHandler) handle(w http.ResponseWriter, req *http.Request) {
	body := []byte{'a'}
	w.Write(body)
}
