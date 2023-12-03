package handler

import (
	"database/sql"
	"fmt"
	"pair-project/entity"
)

func AddProduct(db *sql.DB, product entity.Products) error {
	query := `INSERT INTO products (name, price, stock) VALUES (?, ?, ?)`

	_, err := db.Exec(query, product.Name, product.Price, product.Stock)
	if err != nil {
		return err
	}

	fmt.Println("Product added successfully")
	return nil

}
