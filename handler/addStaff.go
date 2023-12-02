package handler

import (
	"database/sql"
	"fmt"
	"log"
	"pair-project/entity"
)

func AddStaff(db *sql.DB, input entity.AddStaffInput) error {
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

	newStaff := entity.Staff{
		StaffID:  id,
		Name:     input.Name,
		Email:    input.Email,
		Position: input.Position,
	}

	fmt.Printf("\nStaffID | Name | Email | Position\n")
	fmt.Println("----------------------------------------------")

	fmt.Printf("%v | %v | %v | %v\n", newStaff.StaffID, newStaff.Name, newStaff.Email, newStaff.Position)
	fmt.Println("")

	return nil
}
