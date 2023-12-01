package handler

import (
	"database/sql"
	"fmt"
)

var (
	ErrNotFoundRecord         = fmt.Errorf("error record not found")
	ErrMultipleRecordAffected = fmt.Errorf("error multiple record affected")
)

type Handler struct {
	conn *sql.DB
}

func New(conn *sql.DB) *Handler {
	return &Handler{
		conn: conn,
	}
}
