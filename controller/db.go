package controller

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/juandserrano/lana-api/model"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDB(){
  host := "postgres"
  port := "5432"
  user := "postgres"
  password := os.Getenv("POSTGRES_PASSWORD")
  dbname := "lana"

  connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)

  db, err := sql.Open("postgres", connInfo)
  if err != nil {
    log.Fatalf("Error connecting to Postgres: %s", err)
  }
  defer db.Close()

  checkDBExists(db, dbname)
  _, err = db.Exec("CREATE TABLE IF NOT EXISTS transactions (id serial PRIMARY KEY, name varchar(255) NOT NULL, amount DECIMAL NOT NULL, category varchar(255), vendor varchar(255), date DATE NOT NULL DEFAULT CURRENT_DATE)")
  if err != nil {
    log.Fatalf("Error creating table: %s", err)
  }
  query := "select name FROM transactions"
  data, err := db.Query(query)
  if err != nil {
    log.Fatalf("Error querying to Database: %s", err)
  }
  var transaction model.Transaction
  s, _ := data.Columns()
  fmt.Printf("Columns: %s", s) 
  for data.Next() {
    data.Scan(&transaction.Name)
    fmt.Printf("Name is: %s", transaction.Name)
  }
}

func checkDBExists(db *sql.DB, dbname string){
  _, err := db.Exec("SELECT 'CREATE DATABASE" + dbname + "' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '" + dbname + "')")
  if err != nil {
    log.Fatal("Error creating DB")
  }
}

func NewTransaction(t model.Transaction) error {
  _, err := db.Exec("INSERT INTO transactions (name, amount, category, vendor, date) VALUES (%s, %s, %s, %s, %s)",
    t.Name, t.Amount, t.Category, t.Vendor, t.Date)
    if err != nil {
      log.Fatalf("Error insertion into database: %s", err)
      return err
    }
  return nil
}
