package handler

import (
	"database/sql"
	"errors"
	"pair-project/entity"
)

func (h *Handler) UpdateProduct(product *entity.Products) (*entity.Products, error) {
	query := `
		UPDATE products
		SET 
			name = ?,
			price = ?,
			stock = ?
		WHERE
			id = ?
		LIMIT 1
	`

	result, err := h.conn.Exec(query, product.Name, product.Price, product.Stock, product.Product_id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFoundRecord
		default:
			return nil, err
		}
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowAffected != 1 {
		return nil, ErrMultipleRecordAffected
	}

	return product, nil
}
