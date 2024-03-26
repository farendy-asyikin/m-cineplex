package seatservice

import "main.go/schemas"

func (s *seatService) ListUnbookedSeatByFilmIDAndLocationID(filmID, locationID uint64) ([]*schemas.ListSeatResponse, error) {
	seats, err := s.seatRepository.ListUnbookedSeatByFilmIDAndLocationID(filmID, locationID)
	if err != nil {
		return nil, err
	}

	var res []*schemas.ListSeatResponse
	for _, seat := range seats {
		res = append(res, &schemas.ListSeatResponse{
			ID:     seat.ID,
			Row:    seat.Row,
			Column: seat.Column,
			//Location: seat.LocationID,
		})
	}
	return res, nil
}
