package main

import (
	//"encoding/json"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_basic_crud/db"
)

func GetAllItems(c *gin.Context) {
	items, err := db.GetAllItems()
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to load database items: %v", err),
		)
		return
	}

	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")

	item, err := db.GetItem(id)
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to read database: %v", err),
		)
		return
	}

	c.JSON(http.StatusOK, item)
}

func PostItem(c *gin.Context) {
	var item db.Item

	err := c.BindJSON(&item)
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to parse input data: %v", err),
		)
		return
	}

	if err = db.SaveItem(item); err != nil {
		c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to save data: %v", err),
		)
		return
	}

	c.String(http.StatusOK, "OK")
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	if err := db.RemoveItem(id); err != nil {
		c.String(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to remove item: %v", err),
		)
		return
	}

	c.String(http.StatusOK, "OK")
}

func main() {
	router := gin.Default()

	router.GET("/api/items", GetAllItems)

	router.GET("/api/items/:id", GetItem)

	router.POST("/api/items", PostItem)

	router.DELETE("/api/items/:id", DeleteItem)

	router.Run(":8080")
}
