package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sharukh010/go-bookstore/pkg/routes"
)

func main() {
	
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	fmt.Println("Server is starting at 9010 port")
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}