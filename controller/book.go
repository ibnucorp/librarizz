package controller

import (
	"librarizz/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InputBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Konversi String ke Integer
		stock, err := strconv.Atoi(ctx.PostForm("stock"))

		if err != nil {
			ctx.JSON(400, "Stock must be numbers")
		}

		// Jika table products belum di database, maka dibuatkan
		db.AutoMigrate(models.Book{})

		// Memasukkan nilai yang diinput user ke variable newBook
		newBook := models.Book{
			Title:     ctx.PostForm("title"),
			Publisher: ctx.PostForm("publisher"),
			Genre:     ctx.PostForm("genre"),
			Stock:     stock,
		}

		db.Create(&newBook)    // Simpan data di table
		ctx.JSON(201, newBook) // Kirim JSON ke user, sukses input
	}
}

func ShowAllBooks(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var books []models.Book

		db.Find(&books)

		ctx.JSON(200, books)
	}

}

func GetBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param("id")
		var book models.Book

		if err := db.First(&book, "id=?", ID).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "Book ID not found"})
			return
		}

		ctx.JSON(200, book)
	}
}

func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param("id")
		var book models.Book

		if err := db.First(&book, "id=?", ID).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Book ID not found"}) // 400 : invalid syntax request to server
		}

		stock, err := strconv.Atoi(ctx.PostForm("stock"))

		if err != nil {
			ctx.JSON(400, "Stock must be numbers")
		}
		var updatedBook = models.Book{
			ID:        book.ID,
			Title:     ctx.PostForm("title"),
			Publisher: ctx.PostForm("publisher"),
			Genre:     ctx.PostForm("genre"),
			Stock:     stock,
		}

		db.Model(&book).Updates(updatedBook)
		ctx.JSON(200, book)
	}
}

func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ID := ctx.Param("id")
		var book models.Book

		if err := db.First(&book, "id=?", ID).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "Book ID not found."})
			return
		}
		db.Delete(&book)
		ctx.JSON(200, gin.H{"message": "Book succesfully deleted."})
	}
}

func GetBookByGenre(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		genre := ctx.Param("genre")
		var books []models.Book

		if err := db.Where("genre LIKE ?", "%"+genre+"%").Find(&books).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "Books not found"})
			return
		}

		if len(books) == 0 {
			ctx.JSON(404, gin.H{"message": "No books found for the given genre"})
			return
		}
		ctx.JSON(200, books)
	}
}

