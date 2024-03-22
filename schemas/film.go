package schemas

type CreateFilmRequest struct {
	Name       string `json:"name" binding:"required"`
	LocationID string `json:"email" binding:"required,email"`
	IsActive   bool   `gorm:"default:true" json:"is_active"`
}

type UpdateFilmRequest struct {
	Name       *string `json:"name"`
	LocationID *string `json:"email"`
	IsActive   *bool   `json:"is_active"`
}

type DetailFilmResponse struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	LocationID string `json:"location_id"`
	IsActive   bool   `json:"is_active"`
}