package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	kentang "belajar-go-echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.GET("/users", controller.GetAllUsers(db), middleware.JWT([]byte(kentang.SecretJwt)))
	app.POST("/users", controller.CreateUser(db))
	app.Start(":8080")
}
