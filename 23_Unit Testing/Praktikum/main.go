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

func GetBooks(c echo.Context) error {
	var books []Book
	if err := DB.Find(&books).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, books)
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

func GetBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var books Book
	if err := DB.First(&books, bookId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if books.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, books)
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

func UpdateBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var books Book
	if err := DB.First(&books, bookId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if books.ID == 0 {
		return c.String(http.StatusNotFound, "user notfound")
	}
	if err := c.Bind(&books); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := DB.Save(&books); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, books)
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

func DeleteBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	var books Book
	if err := DB.First(&books, bookId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if books.ID == 0 {
		return c.String(http.StatusNotFound, "user not found")
	}
	if err := DB.Delete(&books, bookId).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, books)
}

func initDb() {
	initDBforUser()

}

func userhandler() {
	e := echo.New()
	e.POST("/users/", CreateUser)
	e.GET("/users/", GetUsers, middleware.JWT([]byte(secretJwt)))
	e.GET("/users/:id", GetUser, middleware.JWT([]byte(secretJwt)))
	e.PUT("/users/:id", UpdateUser, middleware.JWT([]byte(secretJwt)))
	e.DELETE("/users/:id", DeleteUser, middleware.JWT([]byte(secretJwt)))
	e.Start(":8080")
}

func CreateBook(c echo.Context) error {
	book := Book{}
	if err := c.Bind(&book); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	if err := DB.Save(&book).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Internal Sever Error")
	}
	return c.JSON(http.StatusOK, book)
}

func bookhandler() {
	e := echo.New()
	e.POST("/books/", CreateBook, middleware.JWT([]byte(secretJwt)))
	e.GET("/books/", GetBooks, middleware.JWT([]byte(secretJwt)))
	e.GET("/books/:id", GetBook)
	e.PUT("/books/:id", UpdateBook)
	e.DELETE("/books/:id", DeleteBook)
	// e.Start(":8080")
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
