package middlewares

import (
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

func Auth() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    tokenString := ctx.GetHeader("Authorization")

    if tokenString == "" {
      ctx.JSON(401, gin.H{"message": "Unauthorized"})
      ctx.Abort()
      return
    }

    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
      return jwtKey, nil
    })

    // Validasi token
    if err != nil || !token.Valid {
      ctx.JSON(401, gin.H{"message": "Unauthorized"})
      ctx.Abort()
      return
    }

    // Ambil klaim dari token
    claims := token.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))
    userRole := claims["role"].(string) // Asumsikan "role" adalah klaim di JWT Anda

    // Set informasi user ke context
    ctx.Set("user_id", userID)
    ctx.Set("role", userRole)

    // Periksa role user
    if userRole != "admin" && userRole != "user" {
      ctx.JSON(403, gin.H{"message": "Invalid role"})
      ctx.Abort()
      return
    }

    ctx.Next()
  }
}
