package main

import "net/http"

type removeUrlHandler struct{}
func (h *removeUrlHandler) ServeHTTP(w http.ResponseWriter,req *http.Request){
h.handle(w,req)
}

func(h *removeUrlHandler) handle(w http.ResponseWriter, req *http.Request){
  body := []byte{'a'}
  w.Write(body)
}
