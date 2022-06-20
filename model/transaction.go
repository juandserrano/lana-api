package model

import (
)

type Transaction struct {
  Type		string `json:"type"`
  Name		string `json:"name"`
  Amount		float64 `json:"amount,string"`
  Id      int `json:"id,string"`
  AccountId int `json:"accountid,string"`
  Month string `json:"month"`
}

type Transactions []Transaction
