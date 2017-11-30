package models

import (
	"gopkg.in/mgo.v2/bson"

	"work-codes/wuyue/app/common"
)

var ChapterVO = common.DB("wuyue").C("chapter")

type Chapter struct {
	Id        bson.ObjectId `bson:"_id"`
	ChapterNo int64         `bson:"chapterNo"`
	Title     string        `bson:"title"`
	BookNo    string        `bson:"bookNo"`
	Content   string        `bson:"content"`
	Url       string        `bson:"url"`
	Status    int64         `bson:"status"`
	CreatedAt int64         `bson:"createdAt"`
	UpdatedAt int64         `bson:"updatedAt"`
}
