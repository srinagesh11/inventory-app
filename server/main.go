package main

import (
	"fmt"
	"os"
	config "server/config"
	db "server/db"
	routes "server/routes"
	storage "server/storage"
)

func main() {
	// get the configuration
	cfg := config.GetConfig()

	// create a new database driver
	driver, err := db.NewDriver(cfg)
	if err != nil {
		fmt.Println("Error creating the database driver: ", err)
		os.Exit(1)
	}
	fmt.Println("Database driver created successfully")
	defer driver.Close()

	// setup the tables
	// time.Sleep(5 * time.Second)
	// if err := driver.SetupTables(); err != nil {
	// 	fmt.Println("Error setting up the tables: ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Tables created successfully")

	// create a new service
	service := storage.NewService(driver)

	// create the routes
	router := routes.RegisterRoutes(service)

	// Listen and serve on http://localhost:8080
	router.Run(":" + cfg.ServerPort)
}
