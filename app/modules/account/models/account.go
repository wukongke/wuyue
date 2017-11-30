package models

import (
	"work-codes/wuyue/app/common"

	"gopkg.in/mgo.v2/bson"
)

var AccountVO = common.DB("wuyue").C("account")

type Account struct {
	Id        bson.ObjectId `bson:"_id"`
	Account   string        `bson:"account"`
	UserPwd   string        `bson:"userPwd"`
	Name      string        `bson:"name"`     //昵称
	TrueName  string        `bson:"trueName"` //真实姓名
	Avatar    string        `bson:"avatar"`
	Intro     string        `bson:"Intro"`
	IsAuther  bool          `bson:"isAuther"` //是否作者
	Status    int64         `bson:"status"`
	CreatedAt int64         `bson:"createAt"`
	UpdatedAt int64         `bson:"updateAt"`
}
