package main

import (
	"net/http"
)

func registerHandlers() map[string]http.Handler{
  
  getURL:=getUrlHandler{}
  removeURL:=removeUrlHandler{}
  updateUrl:=updateUrlHandler{}
  addURL:=shortenUrlHandler{}
  
  m:=make(map[string]http.Handler)
  m["get-url"] = &getURL
  m["remove-url"] = &removeURL
  m["update-url"] = &updateUrl
  m["add-url"]=&addURL

  return m
}
