package bookingservice

import (
	"main.go/models"
	"main.go/schemas"
)

func (s *bookingService) CreateBooking(request schemas.CreateBookingRequest) (*models.Booking, error) {

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
