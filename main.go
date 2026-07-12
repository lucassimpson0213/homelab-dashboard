package main

import (

	// "log"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var templates = template.Must(template.ParseFiles("dist/index.html"))

type PageData struct {
	Path string
}

func main() {
	workingdir, err := os.Getwd()
	check(err)

	fmt.Println("working dir", workingdir)
	http.HandleFunc("/dist/", staticHandler)
	http.HandleFunc("/", rootHandler)
	error := http.ListenAndServe(":8080", nil)
	check(error)
}
func staticHandler(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println("path: ", path)
	dat := PageData{
		Path: path,
	}

	if err := templates.ExecuteTemplate(rw, "index.html", dat); err != nil {
		http.Error(rw, "failed to render page", http.StatusInternalServerError)
	}

}
func rootHandler(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "index.html", nil)
}
