package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	dblayer "url-shortner/DBLayer"
)

const shortUrlSize = 8

type shortenUrlHandler struct {
	db             *sql.DB
	InsertURL      func(db *sql.DB, u *dblayer.URL) error
	CheckUrlExists func(db *sql.DB, u *dblayer.URL) (*dblayer.URL, error)
}

func (h *shortenUrlHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req)
}

func uniqueString(n int) string {
	// seed the randomness with current time
	rand.Seed(time.Now().UnixNano())
	// characters used to create random string
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func (h *shortenUrlHandler) parseRequest(req *http.Request) (*dblayer.URL, error) {
	query := req.URL.Query()
	long_url := query.Get("long_url")
	if len(long_url) == 0 {
		return nil, errors.New("long url cannot be empty")
	}
	short_url := uniqueString(shortUrlSize)
	nativeUrl := &dblayer.URL{
		Original_url: long_url,
		Short_code:   short_url,
		Created_at:   time.Now(),
	}
	return nativeUrl, nil
}
func (h *shortenUrlHandler) handle(w http.ResponseWriter, req *http.Request) {
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
	existUrl, _ := h.CheckUrlExists(h.db, url)
	log.Println("value of existUrl:", existUrl)
	if existUrl != nil {
		httpStatusCode = http.StatusBadRequest
		w.WriteHeader(httpStatusCode)
		response := Response{Success: false, Message: "entry already exists for this url"}
		json, _ := json.Marshal(response)
		_, err := w.Write(json)
		if err != nil {
			os.Exit(5)
		}
		return
	}

	err = h.InsertURL(h.db, url)
	if err != nil {
		log.Printf("%s", err)
		os.Exit(10)
		httpStatusCode = http.StatusFailedDependency
		w.WriteHeader(httpStatusCode)
		response := Response{Success: false}
		json, _ := json.Marshal(response)
		_, err := w.Write(json)
		if err != nil {
			os.Exit(5)
		}
	}

	w.WriteHeader(httpStatusCode)
	response := Response{Success: true, ShortUrl: url.Short_code}
	json, _ := json.Marshal(response)
	_, errr := w.Write(json)
	if errr != nil {
		os.Exit(5)
	}
}
