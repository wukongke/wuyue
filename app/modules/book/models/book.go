package models

import (
	"gopkg.in/mgo.v2/bson"

	"work-codes/wuyue/app/common"
)

var BookVO = common.DB("wuyue").C("book")

type Book struct {
	Id          bson.ObjectId `bson:"_id"`
	BookNo      int64         `bson:"bookNo"`
	Name        string        `bson:"name"`
	AccountId   string        `bson:"accountId"`
	Type        string        `bson:"type"`
	Intro       string        `bson:"intro"`
	Image       string        `bson:"image"`
	Rate        float64       `bson:"rate"`      //总评分
	WordCount   float64       `bson:"wordcount"` //万字
	Serialize   int64         `bson:"serialize"` //是否完结
	LoveCount   int64         `bson:"loveCount"`
	Likes       []string      `bson:"likes"`
	Url         string        `bson:"url"`
	LastChapter string        `bson:"lastChapter"` //最后更新章节
	Status      int64         `bson:"status"`
	CreatedAt   int64         `bson:"createdAt"`
	UpdatedAt   int64         `bson:"updatedAt"`
}
