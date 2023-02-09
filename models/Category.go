package models

import (
	"reflect"

	"github.com/golungo/lungo"
	"github.com/golungo/lungo/query"
)

type Category struct {
	query.Model `json:"-" bson:"-"`
	ID          lungo.ObjectID `json:"_id" bson:"_id"`
	Title       string         `json:"title" bson:"title"`
	Items       []Item         `json:"items,omitempty" bson:"-"`
}

func GetCategoryCollection() query.Model {
	var model Category

	return model.Init(
		"categories", reflect.TypeOf(model),
	).Virtual(
		"items", "_id", "categoryId", "items", false,
	)
}
