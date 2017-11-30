package controllers

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"work-codes/wuyue/app/common"
	accountProxy "work-codes/wuyue/app/modules/account/proxy"
	"work-codes/wuyue/app/modules/book/proxy"
)

var BookController *bookController

type bookController struct {
	echo.Context
}

func (c *bookController) List(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		limit = 15
	}
	filter := bson.M{}
	if search != "" {
		filter["name"] = search
	}
	sort := []string{"-updatedAt"}
	result, _ := proxy.BookProxy.List(filter, page, limit, sort)

	books := []map[string]interface{}{}
	for _, book := range (*result)["docs"].([]map[string]interface{}) {
		books = append(books, *proxy.Convert2Book(book))
	}
	(*result)["docs"] = books
	res := common.ResponseJSON(0, "操作成功", *result)
	return ctx.JSON(200, res)
}

func (c *bookController) Detail(ctx echo.Context) error {
	bookNo := ctx.Param("bookNo")
	if bookNo == "" {
		res := common.ResponseJSON(-1, "bookNo不能为空", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	row, err := proxy.BookProxy.FindOne(bson.M{"bookNo": bookNo})
	if err != nil {
		fmt.Println("book-detail-err: ", err)
		res := common.ResponseJSON(-1, "书籍不存在", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	row = proxy.Convert2Book(*row)
	res := common.ResponseJSON(0, "操作成功", *row)
	return ctx.JSON(200, res)
}

func (c *bookController) Info(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id == "" {
		res := common.ResponseJSON(-1, "id不能为空", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	row, err := proxy.BookProxy.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		fmt.Println("book-info-err: ", err)
		res := common.ResponseJSON(-1, "书籍不存在", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	book := *(proxy.Convert2Book(*row))
	account, err := accountProxy.AccountProxy.FindOne(bson.M{"_id": book["accountId"]})
	book["author"] = ""
	if err == nil {
		book["author"] = (*account)["name"]
	}
	book["intro"] = string([]rune(book["intro"].(string))[:35]) + "..."
	if book["serialize"] == 0 {
		book["serialize"] = "连载"
	} else {
		book["serialize"] = "完成"
	}
	fmt.Println("likes: ", book["likes"])
	if len(book["likes"].([]interface{})) > 0 {
		likes := []interface{}{}
		for _, v := range book["likes"].([]interface{}) {
			tmp, err := proxy.BookProxy.FindOne(bson.M{"name": v.(string)})
			fmt.Println("err: ", err)
			if err == nil {
				likes = append(likes, (*tmp)["_id"])
			}
		}
		book["likes"] = likes
	}
	res := common.ResponseJSON(0, "操作成功", book)
	return ctx.JSON(200, res)
}
