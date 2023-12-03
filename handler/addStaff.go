package handler

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
)

func AddStaff(db *sql.DB, input entity.Staff) error {
	query :=
		`INSERT INTO staff (name, email, position)
        VALUES (?, ?, ?)`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("error preparing sql statement:", err)
		return err
	}

	result, err := stmt.Exec(input.Name, input.Email, input.Position)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("error getting :", err)
		return err
	}

	fmt.Println("Product added successfully with id: ", id)

	return nil
}
