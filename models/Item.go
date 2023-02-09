package models

import (
	"reflect"

	"github.com/golungo/lungo"
	"github.com/golungo/lungo/query"
)

type Item struct {
	query.Model `json:"-" bson:"-"`
	ID          lungo.ObjectID `json:"_id" bson:"_id"`
	Title       string         `json:"title" bson:"title"`
	Content     string         `json:"content" bson:"content"`
	CategoryID  lungo.ObjectID `json:"categoryId" bson:"categoryId"`
	Category    []Category     `json:"category" bson:"-"`
}

func GetItemCollection() query.Model {
	var model Item

	return model.Init(
		"items", reflect.TypeOf(model),
	).Virtual(
		"categories", "categoryId", "_id", "category", false,
	)
}
