package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	dblayer "url-shortner/DBLayer"
)

type getLongUrlHandler struct {
	db         *sql.DB
	getLongUrl func(db *sql.DB, u *dblayer.URL) (*dblayer.URL, error)
}

func (h *getLongUrlHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req)
}

func (h *getLongUrlHandler) parseRequest(req *http.Request) (*dblayer.URL, error) {
	query := req.URL.Query()
	short_code := query.Get("short_code")
	if len(short_code) == 0 {
		return nil, errors.New("long url cannot be empty")
	}
	nativeUrl := &dblayer.URL{
		Short_code: short_code,
	}
	return nativeUrl, nil
}

func (h *getLongUrlHandler) handle(w http.ResponseWriter, req *http.Request) {
	httpStatusCode := http.StatusOK
	url, err := h.parseRequest(req)
	if err != nil {
		httpStatusCode = http.StatusBadRequest
		w.WriteHeader(httpStatusCode)
		response := Response{Success: false}
		json, _ := json.Marshal(response)
		_, err := w.Write(json)
		if err != nil {
			os.Exit(5)
		}
	}
	longurl, err := h.getLongUrl(h.db, url)
	if err != nil {
		log.Printf("%s", err)
		httpStatusCode = http.StatusExpectationFailed
		w.WriteHeader(httpStatusCode)
		response := Response{Success: false}
		json, _ := json.Marshal(response)
		_, err := w.Write(json)
		if err != nil {
			os.Exit(5)
		}
	}
	w.WriteHeader(httpStatusCode)
	response := Response{Success: true, LongUrl: longurl.Original_url}
	json, _ := json.Marshal(response)
	_, err = w.Write(json)
	if err != nil {
		os.Exit(5)
	}
}
