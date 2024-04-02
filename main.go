package main

import (
	"fmt"
	"net/http"
)

func main() {
	//render markdown to the screen
	mux := http.NewServeMux()
	mux.HandleFunc("GET /post/{slug}", PostHandler(FileReader{}))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

func PostHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		postMarkdown, err := sl.Read(slug)
		if err != nil {
			// TODO: Handle different errors in the future
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		fmt.Fprint(w, postMarkdown)
	}
}
