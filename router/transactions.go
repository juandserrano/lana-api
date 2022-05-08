package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	
	"github.com/juandserrano/lana-api/model"
)

func NewTransaction(){
	
	fmt.Println("New transaction")
}

func HandlePost() {
	fmt.Println("Handle post")
}

func DeleteTransaction(){
	fmt.Println("delete transaction")
}

func UpdateTransaction(){
	fmt.Println("update transaction")
}

func ShowTransactions(w http.ResponseWriter, r *http.Request){
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