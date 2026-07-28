package main

import (
	"database/sql"
	"encoding/json"
	// "github.com/lib/pq"
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
	http.HandleFunc("/getlinks", getLinksHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

type Resource struct {
	Url           string `json:"url"`
	ResourceId   string  `json:"resourceId"`
	ResourceName string  `json:"resourceName"`
}

func getLinksHandler(rw http.ResponseWriter, r *http.Request) {
	pg := PgConnection{}

	db, err := pg.getConnection()

	if err != nil {
		log.Fatal("pg connection failed")
	}

	rows, error := db.Query("SELECT resource_id, resource_name, url FROM public.resource")

	if error != nil {
		log.Fatal("problem with db query")
	}

	defer rows.Close()

	//I guess this inits the string array compared to using a compile time static value
	//slices are like references where regular arrays are copied by value
	resources := make([]Resource, 10)
	// https://pkg.go.dev/database/sql#DB.Query
	for rows.Next() {
		var resource_id, resource_name, url string
		if err := rows.Scan(&resource_id, &resource_name,&url); err != nil {
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
