package handler

import (
	"database/sql"
	"pair-project/entity"
)

func AddProduct(db *sql.DB, product entity.Products) error {
	_, err := db.Exec("INSERT INTO products (name, price, stock) VALUES (?, ?, ?)", product.Name, product.Price, product.Stock)
	return err
}
