package is_booked

import (
	"gorm.io/gorm"
	"main.go/models"
)

type IsBookedRepository interface {
	GetBookedSeatBySeatIDAndFilm(SeatID, FilmID uint64) (*models.BookedSeat, error)
}

type isBookedRepository struct {
	db *gorm.DB
}

func NewIsBookedRepository(db *gorm.DB) *isBookedRepository {
	return &isBookedRepository{db}
}
