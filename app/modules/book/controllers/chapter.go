package controllers

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"work-codes/wuyue/app/common"
	"work-codes/wuyue/app/modules/book/proxy"
)

var ChapterController *chapterController

type chapterController struct {
	echo.Context
}

func (c *chapterController) List(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	bookNo, err := strconv.Atoi(ctx.QueryParam("bookNo"))
	if err != nil {
		bookNo = 0
	}
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		limit = 9999
	}
	filter := bson.M{}
	if search != "" {
		filter["name"] = search
	}
	if bookNo != 0 {
		filter["bookNo"] = strconv.Itoa(bookNo)
	}
	fmt.Println("filter: ", filter)
	sort := []string{"-updatedAt"}
	result := proxy.ChapterProxy.List(filter, page, limit, sort)

	chapters := []map[string]interface{}{}
	for _, chapter := range (*result)["docs"].([]map[string]interface{}) {
		chapters = append(chapters, *proxy.Convert2Chapter(chapter))
	}
	(*result)["docs"] = chapters
	res := map[string]interface{}{
		"code": 0,
		"msg":  "操作成功",
		"data": *result,
	}
	return ctx.JSON(200, res)
}

func (c *chapterController) Detail(ctx echo.Context) error {
	chapterNo := ctx.Param("chapterNo")
	if chapterNo == "" {
		res := common.ResponseJSON(-1, "chapterNo不能为空", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	row, err := proxy.ChapterProxy.FindOne(bson.M{"chapterNo": chapterNo})
	if err != nil {
		res := common.ResponseJSON(-1, "章节不存在", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	row = proxy.Convert2Chapter(*row)
	res := common.ResponseJSON(0, "操作成功", *row)
	return ctx.JSON(200, res)
}
