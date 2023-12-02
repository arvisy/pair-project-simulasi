package handler

import (
	"database/sql"
	"pair-project/entity"
)

func AddStaff(db *sql.DB, input entity.AddStaffInput) (entity.Staff, error) {
	query :=
		`INSERT INTO staff (name, email, position)
        VALUES (?, ?, ?)`

	result, err := db.Exec(query, input.Name, input.Email, input.Position)
	if err != nil {
		return entity.Staff{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return entity.Staff{}, err
	}

	newStaff := entity.Staff{
		StaffID:  id,
		Name:     input.Name,
		Email:    input.Email,
		Position: input.Position,
	}

	return newStaff, nil
}
