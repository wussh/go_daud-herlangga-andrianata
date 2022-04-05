package main

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const secretJwt = "kentang"

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

func initDBforUser() {
	dsn := "root:minato123@tcp(localhost:3306)/userlogin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&User{})
}

func initDBforBook() {
	dsn := "root:minato123@tcp(localhost:3306)/bookdata?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&Book{})
}

func GetUsers(c echo.Context) error {
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
	token, err := CreateJwtToken(user.Username)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot create token")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"success": "ok",
		"token":   token,
	})
}

func GetUser(c echo.Context) error {
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

func initDb() {
	initDBforUser()

}

func userhandler() {
	e := echo.New()
	e.POST("/users/", CreateUser)
	e.GET("/users/", GetUsers, middleware.JWT([]byte(secretJwt)))
	e.GET("/users/:id", GetUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
	e.Start(":8080")
}

func bookhandler() {
	// e:=echo.New()
	// e.POST("/books/", CreateBook)
	// e.GET("/books/", GetBooks)
	// e.GET("/books/:id", GetBook)
	// e.PUT("/books/:id", UpdateBook)
	// e.DELETE("/books/:id", DeleteBook)
	// e.Start(":8080")
}

func SecretHandler(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	// select * from order where id=username
	return c.String(http.StatusOK, "Hello, "+username)
}

func CreateJwtToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretJwt))
}

func main() {
	initDb()
	userhandler()
}
