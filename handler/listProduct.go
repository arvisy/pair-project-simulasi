package handler

import (
	"database/sql"
	"pair-project/entity"
)

func ListProduct(db *sql.DB) ([]entity.Products, error) {
	var products []entity.Products

	query := "SELECT * FROM products"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Products
		err := rows.Scan(&product.Product_id, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
