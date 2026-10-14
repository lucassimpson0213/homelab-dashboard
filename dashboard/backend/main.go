package main

import (
	"fmt"
	"encoding/json"
	"homelab-dashboard/db"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	_ "github.com/lib/pq"
)

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080,http://127.0.0.1,http://100.65.51.13")
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
	http.HandleFunc("/api/postlink", postLinkHandler)
	//serves static routing
	http.HandleFunc("/{$}", staticHandler)

	// serves all static files like html, css and js
	http.Handle("/",
		http.StripPrefix("/dist/", static),
	)
	
	//listen on all interfaces, depending on if you wanna do local or cross machine dev
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postLinkHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
	}

	database, err := db.New()

	if err != nil {
		errorStr := fmt.Sprintf("There has been an issue initializing the database, postgres cannot be queried: %s", err.Error())
		slog.Error(errorStr)
		http.Error(rw, "Database failed to be initialized, postgres cannot be queried", http.StatusServiceUnavailable)
		return 
	}

	var resource db.Resource
	fmt.Println(r.Body)
	decodeErr := json.NewDecoder(r.Body).Decode(&resource)
	// creating resource
	if decodeErr != nil {
		// return fmt.Errorf("There has been an error decoding the resource: %w", "decodeErr", decodeErr)
		errorStr := fmt.Sprintf("There has been an error decoding the resource: %s", decodeErr.Error())
	    slog.Error(errorStr)		
		http.Error(rw, "There has been an error decoding the resource", http.StatusInternalServerError)
		return 
	}

	result, err := database.CreateResource(resource)
	slog.Debug("creating resource: ")

	if err != nil {
		errorStr := fmt.Sprintf("There has been an error creating the resource: %s", err.Error())
		slog.Error(errorStr)
	}
	rowsaffected, rowsaerror := result.RowsAffected()

	if rowsaerror != nil {
		http.Error(rw, "failed to get rows affected from database", http.StatusInternalServerError)
	}
	rowsaffectedstr := strconv.FormatInt(rowsaffected, 10)
	slog.Debug(rowsaffectedstr)

	encoded, err := json.Marshal(rowsaffectedstr)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(encoded)

}
func getLinksHandler(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
		return
	}

	repository, err := db.New()

	if err != nil {
		slog.Error("Error occured while calling db.New(): ", "err", err)
	}

	resources, repoerr := repository.GetResources()

	if repoerr != nil {
		slog.Error("Error occured while calling repository.GetResources(): ", "repoerr", repoerr)
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

