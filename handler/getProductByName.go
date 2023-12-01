package handler

import (
	"database/sql"
	"errors"
	"pair-project/entity"
)

func GetProductByName(db *sql.DB, productName string) (*entity.Products, error) {
	query := `
		SELECT product_id, name, price, stock
		WHERE name = ?
		LIMIT 1
	`

	var product entity.Products
	err := db.QueryRow(query, productName).Scan(
		&product.Product_id,
		&product.Name,
		&product.Price,
		&product.Stock,
	)

	if err != nil {
		switch {
		case errors.Is(err, ErrNotFoundRecord):
			return nil, ErrNotFoundRecord
		default:
			return nil, err
		}
	}

	return &product, nil
}
