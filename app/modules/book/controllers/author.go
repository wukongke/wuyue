package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"work-codes/wuyue/app/modules/book/proxy"
)

type AutherController struct {
	echo.Context
}

func (c *AutherController) BookList(ctx echo.Context) error {
	search := ctx.QueryParam("auther")
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
	res := map[string]interface{}{
		"code": 0,
		"msg":  "操作成功",
		"data": *result,
	}
	return ctx.JSON(200, res)
}
