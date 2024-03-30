package filmrepository

import (
	"errors"
	"gorm.io/gorm"
	"main.go/models"
)

func (r *filmRepository) CreateFilm(film models.Film) (*models.Film, error) {
	err := r.db.Save(&film).Error
	if err != nil {
		return nil, err
	}

	return &film, nil
}

func (r *filmRepository) GetFilmByID(ID string) (*models.Film, error) {
	var film models.Film
	err := r.db.Preload("Location").First(&film, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("film not found")
		}

		return nil, err
	}

	return &film, nil
}

func (r *filmRepository) GetFilmByName(name string) (*models.Film, error) {
	var film models.Film

	err := r.db.Where("name = ?", name).First(&film).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &film, nil
}

func (r *filmRepository) UpdateFilm(film models.Film) (*models.Film, error) {
	err := r.db.Model(&film).Updates(map[string]any{
		"name":        film.Name,
		"location_id": film.LocationID,
	}).Error
	if err != nil {
		return nil, err
	}

	return &film, nil
}

func (r *filmRepository) DeleteFilmByID(ID string) error {
	var film models.Film
	err := r.db.Delete(&film, ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *filmRepository) ListFilm() ([]*models.Film, error) {
	var film []*models.Film

	err := r.db.Preload("Location").Find(&film).Error
	if err != nil {
		return nil, err
	}

	return film, nil
}
