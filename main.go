package main

import (
	"database/sql"

	dbAdapter "github.com/BrunoSSanta/hexagonal-go/adapters/db"
	"github.com/BrunoSSanta/hexagonal-go/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := dbAdapter.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 230)

	productService.Enable(product)

}
