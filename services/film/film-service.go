package filmservice

import (
	"main.go/models"
	"main.go/schemas"
)

func (s *filmService) CreateFilm(request schemas.CreateFilmRequest) (*models.Film, error) {

	user := models.Film{
		Name: request.Name,

		Password: string(passwordHash),
	}

	result, err := s.filmRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *filmService) UpdateFilm(request schemas.UpdateFilmRequest, film models.Film) (*models.Film, error) {

	result, err := s.filmRepository.UpdateFilm(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *filmService) DeleteFilmByID(ID string) error {
	err := s.filmRepository.DeleteFilmByID(ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *filmService) GetFilmByID(ID string) (*models.Film, *schemas.DetailFilmResponse, error) {

}
