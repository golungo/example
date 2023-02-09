package main

import (
	"golungo/example/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golungo/lungo"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("scripts/default.config.yaml")
	viper.ReadInConfig()

	uri := viper.Get("database.uri").(string)
	if err := lungo.Init(uri); err != nil {
		panic(err)
	}
	if err := lungo.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := lungo.Disconnect(); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	router.GET("/api/v1/categories/:categoryId", handlers.GetCategoryById)
	router.GET("/api/v1/items/:itemId", handlers.GetItemById)
	router.Run(":8080")
}
