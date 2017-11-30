package main

import (
	"github.com/labstack/echo"

	"work-codes/wuyue/app/common"
	"work-codes/wuyue/app/routes"
)

func main() {
	defer common.MgoClose()

	app := echo.New()
	routes.Router(app)

	app.Static("/wuyue/public", "public")

	app.File("/*", "public/index.html")

	app.Debug = true

	app.Logger.Fatal(app.Start(":1323"))
}
