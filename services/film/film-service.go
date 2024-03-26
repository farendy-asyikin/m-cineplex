package filmservice

import (
	"errors"
	"main.go/models"
	"main.go/schemas"
)

func (s *filmService) CreateFilm(request schemas.CreateFilmRequest) (*models.Film, error) {
	isExist, err := s.filmRepository.GetFilmByName(request.Name)
	if err != nil {
		return nil, err
	}

	if isExist != nil {
		return nil, errors.New("film already exist")
	}

	film := models.Film{
		Name:       request.Name,
		LocationID: request.LocationID,
	}

	res, err := s.filmRepository.CreateFilm(film)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *filmService) UpdateFilm(request schemas.UpdateFilmRequest, film models.Film) (*models.Film, error) {
	if request.Name != nil {
		isExist, err := s.filmRepository.GetFilmByName(*request.Name)
		if err != nil {
			return nil, err
		}
		if isExist != nil && isExist.ID != film.ID {
			return nil, errors.New("film already exist")
		}

		film.Name = *request.Name
	}

	if request.LocationID != nil {
		film.LocationID = *request.LocationID
	}

	result, err := s.filmRepository.UpdateFilm(film)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *filmService) ListFilm() ([]*schemas.ListFilmResponse, error) {
	films, err := s.filmRepository.ListFilm()
	if err != nil {
		return nil, err
	}
	var res []*schemas.ListFilmResponse
	for _, film := range films {
		res = append(res, &schemas.ListFilmResponse{
			Name:     film.Name,
			Location: film.Location.Name,
		})
	}
	return res, nil
}

func (s *filmService) DeleteFilmByID(ID string) error {
	err := s.filmRepository.DeleteFilmByID(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *filmService) GetFilmByID(ID string) (*models.Film, error) {
	film, err := s.filmRepository.GetFilmByID(ID)
	if err != nil {
		return nil, err
	}

	return film, nil
}
