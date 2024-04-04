package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
)

func main() {
	//render markdown to the screen
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /post/{slug}", PostHandler(FileReader{}))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}

}

type HomePageData struct {
	Posts []PostTitle
}

type PostTitle struct {
	Title string
	Slug  string
}

type PostData struct {
	Metadata Metadata
	Content  template.HTML
}

type Metadata struct {
	Title       string   `yaml:"title"`
	Author      string   `yaml:"author"`
	Description string   `yaml:"description"`
	Date        string   `yaml:"date"`
	Tags        []string `yaml:"tags"`
}

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

func getPostTitles() []PostTitle {
	files, err := os.ReadDir("posts")
	if err != nil {
		panic(err)
	}
	var titles []PostTitle
	for _, file := range files {
		//open file
		f, err := os.Open(fmt.Sprintf("posts/%s", file.Name()))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		//read file
		b, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		//parse frontmatter
		var metadata Metadata
		_, err = frontmatter.Parse(strings.NewReader(string(b)), &metadata)
		if err != nil {
			panic(err)
		}
		var postTile = PostTitle{
			Title: metadata.Title,
			Slug:  fmt.Sprintf("/post/%s", strings.Split(file.Name(), ".")[0]),
		}
		titles = append(titles, postTile)
	}
	return titles
}
