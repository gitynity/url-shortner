package main

import "net/http"

type updateUrlHandler struct{}
func (h *updateUrlHandler) ServeHTTP(w http.ResponseWriter,req *http.Request){
h.handle(w,req)
}

func(h *updateUrlHandler) handle(w http.ResponseWriter, req *http.Request){
  body := []byte{'a'}
  w.Write(body)
}
