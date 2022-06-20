package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

)

func GetAccountSummary(w http.ResponseWriter, r *http.Request){

  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    log.Fatalf("Error in ReadAll: %s", err)
    return
  }

  defer r.Body.Close()

  var accountNumber int
  err = json.Unmarshal(body, &accountNumber)
  if err != nil {
    log.Fatalf("Error unmarshalling accountNumber: %s", err)
    return
  }
  
  

}
