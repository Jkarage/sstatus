package main

import (
	"fmt"
	"net/http"
)

func memstats(w http.ResponseWriter, r *http.Request) {

}

func cpustats(w http.ResponseWriter, r *http.Request) {

}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found")
}
