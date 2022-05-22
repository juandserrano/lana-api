package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/juandserrano/lana-api/controller"
	"github.com/juandserrano/lana-api/model"
)

func NewTransaction(w http.ResponseWriter, r *http.Request){
  var transaction model.Transaction
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    log.Fatalf("Error in ReadAll: %s", err)
    return
  }
  defer r.Body.Close()
  log.Printf("Receiving Body: %s", body)

  err = json.Unmarshal(body, &transaction)
  if err != nil {
    log.Fatalf("Error unmarshalling transaction: %s", err)
    return
  }
  log.Printf("Unmarshalled Body: %T, , %T, %T, %T", transaction.Name, transaction.Amount, transaction.Date, transaction.Vendor)

  err = controller.NewTransaction(controller.GetDB(), transaction)
  if err != nil {
    log.Fatalf("Error on database: %s", err)
    return
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
			Date: "10/10.2010",
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
