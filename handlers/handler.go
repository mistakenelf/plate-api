package handlers

import "database/sql"

// Handler adds db to a handler
type Handler struct {
	DB *sql.DB
}
