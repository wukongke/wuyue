package models

import (
	"gopkg.in/mgo.v2/bson"

	"work-codes/wuyue/app/common"
)

var TypeVO = common.DB("wuyue").C("type")

type Type struct {
	Id        bson.ObjectId `bson:"_id"`
	TypeNo    int64         `bson:"typeNo"`
	Name      string        `bson:"name"`
	Status    int64         `bson:"status"`
	CreatedAt int64         `bson:"createdAt"`
	UpdatedAt int64         `bson:"updatedAt"`
}
