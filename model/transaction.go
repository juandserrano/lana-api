package model

import (
	"time"
)

type Transaction struct {
	Date		time.Time	`json:"date"`
	Name		string		`json:"name"`
	Amount		float64		`json:"amount"`
	Category	string		`json:"category"`
	Vendor		string		`json:"vendor"`
}

type Transactions []Transaction