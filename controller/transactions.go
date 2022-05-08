package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/juandserrano/lana-api/model"
)

func Transactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showTransactions(w)
		break
	case "POST":
		handlePost()
		break
	case "UPDATE":
		updateTransaction()
		break
	default:
		fmt.Println("Error code")
	}
}

func newTransaction(){
	fmt.Println("New transaction")
}

func handlePost() {
	fmt.Println("Handle post")
}

func deleteTransaction(){
	fmt.Println("delete transaction")
}

func updateTransaction(){
	fmt.Println("update transaction")
}

func showTransactions(w http.ResponseWriter){
	transactions := model.Transactions{
		model.Transaction{
			Name: "UberEats",
			Amount: 69.9,
			Date: time.Now(),
			Vendor: "Shawarma",
			Category: "Restaurants",
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		panic(err)
	}
	fmt.Println("Show transaction console")
}