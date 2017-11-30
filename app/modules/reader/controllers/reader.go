package controllers

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"

	"work-codes/wuyue/app/common"
	"work-codes/wuyue/app/modules/book/proxy"
)

var ReaderController *readerController

type readerController struct {
	echo.Context
}

func (c *readerController) Title(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id == "" {
		res := common.ResponseJSON(-1, "id不能为空", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	book, err := proxy.BookProxy.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		fmt.Println("book-err: ", err)
		res := common.ResponseJSON(-1, "小说不存在", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	chapters, err := proxy.ChapterProxy.Find(bson.M{"bookNo": (*book)["bookNo"]})
	if err != nil {
		fmt.Println("chapters-err: ", err)
	}
	data := map[string]interface{}{
		"id":   (*book)["_id"],
		"name": (*book)["name"],
	}
	// titles := ""
	// if len(chapters) != 0 {
	// 	for _, v := range chapters {
	// 		titles += (*v)["title"].(string) + "-"
	// 	}
	// }
	titles := []map[string]interface{}{}
	for _, v := range chapters {
		chap := map[string]interface{}{
			"id":    (*v)["_id"],
			"title": (*v)["title"],
		}
		titles = append(titles, chap)
	}
	data["titles"] = titles
	res := common.ResponseJSON(0, "操作成功", data)
	return ctx.JSON(200, res)
}

func (c *readerController) Info(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id == "" {
		res := common.ResponseJSON(-1, "id不能为空", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	bookId := ctx.QueryParam("book")
	if bookId == "" {
		res := common.ResponseJSON(-1, "bookId不能为空", map[string]interface{}{})
		return ctx.JSON(500, res)
	}
	var chapter map[string]interface{}
	if id == "1" {
		bookRow, err := proxy.BookProxy.FindOne(bson.M{"_id": bson.ObjectIdHex(bookId)})
		if err != nil {
			res := common.ResponseJSON(-1, "小说不存在", map[string]interface{}{})
			return ctx.JSON(500, res)
		}
		row, err := proxy.ChapterProxy.FindOne(bson.M{"bookNo": (*bookRow)["bookNo"]})
		if err != nil {
			fmt.Println("book-info-err: ", err)
			res := common.ResponseJSON(-1, "章节不存在", map[string]interface{}{})
			return ctx.JSON(500, res)
		}
		chapter = *(proxy.Convert2Chapter(*row))
	} else {
		row, err := proxy.ChapterProxy.FindOne(bson.M{"_id": bson.ObjectIdHex(id)})
		if err != nil {
			fmt.Println("chap-info-err: ", err)
			res := common.ResponseJSON(-1, "章节不存在", map[string]interface{}{})
			return ctx.JSON(500, res)
		}
		chapter = *(proxy.Convert2Chapter(*row))
	}
	book, err := proxy.BookProxy.FindOne(bson.M{"bookNo": chapter["bookNo"]})
	chapter["bookName"] = ""
	if err == nil {
		chapter["bookName"] = (*book)["name"]
	}
	res := common.ResponseJSON(0, "操作成功", chapter)
	return ctx.JSON(200, res)
}
