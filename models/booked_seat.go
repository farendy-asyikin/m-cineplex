package models

type BookedSeat struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	SeatID     uint64 `gorm:"not null" json:"seat_id"`
	FilmID     uint64 `gorm:"not null" json:"film_id"`
	LocationID uint64 `gorm:"not null" json:"location_id"`
}
