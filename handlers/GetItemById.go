package handlers

import (
	"golungo/example/controllers"

	"github.com/gin-gonic/gin"
	"github.com/golungo/lungo"
)

func GetItemById(c *gin.Context) {
	itemId, err := lungo.ObjectIDFromHex(c.Param("itemId"))
	if err != nil {
		c.JSON(404, "Not Found")
		c.Abort()
		return
	}

	Item, err := controllers.GetItemById(itemId)
	if err != nil {
		c.JSON(404, "Not Found")
		c.Abort()
		return
	}

	c.JSON(200, Item)
	c.Abort()
}
