package controllers

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"

	"work-codes/wuyue/app/common"
	accountProxy "work-codes/wuyue/app/modules/account/proxy"
	bookProxy "work-codes/wuyue/app/modules/book/proxy"
)

var IndexController *indexController

type indexController struct {
	echo.Context
}

func (c *indexController) BookList(ctx echo.Context) error {
	filter := bson.M{}
	page := 1
	limit := 9999
	sort := []string{"-updatedAt"}
	result, err := bookProxy.BookProxy.List(filter, page, limit, sort)
	data := map[string]interface{}{}
	if err != nil {
		fmt.Println("err: ", err)
		res := common.ResponseJSON(-1, "操作失败", data)
		return ctx.JSON(500, res)
	}

	books := []map[string]interface{}{}
	for _, book := range (*result)["docs"].([]map[string]interface{}) {
		tmp := *bookProxy.Convert2Book(book)
		account, err := accountProxy.AccountProxy.FindOne(bson.M{"_id": tmp["accountId"]})
		tmp["author"] = ""
		if err == nil {
			tmp["author"] = (*account)["name"]
		}
		tmp["intro"] = string([]rune(tmp["intro"].(string))[:35]) + "..."
		if tmp["serialize"] == 0 {
			tmp["serialize"] = "连载"
		} else {
			tmp["serialize"] = "完成"
		}
		books = append(books, tmp)
	}
	(*result)["docs"] = books
	res := common.ResponseJSON(0, "操作成功", *result)
	return ctx.JSON(200, res)
}
