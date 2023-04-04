package repository

import (
	"database/sql"

	"github.com/Quoc-Bao-213/getgo-project/api/internal/repository/inventory"
)

type Registry interface {
	// Inventory returns the inventory repo
	Inventory() inventory.Repository
}

// New returns an implementation instance which satisfying Registry
func New(db *sql.DB) Registry {
	return impl{
		inventory: inventory.New(db),
	}
}

type impl struct {
	inventory inventory.Repository
}

// Inventory returns the inventory repo
func (i impl) Inventory() inventory.Repository {
	return i.inventory
}
