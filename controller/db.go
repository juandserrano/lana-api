package controller

import (
	"fmt"
	"os"
	"strings"

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
  monthid := getMonthId(t.Month)
  typeid := getTypeId(t.Type)
  sqlStatement := `INSERT INTO "transactions" ("id", "name", "amount", "accountid", "monthid", "typeid") VALUES ($1, $2, $3, $4, $5, $6)`
  _, err := db.Exec(sqlStatement, t.Id, t.Name, t.Amount, t.AccountId, monthid, typeid)

  if err != nil {
    return err
  }
  return nil
}

func GetAllTransactions(db *sql.DB) (model.Transactions, error) {

  err := db.Ping()
  if err != nil {
    return nil, err
  }


query := `SELECT transactions.id, transactions.name, amount, accountid, months.name, types.name FROM transactions
INNER JOIN months ON transactions.monthid = months.id
INNER JOIN types ON transactions.typeid = types.id;`
  rows, err := db.Query(query)
  tList := model.Transactions{}

  defer rows.Close()
  for rows.Next() {
    var name, month, ttype string
    var accountid, id int
    var amount float64
    err = rows.Scan(&id, &name, &amount, &accountid, &month, &ttype)
    if err != nil {
      return nil, err
    }
    tList = append(tList, model.Transaction{
      Name: name,
      Id: id,
      Amount: amount,
      AccountId: accountid,
      Month: month,
      Type: ttype,
    })
  }

  return tList, nil;
}

func getMonthId(month string) int {
  var monthid int
  switch month {
  case "January":
    monthid = 1
  case "February":
    monthid = 2
  case "March":
    monthid = 3
  case "April":
    monthid = 4
  case "May":
    monthid = 5
  case "June":
    monthid = 6
  case "July":
    monthid = 7
  case "August":
    monthid = 8
  case "September":
    monthid = 9
  case "October":
    monthid = 10
  case "November":
    monthid = 11
  case "December":
    monthid = 12
  }
return monthid
}

func getTypeId(ttype string) int {
  if strings.Contains(ttype, "income") {
    return 1
  }
  return 2
}
