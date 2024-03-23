package schemas

type CreateBookingRequest struct {
	UserID     uint64 `json:"user_id"`
	SeatID     uint64 `json:"seat_id"`
	FilmID     uint64 `json:"film_id"`
	LocationID uint64 `json:"location_id"`
}
