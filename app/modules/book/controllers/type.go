package controllers

import (
	"fmt"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"

	"work-codes/wuyue/app/common"
	accountProxy "work-codes/wuyue/app/modules/account/proxy"
	"work-codes/wuyue/app/modules/book/proxy"
)

var TypeController *typeController

type typeController struct {
	echo.Context
}

var types = map[int]string{
	1: "玄幻",
	2: "修真",
	3: "都市",
	4: "历史",
	5: "网游",
	6: "科幻",
}

func (c *typeController) List(ctx echo.Context) error {
	id, ok := strconv.Atoi(ctx.QueryParam("type"))
	if ok != nil {
		id = 0
	}
	data := map[string]interface{}{
		"types": map[int]string{
			id: types[id],
		},
	}
	res := common.ResponseJSON(0, "操作成功", data)
	return ctx.JSON(200, res)
}

func (c *typeController) Detail(ctx echo.Context) error {
	id, ok := strconv.Atoi(ctx.Param("id"))
	if ok != nil {
		id = 0
	}
	data := types[id]
	res := map[string]interface{}{
		"code": 0,
		"msg":  "操作成功",
		"data": data,
	}
	return ctx.JSON(200, res)
}

func (c *typeController) BookList(ctx echo.Context) error {
	typeNo, ok := strconv.Atoi(ctx.QueryParam("type"))
	if ok != nil {
		typeNo = 1
	}
	filter := bson.M{}
	typeRes, err := proxy.TypeProxy.FindOne(bson.M{"typeNo": typeNo})
	if err == nil {
		filter["type"] = (*typeRes)["name"]
	}
	page := 1
	limit := 9999
	sort := []string{"-updatedAt"}
	result, err := proxy.BookProxy.List(filter, page, limit, sort)
	data := map[string]interface{}{}
	if err != nil {
		fmt.Println("err: ", err)
		res := common.ResponseJSON(-1, "操作失败", data)
		return ctx.JSON(500, res)
	}

	books := []map[string]interface{}{}
	for _, book := range (*result)["docs"].([]map[string]interface{}) {
		tmp := *proxy.Convert2Book(book)
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
