package main

import (
	"fmt"

	// "log"

	"html/template"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var templates = template.Must(template.ParseFiles("frontend/html/index.html"))

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", rootHandler)

}

func rootHandler(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "index.html", nil)
}
