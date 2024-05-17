package middlewares

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

func AuthorizeJWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        const BEARER_SCHEMA = "Bearer "
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, BEARER_SCHEMA)
        token, err := validateToken(tokenString)
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            return
        }

        c.Set("userID", claims["user_id"])
        c.Next()
    }
}

func validateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte("your_secret_key"), nil
    })
}
