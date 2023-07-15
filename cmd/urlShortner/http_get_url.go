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

type Response struct {
	ShortUrl string `json:"short_url,omitempty"`
	Success  bool   `json:"success,omitempty"`
	Message  string `json:"message,omitempty"`
}

type getUrlHandler struct {
	db          *sql.DB
	getShortUrl func(db *sql.DB, u *dblayer.URL) (*dblayer.URL, error)
}

func (h *getUrlHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req)
}

func (h *getUrlHandler) parseRequest(req *http.Request) (*dblayer.URL, error) {
	query := req.URL.Query()
	long_url := query.Get("long_url")
	if len(long_url) == 0 {
		return nil, errors.New("long url cannot be empty")
	}
	nativeUrl := &dblayer.URL{
		Original_url: long_url,
		Short_code:   "",
	}
	return nativeUrl, nil
}

func (h *getUrlHandler) handle(w http.ResponseWriter, req *http.Request) {
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
	shorturl, err := h.getShortUrl(h.db, url)
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
	response := Response{Success: true, ShortUrl: shorturl.Short_code}
	json, _ := json.Marshal(response)
	_, err = w.Write(json)
	if err != nil {
		os.Exit(5)
	}
}
