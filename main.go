package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/antixcode6/cyderes/pkg/api"
	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	app.Path("/api/v1/query").HandlerFunc(api.Ingest).Methods("GET")
	log.Println("---------------------------------")
	log.Println("Running on port 3000")
	log.Println("---------------------------------")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 3000), app))
}
