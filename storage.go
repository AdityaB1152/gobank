package main

import "database/sql"

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
	UpdateAccount(*Account) error
}

type PostgresStore struct {
	db *sql.DB
}
