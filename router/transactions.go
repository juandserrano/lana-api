package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/juandserrano/lana-api/controller"
	"github.com/juandserrano/lana-api/model"
)

func NewTransaction(w http.ResponseWriter, r *http.Request){
  var transaction model.Transaction
  body, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  if err != nil {
    log.Fatalf("Error in ReadAll: %s", err)
  }
  err = json.Unmarshal([]byte(body), &transaction)
  if err != nil {
    fmt.Fprintf(w, "Error unmarshalling transaction!: %s", err)
    log.Fatalf("Error unmarshalling transaction: %s", err)
  }
  log.Printf("Unmarshalled Body: %s, %s, %s", transaction.Name, transaction.Date, transaction.Vendor)
  err = controller.NewTransaction(transaction)
  if err != nil {
    log.Fatalf("%s", err)
    fmt.Fprintf(w, "Error processing transaction!: %s", err)
  }

  fmt.Fprintf(w, "Added!")

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
