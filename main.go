package main

import (
	"log"
	"net/http"

	"github.com/Chris-Ju/FISCO-BCOS-DEMO/router"
)

func main() {
	log.Printf("Server started")
	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}