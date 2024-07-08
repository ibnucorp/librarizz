package controller

import (
  "librarizz/models"
  "strconv"

  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
)

func InputBook(db *gorm.DB) gin.HandlerFunc {
  return func(ctx *gin.Context) {
    // Parse and convert stock
    stock, err := strconv.Atoi(ctx.PostForm("stock"))
    if err != nil {
      ctx.JSON(400, gin.H{"message": "Stock must be a number"})
      return
    }

    // Create a new book instance
    var newBook models.Book
    newBook.Title = ctx.PostForm("title") // Add this line
    newBook.Publisher = ctx.PostForm("publisher") // Add this line
    newBook.Genre = ctx.PostForm("genre") // Add this line
    newBook.Stock = stock

    // Save data and send response
    db.Create(&newBook)
    ctx.JSON(201, newBook)
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
      ctx.JSON(400, gin.H{"message": "Book ID not found"})
      return
    }

    stock, err := strconv.Atoi(ctx.PostForm("stock"))
    if err != nil {
      ctx.JSON(400, gin.H{"message": "Stock must be a number"})
      return
    }

    var updatedBook = models.Book{
      ID:        book.ID,
      Title:    ctx.PostForm("title"),
      Publisher: ctx.PostForm("publisher"),
      Genre:    ctx.PostForm("genre"),
      Stock:    stock,
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
      ctx.JSON(404, gin.H{"message": "Book ID not found"})
      return
    }

    db.Delete(&book)
    ctx.JSON(200, gin.H{"message": "Book successfully deleted"})
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
