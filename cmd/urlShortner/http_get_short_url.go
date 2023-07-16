package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	dblayer "url-shortner/DBLayer"

	"github.com/redis/go-redis/v9"
)

type Response struct {
	ShortUrl string `json:"short_url,omitempty"`
	LongUrl  string `json:"long_url,omitempty"`
	Success  bool   `json:"success,omitempty"`
	Message  string `json:"message,omitempty"`
}

type getShortUrlHandler struct {
	db          *sql.DB
	rdb         *redis.Client
	getShortUrl func(db *sql.DB, u *dblayer.URL) (*dblayer.URL, error)
}

func (h *getShortUrlHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req)
}

func (h *getShortUrlHandler) parseRequest(req *http.Request) (*dblayer.URL, error) {
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

func (h *getShortUrlHandler) handle(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
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
	cacheURL, err := h.rdb.Get(ctx, url.Original_url).Result()
	if err == nil {
		w.WriteHeader(httpStatusCode)
		response := Response{Success: true, ShortUrl: cacheURL, Message: "URL fetched from cache"}
		json, _ := json.Marshal(response)
		_, err := w.Write(json)
		if err != nil {
			os.Exit(5)
		}
		return
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
		return
	}
	err = h.rdb.Set(ctx, url.Original_url, shorturl.Short_code, 0).Err()
	if err != nil {
		log.Printf("%s", err)
		httpStatusCode = http.StatusExpectationFailed
		w.WriteHeader(httpStatusCode)
		response := Response{Success: false, Message: "failed to write in cache"}
		json, _ := json.Marshal(response)
		_, err := w.Write(json)
		if err != nil {
			os.Exit(5)
		}
		return
	}
	w.WriteHeader(httpStatusCode)
	response := Response{Success: true, ShortUrl: shorturl.Short_code}
	json, _ := json.Marshal(response)
	_, err = w.Write(json)
	if err != nil {
		os.Exit(5)
	}
}
