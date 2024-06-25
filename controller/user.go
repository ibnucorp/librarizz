package controller

import (
	"librarizz/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db.AutoMigrate(models.User{})

		var existingUser models.User

		if err := db.Where("email = ?", ctx.PostForm("email")).First(&existingUser).Error; err == nil {
			ctx.JSON(400, gin.H{"message": "Email already exists."})
			return
		}

		hash, _ := HashPassword(ctx.PostForm("password"))

		var newUser = models.User{
			Email:    ctx.PostForm("email"),
			Name:     ctx.PostForm("name"),
			Password: hash,
			Status:   ctx.PostForm("status"),
		}

		if err := db.Create(&newUser).Error; err != nil {
			ctx.JSON(500, gin.H{"message": "Internal server error."})
			return
		}

		ctx.JSON(201, gin.H{"message": "User registered successfully."})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User

		if err := db.Where("email = ?", ctx.PostForm("email")).First(&user).Error; err != nil {
			ctx.JSON(401, gin.H{"message": "Email not found."})
			return
		}

		if !CheckPasswordHash(ctx.PostForm("password"), user.Password) {
			ctx.JSON(401, gin.H{"message": "Password invalid."})
			return
		}

		token, err := CreateToken(user.ID)

		if err != nil {
			ctx.JSON(500, gin.H{"message": "Internal server error."})
			return
		}

		ctx.JSON(200, gin.H{
			"name":  user.Name,
			"token": token})
	}
}

