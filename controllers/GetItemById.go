package controllers

import (
	"encoding/json"

	"golungo/example/models"

	"github.com/golungo/lungo"
)

func GetItemById(itemId lungo.ObjectID) (models.Item, error) {
	var Items []models.Item

	filter := lungo.Filter{
		"_id": itemId,
	}

	populate := lungo.Fields{
		"category",
	}

	result, err := models.GetItemCollection().Match(filter).Lookup(populate).Exec()
	if err != nil {
		return models.Item{}, err
	}

	if err := json.Unmarshal(result, &Items); err != nil {
		return models.Item{}, err
	}

	if len(Items) == 0 {
		return models.Item{}, err
	}

	return Items[0], nil
}
