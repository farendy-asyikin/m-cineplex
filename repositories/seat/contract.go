package seatrepository

import (
	"gorm.io/gorm"
	"main.go/models"
)

type SeatRepository interface {
	ListNotBookedSeat() ([]*models.Seat, error)
}

type seatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) *seatRepository {
	return &seatRepository{db}
}
