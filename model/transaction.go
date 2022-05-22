package model

import (
)

type Transaction struct {
	Date		string
	Name		string
  Amount		float64 `json:"amount,string"`
	Category	string
	Vendor		string
}

type Transactions []Transaction
