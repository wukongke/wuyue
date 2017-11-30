package controllers

import "github.com/labstack/echo"

var DemoController *demoController

type demoController struct {
	echo.Context
}

func (c *demoController) Demo(ctx echo.Context) error {
	data := map[string]interface{}{
		"code": 0,
		"msg":  "操作成功",
		"data": map[string]interface{}{},
	}
	return ctx.JSON(200, data)
}
