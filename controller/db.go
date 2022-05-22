package controller

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/juandserrano/lana-api/model"
	_ "github.com/lib/pq"
)
var db *sql.DB
func ConnectToDB() (error){
  host := "postgres"
  port := "5432"
  user := "postgres"
  password := os.Getenv("POSTGRES_PASSWORD")
  dbname := "lana"

  connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

  var err error
  db, err = sql.Open("postgres", connInfo)
  if err != nil {
    return err
  }

  return nil
}

func GetDB() (*sql.DB, error){
  err := db.Ping()
  if err != nil {
    return nil, err
  }
  return db, nil
}

func NewTransaction(db *sql.DB, t model.Transaction) error {
  sqlStatement := `INSERT INTO "transactions" ("name", "amount", "category", "vendor", "date") VALUES ($1, $2, $3, $4, $5)`
  _, err := db.Exec(sqlStatement, t.Name, t.Amount, "category", t.Vendor, t.Date)

  if err != nil {
    return err
  }
  return nil
}
