package db

import (
	"database/sql"
	"fmt"
	config "server/config"
	models "server/models"

	_ "github.com/go-sql-driver/mysql"
)

type driver struct {
	db *sql.DB
}

func NewDriver(cfg *config.Config) (Driver, error) {

	fmt.Println("Creating a new database driver with the following configuration:")
	fmt.Println("DB_USER: ", cfg.DbUser)
	fmt.Println("DB_PASSWORD: ", "********")
	fmt.Println("DB_NAME: ", cfg.DbName)

	// create the database connection
	connString := "root" + ":" + "srinagesh" + "@tcp(" + cfg.DbHost + ":3306)/" + cfg.DbName
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return &driver{db: db}, nil
}

// setup the tables
func (d *driver) SetupTables() error {
	query := `
		CREATE TABLE IF NOT EXISTS items (
			id INT NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			quantity INT NOT NULL,
			unit_price FLOAT NOT NULL,
			PRIMARY KEY (id)
		)
	`
	result, err := d.db.Query(query)
	if err != nil {
		return err
	}
	defer result.Close()
	return nil
}

// Close closes the database connection
func (d *driver) Close() error {
	return d.db.Close()
}

// creates a new item
// returns the id of the newly created item
// returns an error if the operation fails
func (d *driver) CreateItem(item models.Item) error {
	query := "INSERT INTO items (id, name, quantity, unit_price) VALUES (?, ?, ?, ?)"
	result, err := d.db.Query(query, item.Id, item.Name, item.Quantity, item.UnitPrice)
	if err != nil {
		return err
	}
	defer result.Close()
	return nil
}

// returns the list of items
// returns an error if the operation fails
func (d *driver) GetItems() ([]models.Item, error) {
	query := "SELECT id, name, quantity, unit_price FROM items"
	result, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	items := []models.Item{}
	for result.Next() {
		var item models.Item
		err = result.Scan(&item.Id, &item.Name, &item.Quantity, &item.UnitPrice)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// returns the item with the specified id
// returns an error if the operation fails
func (d *driver) GetItem(id int) (models.Item, error) {
	query := "SELECT * FROM items WHERE id = ?"
	result, err := d.db.Query(query, id)
	if err != nil {
		return models.Item{}, err
	}
	defer result.Close()
	if result.Next() {
		var item models.Item
		if err := result.Scan(&item.Id, &item.Name, &item.Quantity, &item.UnitPrice); err != nil {
			return models.Item{}, err
		}
		return item, nil
	}
	return models.Item{}, nil
}

// updates the item with the specified id
// returns an error if the operation fails
func (d *driver) UpdateItem(id int, item models.Item) error {
	query := "UPDATE items SET name = ?, quantity = ?, unit_price = ? WHERE id = ?"
	result, err := d.db.Query(query, item.Name, item.Quantity, item.UnitPrice, id)
	if err != nil {
		return err
	}
	defer result.Close()
	return nil
}

// deletes the item with the specified id
// returns an error if the operation fails
func (d *driver) DeleteItem(id int) error {
	query := "DELETE FROM items WHERE id = ?"
	result, err := d.db.Query(query, id)
	if err != nil {
		return err
	}
	defer result.Close()
	return nil
}
