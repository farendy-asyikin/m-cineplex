package filmrepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type FilmRepository interface {
	CreateFilm(film models.Film) (*models.Film, error)
	GetFilmByID(ID string) (*models.Film, error)
	GetFilmByName(name string) (*models.Film, error)
	UpdateFilm(film models.Film) (*models.Film, error)
	DeleteFilmByID(ID string) error
}

type filmRepository struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) *filmRepository {
	return &filmRepository{db}
}
