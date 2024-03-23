package seatservice

import "main.go/schemas"

func (s *seatService) ListSeat() ([]*schemas.ListSeatResponse, error) {
	seats, err := s.seatRepository.ListNotBookedSeat()
	if err != nil {
		return nil, err
	}

	var res []*schemas.ListSeatResponse
	for _, seat := range seats {
		res = append(res, &schemas.ListSeatResponse{
			ID:         seat.ID,
			Row:        seat.Row,
			Column:     seat.Column,
			LocationID: seat.LocationID,
		})
	}
	return res, nil
}
