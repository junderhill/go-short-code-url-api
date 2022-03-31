package urlshortener

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := 8080
	r := mux.NewRouter()

	r.HandleFunc("/", handleNewUrl).Methods("POST")
	r.HandleFunc("/{url}", handleGetUrl).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + string(port),
	}

	log.Fatal(srv.ListenAndServe())
}

func handleNewUrl(w http.ResponseWriter, r *http.Request) {

	// Get the URL from the request
	url := r.FormValue("url")

	//generate slugs until we find one that doesn't exist
	slug := generateSlug()
	for slugAlreadyExists(slug) {
		slug = generateSlug()
	}

	//save slug to database
	persistSlug(slug, url)

	//return slug

	w.WriteHeader(http.StatusOK)
}

func persistSlug(slug string, url string) {
	//create new record in dynamo db
}

func generateSlug() string {
	return ""
}

func slugAlreadyExists(slug string) bool {
	return true
}

func handleGetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
