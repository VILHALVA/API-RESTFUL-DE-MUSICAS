package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "codigo/service"
)

type AuthController interface {
    Login(ctx *gin.Context)
}

type authController struct {
    authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
    return &authController{
        authService: authService,
    }
}

func (c *authController) Login(ctx *gin.Context) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := ctx.ShouldBindJSON(&credentials); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    token, err := c.authService.Login(credentials.Username, credentials.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"token": token})
}
