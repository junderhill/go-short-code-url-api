package urlshortener

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
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
	fmt.Println(w, "%s %s", r.URL, slug)
}

func handleGetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func persistSlug(slug string, url string) {
	//create new record in dynamo db
}

func generateSlug() string {
	slugLength := 5
	return randomString(slugLength)
}

func randomString(n int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func slugAlreadyExists(slug string) bool {
	//check if slug exists in dynamo db
	return false
}

func getUrlBySlug(slug string) string {
	//get url from dynamo db

	client := getDynamoClient()

	out, err := client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String("url-shortener"),
		Key:       map[string]dynamodb.AttributeValue{"slug": {S: aws.String(slug)}}})

	return ""
}

func getDynamoClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc
}
