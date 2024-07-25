package main

import (
	"database/sql"
)

type Store struct {
	// Users 
	CreateUser() error
}


// holds methods to communicate with db 
type Storage struct {
	db *sql.DB
}


// create constructor for store 
func NewStore(db *sql.DB) *Storage {
	return &Storage {
		db: db,
	}
}

func (s *Storage) CreateUser() error {
	return nil
}