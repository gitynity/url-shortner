package main

import( 
  "github.com/gorilla/mux"
  "net/http"
)


func main(){
  r:=mux.NewRouter()
  handlers:=registerHandlers()
  r.Handle("/get-url", handlers["get-url"])
  r.Handle("/remove-url",handlers["remove-url"])
  r.Handle("update-url",handlers["update-url"])
  r.Handle("/add-url",handlers["add-url"])
  
/*
In Go, when you pass a value of a struct to a function or method that expects an interface, the value is automatically converted to the interface type. However, if the interface is defined with a pointer receiver method, you need to pass a pointer to the struct instead of the value itself.

In the case of mux.Router.Handle(), it expects an http.Handler interface value, which has the ServeHTTP method defined with a pointer receiver. Therefore, you need to pass a pointer to the getUrlHandler struct using the & operator, like &handler, to satisfy the interface requirement.
*/
  
  http.ListenAndServe("localhost:8080", r)
/*
Notice that ListenAndServe needs (path,handler) but we have passed (path, mux.Router)
This works because mux.Router also implements ServeHTTP and is therefor a handler as well.
which is basically this--> it matches the handler as per the route and plugs the ServeHTTP of that handler
*/
}
