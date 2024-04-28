package model

import (
	"github.com/areviksol/backend_task/database"
)

type Model struct {
	DB *database.Database
}

func NewModel(db *database.Database) *Model {
	return &Model{DB: db}
}

func (m *Model) CheckRecord(identifier string) (bool, error) {
	return m.DB.CheckRecord(identifier)
}

func (m *Model) AddRecord(identifier string) error {
	return m.DB.AddRecord(identifier)
}
