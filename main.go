package main

import (
	"FrequenctCountOfTexts/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/topTenWords", handler.TopTenMostUsedWords).Methods("POST")
	log.Fatal(http.ListenAndServe(":8098", router))
}

