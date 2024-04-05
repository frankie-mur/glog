package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /post/{slug}", PostHandler(FileReader{}))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
