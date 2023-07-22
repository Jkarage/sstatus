package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/memstats", memstats)
	r.HandleFunc("/cpustats", cpustats)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
