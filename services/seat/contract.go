package seatservice

import (
	seatrepository "main.go/repositories/seat"
	"main.go/schemas"
)

type SeatService interface {
	ListSeat() ([]*schemas.ListSeatResponse, error)
}

type seatService struct {
	seatRepository seatrepository.SeatRepository
}

func NewSeatService(
	seatRepository seatrepository.SeatRepository,
) *seatService {
	return &seatService{
		seatRepository: seatRepository,
	}
}
