package main

import (
	"log"
	"net/http"

	"github.com/cormoran/test2doc/example/mux/foos"
	"github.com/cormoran/test2doc/example/mux/widgets"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	foos.AddRoutes(router)
	widgets.AddRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
