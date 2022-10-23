package db

import (
	models "server/models"
)

// Driver is the database driver
type Driver interface {
	// creates the database table
	SetupTables() error

	// creates a new item
	// returns the id of the newly created item
	// returns an error if the operation fails
	CreateItem(item models.Item) error

	// returns the list of items
	// returns an error if the operation fails
	GetItems() ([]models.Item, error)

	// returns the item with the specified id
	// returns an error if the operation fails
	GetItem(id int) (models.Item, error)

	// deletes the item with the specified id
	// returns an error if the operation fails
	DeleteItem(id int) error

	// updates the item with the specified id
	// returns an error if the operation fails
	UpdateItem(id int, item models.Item) error

	// closes the database connection
	Close() error
}