package main

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

var db *sql.DB

type Product struct {
  Id int
  Name string
  Description string
  Price int
  Quantity int
}

type Products struct {
  Products []Product
}
func main() {
  var err error

  db, err = sql.Open("postgres", "host=127.0.0.1 user=api password=123456 dbname=api sslmode=disable")
  if err != nil {
    panic(err)
  }
  
  defer db.Close()

  fmt.Println("# Starting server...")

  http.HandleFunc("/v1/products/", getProducts)
  log.Fatal(http.ListenAndServe(":8080", nil))

func getProducts(w http.ResponseWriter, r *http.Request){
  w_array := Products{}

  fmt.Println("# Querying...")
  rows, err := db.Query("SELECT id, ")
}

}
