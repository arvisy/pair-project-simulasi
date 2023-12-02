package handler

import (
	"database/sql"
	"errors"
	"fmt"
)

func GetProductIDByName(db *sql.DB, productName string) (int, error) {
	query := `
		SELECT product_id
		FROM products
		WHERE name = ?
		LIMIT 1
	`

	var productID int64
	err := db.QueryRow(query, productName).Scan(&productID)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFoundRecord):
			return -1, ErrNotFoundRecord
		default:
			return -1, fmt.Errorf("error getting product by name: %v", err)
		}
	}

	return int(productID), nil
}
