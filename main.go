package main

import (
	"go-web-api/book"
	"go-web-api/handler"
	"go-web-api/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root@tcp(127.0.0.1:3306)/db_pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db Connection Failed")
	}

	db.AutoMigrate(&user.User{})

	routes := gin.Default()

	//books
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewService(bookService)
	routes.POST("/book", bookHandler.PostBook)
	routes.GET("/books", bookHandler.GetAllBook)
	routes.GET("/book/:id", bookHandler.GetBook)
	routes.PUT("/book/:id", bookHandler.UpdateBook)
	routes.DELETE("/book/:id", bookHandler.DeleteBook)

	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewHandler(userService)
	routes.GET("/users", userHandler.GetUsers)
	routes.GET("/users/:id", userHandler.GetUser)
	routes.POST("/user", userHandler.CreateUser)
	routes.DELETE("/users/:id", userHandler.Deleteuser)

	routes.Run()
}
