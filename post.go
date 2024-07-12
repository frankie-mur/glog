package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"

	"github.com/adrg/frontmatter"
)

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
	fmt.Printf("titles: %+v\n", titles)
	return titles
}
