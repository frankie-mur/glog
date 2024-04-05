package main

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/home.go.tmpl")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, HomePageData{
		Posts: getPostTitles(),
	})
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
		var metadata Metadata

		rest, err := frontmatter.Parse(strings.NewReader(postMarkdown), &metadata)
		if err != nil {
			http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
			return
		}

		var buf bytes.Buffer
		err = mdRenderer.Convert([]byte(rest), &buf)
		if err != nil {
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
			return
		}
		tpl, err := template.ParseFiles("templates/post.go.tmpl")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, PostData{
			Metadata: metadata,
			Content:  template.HTML(buf.String()),
		})
	}
}
