package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func initDB() {
	dsn := "root:minato123@tcp(localhost:3306)/testingorm?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&User{})
}

func GetUser(c echo.Context) error {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := DB.Save(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Internal Sever Error")
	}
	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var users User
	if err := DB.First(&users, userId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if users.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var users User
	if err := DB.First(&users, userId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if users.ID == 0 {
		return c.String(http.StatusNotFound, "user notfound")
	}
	if err := c.Bind(&users); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := DB.Save(&users); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, users)
}

func DeleteUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	var users User
	if err := DB.First(&users, userId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if users.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := DB.Delete(&users, userId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, users)
}

func main() {
	initDB()
	e := echo.New()
	e.POST("/users/", CreateUser)
	e.GET("/users/", GetUser)
	e.GET("/users/:id", GetUsers)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
	e.Start(":8080")
}
