package model

import (
)

type Account struct {
  Number		string `json:"number"`
  Username		string `json:"username"`
  Balance		float64 `json:"balance,string"`
  Spent	string `json:"spent,string"`
  Budgeted		string `json:"budgeted,string"`
}
