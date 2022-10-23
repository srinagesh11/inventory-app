package storage

import (
	"database/sql"
	db "server/db"
	models "server/models"
)

// `service` implements the `Service` interface
type service struct {
	driver db.Driver
}

// NewService creates a new service
func NewService(driver db.Driver) Service {
	return &service{driver: driver}
}

// `CreateItem` creates a new item in the storage.
func (s *service) CreateItem(item models.Item) error {
	return s.driver.CreateItem(item)
}

// `GetItems` returns the items from the storage.
func (s *service) GetItems() ([]models.Item, error) {
	items, err := s.driver.GetItems()
	if err != nil {
		if err == sql.ErrNoRows {
			return []models.Item{}, nil
		}
		return nil, err
	}
	return items, nil
}

// `GetItem` returns the item with the specified ID from the storage.
func (s *service) GetItem(id int) (models.Item, error) {
	return s.driver.GetItem(id)
}

// `UpdateItem` updates the item with the specified ID in the storage.
func (s *service) UpdateItem(id int, item models.Item) error {
	return s.driver.UpdateItem(id, item)
}

// `DeleteItem` deletes the item with the specified ID from the storage.
func (s *service) DeleteItem(id int) error {
	return s.driver.DeleteItem(id)
}
