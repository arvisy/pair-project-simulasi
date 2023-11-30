package handler

import "database/sql"

type Handler struct {
	conn *sql.DB
}

func New(conn *sql.DB) *Handler {
	return &Handler{
		conn: conn,
	}
}
