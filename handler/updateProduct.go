package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"pair-project/entity"
)

var (
	ErrNotFoundRecord         = fmt.Errorf("error record not found")
	ErrMultipleRecordAffected = fmt.Errorf("error multiple record affected")
)

func UpdateProduct(db *sql.DB, inputUpdatedProduct *entity.Products) error {
	query := `
		UPDATE products
		SET 
			name = ?,
			price = ?,
			stock = ?
		WHERE
			product_id = ?
		LIMIT 1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("error preparing sql statement:", err)
		return err
	}

	result, err := stmt.Exec(inputUpdatedProduct.Name, inputUpdatedProduct.Price, inputUpdatedProduct.Stock, inputUpdatedProduct.Product_id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			fmt.Println("error data not found")
		default:
			fmt.Println("error db exec:", err)
		}
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error getting row affected:", err)
		return err
	}

	if rowAffected != 1 {
		fmt.Println("error multiple row are affected:", err)
		return err
	}

	fmt.Println("\nProduct updated successfully")
	fmt.Println()

	return nil
}
