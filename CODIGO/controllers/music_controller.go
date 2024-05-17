package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "codigo/models"
    "codigo/service"
)

type MusicController interface {
    GetAllMusics(ctx *gin.Context)
    GetMusicByID(ctx *gin.Context)
    CreateMusic(ctx *gin.Context)
    UpdateMusic(ctx *gin.Context)
    DeleteMusic(ctx *gin.Context)
    SearchMusic(ctx *gin.Context)
}

type musicController struct {
    service service.MusicService
}

func NewMusicController(service service.MusicService) MusicController {
    return &musicController{
        service: service,
    }
}

func (c *musicController) GetAllMusics(ctx *gin.Context) {
    musics := c.service.FindAll()
    ctx.JSON(http.StatusOK, musics)
}

func (c *musicController) GetMusicByID(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    music, err := c.service.FindByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
        return
    }
    ctx.JSON(http.StatusOK, music)
}

func (c *musicController) CreateMusic(ctx *gin.Context) {
    var newMusic models.Music
    if err := ctx.ShouldBindJSON(&newMusic); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.service.Save(newMusic)
    ctx.JSON(http.StatusCreated, newMusic)
}

func (c *musicController) UpdateMusic(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var updatedMusic models.Music
    if err := ctx.ShouldBindJSON(&updatedMusic); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err = c.service.Update(id, updatedMusic)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
        return
    }
    ctx.JSON(http.StatusOK, updatedMusic)
}

func (c *musicController) DeleteMusic(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    err = c.service.Delete(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Music not found"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Music deleted"})
}

func (c *musicController) SearchMusic(ctx *gin.Context) {
    title := ctx.Query("title")
    artist := ctx.Query("artist")
    musics := c.service.Search(title, artist)
    ctx.JSON(http.StatusOK, musics)
}
