package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://server4:8080")
	if err != nil {
		fmt.Fprint(w, "Something went wrong")
	}
	body, err := io.ReadAll(resp.Body)
	sb := string(body)
	fmt.Fprint(w, sb)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)

	http.ListenAndServe(":4000", mux)
}