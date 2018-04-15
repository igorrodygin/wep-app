package main

import (
	"net/http"
	_ "github.com/heroku/x/hmetrics/onload"
	"fmt"
	"os"
	"log"
)


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":" + port, nil)
}

