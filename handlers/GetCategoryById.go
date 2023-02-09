package handlers

import (
	"golungo/example/controllers"

	"github.com/gin-gonic/gin"
	"github.com/golungo/lungo"
)

func GetCategoryById(c *gin.Context) {
	categoryId, err := lungo.ObjectIDFromHex(c.Param("categoryId"))
	if err != nil {
		c.JSON(404, "Not Found")
		c.Abort()
		return
	}

	Category, err := controllers.GetCategoryById(categoryId)
	if err != nil {
		c.JSON(404, "Not Found")
		c.Abort()
		return
	}

	c.JSON(200, Category)
	c.Abort()
}
