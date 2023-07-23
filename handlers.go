package main

import (
	"fmt"
	"net/http"
)

func memstats(w http.ResponseWriter, r *http.Request) {
	mem, err := getMemoryStatus()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal server Error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Total Memory: %d KB\nFree Memory:  %d  KB\nAvailable:    %d KB\n", mem.Total, mem.Free, mem.Available)
}

func cpustats(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	t, i, err := getCPUStats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "An Internal error occurred")
	}

	fmt.Fprintf(w, "CPU usage time:    %f\nThe cpu idle time: %f", t, i)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found")
}
