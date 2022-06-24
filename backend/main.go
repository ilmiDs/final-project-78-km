package main

import (
	"database/sql"

	"github.com/ilmiDs/final-project-78-km/api"
	"github.com/ilmiDs/final-project-78-km/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "backend/database/assigment/cashier-app/db/cashier-app.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)
	// productsRepo := repository.NewProductRepository(db)
	// cartItemRepo := repository.NewCartItemRepository(db)
	// salesRepo := repository.NewSalesRepository(db)
	// transactionRepo := repository.NewTransactionRepository(db, *productsRepo, *cartItemRepo)

	mainAPI := api.NewAPI(*usersRepo)
	mainAPI.Start()
}
