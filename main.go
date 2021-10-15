package main

import (
	"log"
	"net/http"

	"./handler"
)

var (
	route handler.Handler = handler.NewGetHandler()
)

func main() {
	http.HandleFunc("/api/blp", route.BlpHandler)
	http.HandleFunc("/api/kasikorn", route.KAHandler)
	http.HandleFunc("/api/ks", route.KSHandler)
	http.HandleFunc("/api/kt", route.KTHandler)
	http.HandleFunc("/api/mfc", route.MFCHandler)
	http.HandleFunc("/api/scb", route.SCBHandler)
	http.HandleFunc("/api/aldn", route.AldnHandler)
	http.HandleFunc("/api/ins", route.InsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
