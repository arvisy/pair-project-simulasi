package handler

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
)

func SalesRecap(db *sql.DB, date1 string, date2 string) error {
	fmt.Printf("\nGetting sales from %s to %s.\n", date1, date2)

	// sql query for recap
	query := `SELECT sale_id, name, quantity, sale_date FROM sales JOIN products ON products.product_id = sales.product_id WHERE sale_date >= ? AND sale_date <= ?`

	// prepare query
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("error preparing sql statement:", err)
		return err
	}

	// run query w/arguments
	rows, err := stmt.Query(date1, date2)
	if err != nil {
		log.Println("error in sql query:", err)
		return err
	}
	defer rows.Close()

	var sales []entity.Recap

	// scan query result
	for rows.Next() {
		var sl entity.Recap
		err := rows.Scan(&sl.Sale_id, &sl.Product_name, &sl.Quantity, &sl.Sale_date)
		if err != nil {
			log.Println("error scanning rows:", err)
			return err
		}
		sales = append(sales, sl)
	}

	if err := rows.Err(); err != nil {
		log.Println("error in row results:", err)
		return err
	}

	// print result
	fmt.Printf("\nSaleID | ProductName | Quantity | SaleDate\n")
	fmt.Println("----------------------------------------------")

	for _, s := range sales {
		fmt.Printf("%d | %s | %d | %s\n", s.Sale_id, s.Product_name, s.Quantity, s.Sale_date)
	}

	return nil
}
