package handler

import "pair-project/entity"

func (h *Handler) AddStaff(newStaff *entity.Staff) (*entity.Staff, error) {
	query := `
		INSERT INTO staff (name, email, position)
		VALUES (?, ?, ?)
	`

	result, err := h.conn.Exec(query, newStaff.Name, newStaff.Email, newStaff.Position)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	newStaff.StaffID = id
	return newStaff, nil
}
