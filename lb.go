package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello world")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)

	http.ListenAndServe(":4000", mux)
}