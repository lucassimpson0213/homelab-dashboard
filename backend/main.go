package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	_ "path/filepath"

	_ "github.com/lib/pq"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080,http://127.0.0.1")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("x-custom-header", "Hello")
}


var templates = template.Must(template.ParseFiles("/home/lucassimpson/DEV/homelab-dashboard/frontend/dist/index.html"))

type PageData struct {
	Path string
}

func main() {
	static := http.FileServer(http.Dir("../frontend/dist"))
	http.HandleFunc("/api/getlinks", getLinksHandler)
	// http.HandleFunc("/api/postlink", postLinkHandler)
	//serves static routing
	http.HandleFunc("/{$}", staticHandler)

	// serves all static files like html, css and js
	http.Handle("/",
		http.StripPrefix("/dist/", static),
	)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type Resource struct {
	Url          string `json:"url"`
	ResourceId   string `json:"resourceId"`
	ResourceName string `json:"resourceName"`
}

func postLinkHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
	}

}
func getLinksHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
		return
	}
	pg := PgConnection{}

	db, err := pg.getConnection()

	if err != nil {
		log.Fatal("pg connection failed")
	}

	rows, error := db.Query("SELECT resource_id, resource_name, url FROM public.resource")

	if error != nil {
		log.Println(error)
		log.Fatal("problem with db query")
	}

	defer rows.Close()

	//I guess this inits the string array compared to using a compile time static value
	//slices are like references where regular arrays are copied by value
	resources := make([]Resource, 10)
	// https://pkg.go.dev/database/sql#DB.Query
	for rows.Next() {
		var resource_id, resource_name, url string
		if err := rows.Scan(&resource_id, &resource_name, &url); err != nil {
			log.Print("There has been an error with the rows scanner for the database; in get links handler")
			log.Fatal(err)
		}
		resource := Resource{ResourceId: resource_id, ResourceName: resource_name, Url: url}

		resources = append(resources, resource)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	b, error := json.Marshal(resources)
	if error != nil {
		log.Fatal(error)
	}

	rw.Write(b)
	rw.WriteHeader(http.StatusOK)

}
func linkHandler(rw http.ResponseWriter, r *http.Request) {

}
func staticHandler(rw http.ResponseWriter, r *http.Request) {
	cors(rw)

	if r.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
		return
	}

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

// TODO make repoistory for database to decouple the database
func (pg *PgConnection) getConnection() (*sql.DB, error) {
	if pg.connect != nil {
		return pg.connect, nil
	}
	db, err := sql.Open("postgres", "host=localhost dbname=postgres connect_timeout=5 sslmode=disable user=lucas password=lucas")
	if err != nil {
		log.Fatal("postgres failed to init")
		return nil, err
	}

	pg.connect = db
	return pg.connect, nil

}
