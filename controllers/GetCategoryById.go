package controllers

import (
	"encoding/json"

	"golungo/example/models"

	"github.com/golungo/lungo"
)

func GetCategoryById(categoryId lungo.ObjectID) (models.Category, error) {
	var Categories []models.Category

	filter := lungo.Filter{
		"_id": categoryId,
	}

	populate := lungo.Fields{
		"items",
	}

	result, err := models.GetCategoryCollection().Match(filter).Lookup(populate).Exec()
	if err != nil {
		return models.Category{}, err
	}

	if err := json.Unmarshal(result, &Categories); err != nil {
		return models.Category{}, err
	}

	if len(Categories) == 0 {
		return models.Category{}, err
	}

	return Categories[0], nil
}
