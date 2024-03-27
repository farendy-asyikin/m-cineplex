package bookingservice

import (
	"errors"
	"fmt"
	"main.go/models"
	"main.go/schemas"
)

func (s *bookingService) CreateBooking(request schemas.CreateBookingRequest) (*models.Booking, error) {

	isBooked, _ := s.isBookedRepository.GetBookedSeatBySeatIDAndFilm(request.FilmID, request.SeatID)
	fmt.Println(isBooked)
	if isBooked.ID != 0 {
		return nil, errors.New("Seat Already Booked")
	}

	bkg := models.Booking{
		UserID:     request.UserID,
		SeatID:     request.SeatID,
		FilmID:     request.FilmID,
		LocationID: request.LocationID,
	}
	booking, err := s.bookingRepository.CreateBooking(bkg)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (s *bookingService) GetBookingByID(ID string) (*models.Booking, error) {
	booking, err := s.bookingRepository.GetBookingByID(ID)
	if err != nil {
		return nil, err
	}

	return booking, nil
}
