package main

import (
	"librarizz/controller"
	"librarizz/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db, err := db.InitDB()

	if err != nil {
		panic(err)
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Selamat datang di Librarizz",
		})
	})

	router.POST("/books", controller.InputBook(db))
	router.GET("/books", controller.ShowAllBooks(db))
	router.GET("/books/:id", controller.GetBook(db))
  router.GET("/books/genre=:genre", controller.GetBookByGenre(db))
  router.PUT("/books/:id", controller.UpdateBook(db))
  router.DELETE("/books/:id", controller.DeleteBook(db))
  router.POST("/register", controller.Register(db))
  router.POST("/login", controller.Login(db))

  
  router.Run(":8080")
}
