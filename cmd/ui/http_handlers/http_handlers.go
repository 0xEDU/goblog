package http_handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type HomeData struct {
	Title string
}

func getStaticFilePath(fileType string, template string) (string, error) {
	e, err := os.Executable()
	if err != nil {
		return "", err
	}
	wd := filepath.Dir(e)
	return filepath.Join(wd, "static", fileType, template), nil
}

func getPosts() []

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeData := HomeData{"Hello, world!"}
	homePath, err := getStaticFilePath("html", "home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles(homePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, homeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Serving home page")
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	cssFile := strings.TrimPrefix(r.URL.Path, "/css/")
	cssPath, err := getStaticFilePath("css", cssFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Println("Serving css file: " + cssFile)
	http.ServeFile(w, r, cssPath)
}
