package storage

import (
	models "server/models"
)

// `Service` is the interface that provides the storage operations.
type Service interface {
	// `CreateItem` creates a new item in the storage.
	CreateItem(item models.Item) error

	// `GetItems` returns the items from the storage.
	GetItems() ([]models.Item, error)

	// `GetItem` returns the item with the specified ID from the storage.
	GetItem(id int) (models.Item, error)

	// `UpdateItem` updates the item with the specified ID in the storage.
	UpdateItem(id int, item models.Item) error

	// `DeleteItem` deletes the item with the specified ID from the storage.
	DeleteItem(id int) error
}