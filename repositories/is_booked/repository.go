package is_booked

import (
	"main.go/models"
)

func (r *isBookedRepository) GetBookedSeatBySeatIDAndFilm(SeatID, FilmID uint64) (*models.BookedSeat, error) {
	var isBooked models.BookedSeat
	r.db.First(&isBooked, SeatID, FilmID)

	return &isBooked, nil
}
