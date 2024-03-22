package filmservice

import (
	"main.go/models"
	filmrepository "main.go/repositories/film"
	"main.go/schemas"
)

type FilmService interface {
	CreateFilm(request schemas.CreateFilmRequest) (*models.Film, error)
	UpdateFilm(request schemas.UpdateFilmRequest, film models.Film) (*models.Film, error)
	DeleteFilmByID(ID string) error
	GetFilmByID(ID string) (*models.Film, *schemas.DetailFilmResponse, error)
}

type filmService struct {
	filmRepository filmrepository.FilmRepository
}

func NewFilmService(
	filmRepository filmrepository.FilmRepository,
) *filmService {
	return &filmService{
		filmRepository: filmRepository,
	}
}
