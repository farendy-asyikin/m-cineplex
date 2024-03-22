package filmhandler

import (
	"github.com/gin-gonic/gin"
	filmservice "main.go/services/film"
)

type FilmHandler interface {
	CreateFilm(ctx *gin.Context)
	GetFilmById(ctx *gin.Context)
	UpdateFilm(ctx *gin.Context)
	DeleteFilmByID(ctx *gin.Context)
}

type filmHandler struct {
	filmService filmservice.FilmService
}

func NewFilmHandler(filmService filmservice.FilmService) *filmHandler {
	return &filmHandler{filmService}
}
