package locationservice

import "main.go/schemas"

func (s *locationService) ListLocation() ([]*schemas.ListLocationResponse, error) {

	locations, err := s.locationRepository.ListLocation()
	if err != nil {
		return nil, err
	}
	var res []*schemas.ListLocationResponse
	for _, location := range locations {
		res = append(res, &schemas.ListLocationResponse{
			ID:   location.ID,
			Name: location.Name,
		})
	}
	return res, nil
}
