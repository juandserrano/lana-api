package model

import (
)

type Transaction struct {
  Date		string `json:"date"`
  Name		string `json:"name"`
  Amount		float64 `json:"amount,string"`
  Category	string `json:"category"`
  Vendor		string `json:"vendor"`
  UUID      string `json:"uuid"`
}

type Transactions []Transaction
