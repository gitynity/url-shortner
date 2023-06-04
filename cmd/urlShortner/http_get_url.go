package main

import "net/http"

type getUrlHandler struct{}
func (h *getUrlHandler) ServeHTTP(w http.ResponseWriter,req *http.Request){
h.handle(w,req)
}

func(h *getUrlHandler) handle(w http.ResponseWriter, req *http.Request){
  body := []byte{'a'}
  w.Write(body)
}
