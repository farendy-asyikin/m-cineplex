package models

import "time"

type Booking struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64    `gorm:"not null" json:"user_id"`
	User        User      `gorm:"foreignkey:UserID" json:"user"`
	SeatID      uint64    `gorm:"not null" json:"seat_id"`
	Seat        Seat      `gorm:"foreignkey:LocationID" json:"seat"`
	FilmID      uint64    `gorm:"not null" json:"film_id"`
	Film        Film      `gorm:"foreignkey:FilmID" json:"film"`
	LocationID  uint64    `gorm:"not null" json:"location_id"`
	Location    Location  `gorm:"foreignkey:LocationID" json:"location"`
	BookingTime time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"booking_time"`
}
