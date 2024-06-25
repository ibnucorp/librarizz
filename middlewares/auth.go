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

        // if token invalid
        if err != nil || !token.Valid {
            ctx.JSON(401, gin.H{"message": "Unauthorized"})
            ctx.Abort()
            return
        }

        // if token valid
        claims := token.Claims.(jwt.MapClaims)
        userID := uint(claims["user_id"].(float64))
        ctx.Set("user_id", userID)

        ctx.Next()
    }
}
