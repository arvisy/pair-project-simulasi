package handler

import (
	"database/sql"
	"pair-project/entity"
)

func ListStaff(db *sql.DB) ([]entity.Staff, error) {
	var staffs []entity.Staff

	query := "SELECT staff_id, name, email, position FROM staff"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var staff entity.Staff
		err := rows.Scan(&staff.StaffID, &staff.Name, &staff.Email, &staff.Position)
		if err != nil {
			return nil, err
		}
		staffs = append(staffs, staff)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return staffs, nil
}
