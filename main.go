package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
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
		mdRenderer := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("dracula"),
				),
			),
		)
		if err != nil {
			http.Error(w, "Error creating markdown renderer", http.StatusInternalServerError)
			return
		}
		var buf bytes.Buffer
		err = mdRenderer.Convert([]byte(postMarkdown), &buf)
		if err != nil {
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
			return
		}
		if err != nil {
			// TODO: Handle different errors in the future
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		io.Copy(w, &buf)
	}
}
