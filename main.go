package main

import (
	"database/sql"
	"github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var templates = template.Must(template.ParseFiles("./frontend/dist/index.html"))

type PageData struct {
	Path string
}

func main() {

	static := http.FileServer(http.Dir("./frontend/dist"))

	http.Handle("/dist/",
		http.StripPrefix("/dist/", static),
	)

	http.HandleFunc("/", staticHandler)
	http.HandleFunc("/getlink", getLinksHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
func getLinksHandler(rw http.ResponseWriter, r *http.Request) {
	pg := PgConnection{}

	db, err := pg.getConnection()

	if err != nil {
		log.Fatal("pg connection failed")
	}

	db.Ping()
}
func linkHandler(rw http.ResponseWriter, r *http.Request) {

}
func staticHandler(rw http.ResponseWriter, r *http.Request) {

	dat := PageData{
		Path: r.URL.Path,
	}

	if err := templates.Execute(rw, dat); err != nil {
		log.Printf("template error: %v", err)
		http.Error(rw, "failed to render page", http.StatusInternalServerError)
		return
	}
}
func rootHandler(rw http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(rw, "index.html", nil)
}

type PgConnection struct {
	connect *sql.DB
}

func (pg *PgConnection) getConnection() (*sql.DB, error) {
	if pg.connect != nil {
		return pg.connect, nil
	}
	db, err := sql.Open("postgres", "host=localhost dbname=resource connect_timeout=5")
	if err != nil {
		log.Fatal("postgres failed to init")
		return nil, err
	}

	pg.connect = db
	return pg.connect, nil

}
