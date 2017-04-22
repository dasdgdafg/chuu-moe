package main

import (
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// local directories that assets are in
const imageDir = "images"
const htmlDir = "generated html"
const otherDir = "other assets"

var staticPages = map[string]string{
	"/rules":     htmlDir + "/rules.html",
	"/contact":   htmlDir + "/contact.html",
	"/rules.jpg": otherDir + "/86513819913bb4f8089f303e01f2dac3.jpeg",
}

var templates = template.Must(template.ParseFiles(htmlDir + "/index.html"))

func makeDefaultHandler(filenames []string) http.HandlerFunc {
	count := len(filenames)
	return func(w http.ResponseWriter, r *http.Request) {
		// serve index.html for /
		if r.URL.Path == "/" {
			//http.ServeFile(w, r, htmlDir+"/index.html")
			err := templates.ExecuteTemplate(w, "index.html", filenames[rand.Intn(count)])
			if err != nil {
				http.Error(w, "500 internal server error", 500)
			}
		} else {
			// serve one of the static pages if they exist
			page := staticPages[r.URL.Path]
			if page != "" {
				http.ServeFile(w, r, page)
			} else {
				http.NotFound(w, r)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	files, _ := ioutil.ReadDir(imageDir)
	filenames := make([]string, len(files))
	for i, f := range files {
		filenames[i] = f.Name()
	}
	files = nil

	imageHandler := http.FileServer(http.Dir(imageDir))
	defaultHandler := makeDefaultHandler(filenames)

	http.HandleFunc("/", defaultHandler)
	http.Handle("/images/", http.StripPrefix("/images/", imageHandler))
	http.ListenAndServe(":8080", nil)
}
