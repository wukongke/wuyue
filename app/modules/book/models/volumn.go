package models

import (
	"gopkg.in/mgo.v2/bson"

	"work-codes/wuyue/app/common"
)

var VolumnVO = common.DB("wuyue").C("volumn")

type Volumn struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	BookId    string        `bson:bookId`
	WorkCount float64       `bson:workCount` //万字
	Status    int64         `bson:"status"`
	CreatedAt int64         `bson:"createdAt"`
	UpdatedAt int64         `bson:"updatedAt"`
}
