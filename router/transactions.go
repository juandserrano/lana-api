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

  db, _ := controller.GetDB()
  err = controller.NewTransaction(db, transaction)
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

  db, _ := controller.GetDB()
  tList, err := controller.GetAllTransactions(db)
  if err != nil {
    log.Fatalf("Error on database: %s", err)
    return
  }

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tList); err != nil {
		panic(err)
	}
}
