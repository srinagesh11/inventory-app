package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Unit_Price float64 `json:"price"`
}

var items = []item{
	{ID: "1", Name: "Blue Train", Quantity: 5, Unit_Price: 56.99},
	{ID: "2", Name: "Jeru", Quantity: 10, Unit_Price: 17.99},
	{ID: "3", Name: "Sarah Vaughan and Clifford Brown", Quantity: 15, Unit_Price: 39.99},
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range items {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func postItems(c *gin.Context) {
	var newItem item

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}
func main() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/item", getItems)
		v1.GET("/item/:id", getItemByID)
		v1.POST("/item", postItems)
	}

	router.Run("localhost:8080")
}
