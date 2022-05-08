package main

import (
	"log"
	"net/http"

	"github.com/juandserrano/lana-api/controller"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/api/transactions/", controller.Transactions)
	log.Fatal(http.ListenAndServe(":3003",nil))
}