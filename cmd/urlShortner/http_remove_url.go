package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	dblayer "url-shortner/DBLayer"

	"github.com/redis/go-redis/v9"
)

type removeUrlHandler struct {
	db        *sql.DB
	rdb       *redis.Client
	DeleteURL func(db *sql.DB, u *dblayer.URL) error
}

func (h *removeUrlHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.handle(w, req)
}

func (h *removeUrlHandler) parseRequest(req *http.Request) (*dblayer.URL, error) {
	query := req.URL.Query()
	long_url := query.Get("long_url")
	if len(long_url) == 0 {
		return nil, errors.New("long url cannot be empty")
	}
	nativeUrl := &dblayer.URL{
		Original_url: long_url,
	}
	return nativeUrl, nil
}
func (h *removeUrlHandler) handle(w http.ResponseWriter, req *http.Request) {
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
	h.rdb.Del(ctx, url.Original_url, url.Short_code).Err()
	err = h.DeleteURL(h.db, url)
	if err != nil {
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
	response := Response{Success: true, Message: "Deleted"}
	json, _ := json.Marshal(response)
	_, err = w.Write(json)
	if err != nil {
		os.Exit(5)
	}

}
