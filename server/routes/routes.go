package routes

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	models "server/models"
	storage "server/storage"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// creates the rotues
func RegisterRoutes(service storage.Service) *gin.Engine {

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// creates a new item
	router.POST("/api/item", func(c *gin.Context) {
		var item models.Item
		err := c.BindJSON(&item)
		if err != nil {
			handleError(c, err)
			return
		}

		err = validateItem(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if item.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "The item id is required",
			})
			return
		}

		if err := service.CreateItem(item); err != nil {
			handleError(c, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"item": item,
		})
	})

	// get the list of items
	router.GET("/api/item", func(c *gin.Context) {
		items, err := service.GetItems()
		if err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	})

	// get item with specific id
	router.GET("/item/:id", func(c *gin.Context) {
		id := c.Param("id")
		itemId, err := validateItemId(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		item, err := service.GetItem(itemId)
		if err != nil {
			handleError(c, err)
			return
		}
		if item.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Item not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"item": item,
		})
	})

	// delete item with specific id
	router.DELETE("/api//item/:id", func(c *gin.Context) {
		id := c.Param("id")
		itemId, err := validateItemId(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := service.DeleteItem(itemId); err != nil {
			handleError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	// update item with specific id
	router.PATCH("/api/item/:id", func(c *gin.Context) {
		var item models.Item
		err := c.BindJSON(&item)
		if err != nil {
			handleError(c, err)
			return
		}

		err = validateItem(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if item.Id != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "The item id must not be provided",
			})
			return
		}

		id := c.Param("id")
		itemId, err := validateItemId(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := service.UpdateItem(itemId, item); err != nil {
			handleError(c, err)
			return
		}

		item.Id = itemId

		c.JSON(http.StatusOK, gin.H{
			"item": item,
		})
	})

	// retrieves the list of items as a CSV file
	router.GET("/api/item/csv", func(c *gin.Context) {
		items, err := service.GetItems()
		if err != nil {
			handleError(c, err)
			return
		}

		csv := "id,name,quantity,unit_price\n"
		for _, item := range items {
			csv += fmt.Sprintf("%d,%s,%d,%f\n", item.Id, item.Name, item.Quantity, item.UnitPrice)
		}

		c.Data(http.StatusOK, "text/csv", []byte(csv))
	})

	return router
}

// handles the server errors
func handleError(c *gin.Context, err error) {
	fmt.Println("Error:", err)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Item not found",
		})
	} else if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Item already exists",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
	}
}

// validate the item id
func validateItemId(id string) (int, error) {
	if id == "" {
		return 0, errors.New("The item id is required")
	}
	itemId, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("The item id must be a number")
	}
	return itemId, nil
}

// validates the item
func validateItem(item models.Item) error {
	if item.Name == "" {
		return errors.New("The item name is required")
	}

	if item.Quantity <= 0 {
		return errors.New("The item quantity must be greater than 0")
	}

	if item.UnitPrice <= 0 {
		return errors.New("The item unit_price must be greater than 0")
	}

	return nil
}
