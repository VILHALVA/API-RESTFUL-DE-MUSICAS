package main

import (
    "github.com/gin-gonic/gin"
    "codigo/controllers"
    "codigo/middlewares"
    "codigo/service"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "codigo/docs"
)

var (
    musicService service.MusicService = service.NewMusicService()
    authService  service.AuthService  = service.NewAuthService()
    musicController controllers.MusicController = controllers.NewMusicController(musicService)
    authController controllers.AuthController = controllers.NewAuthController(authService)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
    r := gin.Default()

    // Swagger setup
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Rotas de autenticação
    r.POST("/auth/login", authController.Login)

    // Rotas CRUD de músicas protegidas por JWT
    musicRoutes := r.Group("/api/music", middlewares.AuthorizeJWT())
    {
        musicRoutes.GET("/", musicController.GetAllMusics)
        musicRoutes.GET("/:id", musicController.GetMusicByID)
        musicRoutes.POST("/", musicController.CreateMusic)
        musicRoutes.PUT("/:id", musicController.UpdateMusic)
        musicRoutes.DELETE("/:id", musicController.DeleteMusic)
        musicRoutes.GET("/search", musicController.SearchMusic)
    }

    // Inicia o servidor na porta 8080
    r.Run(":8080")
}
