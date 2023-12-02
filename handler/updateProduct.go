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

func UpdateProduct(db *sql.DB, oldproductName string, inputUpdatedProduct *entity.UpdateProductInput) error {
	productID, err := GetProductIDByName(db, oldproductName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	updatedProduct := entity.Products{
		Product_id: productID,
		Name:       inputUpdatedProduct.Name,
		Price:      inputUpdatedProduct.Price,
		Stock:      inputUpdatedProduct.Stock,
	}

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

	result, err := stmt.Exec(updatedProduct.Name, updatedProduct.Price, updatedProduct.Stock, updatedProduct.Product_id)
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

	fmt.Printf("\nProductID | ProductName | ProductPrice | ProductStock |\n")
	fmt.Println("----------------------------------------------")
	fmt.Printf("%v | %v | %v | %v\n", updatedProduct.Product_id, updatedProduct.Name, updatedProduct.Price, updatedProduct.Stock)
	fmt.Println("")

	return nil
}
