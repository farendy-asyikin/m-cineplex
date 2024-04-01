package schemas

type CreateFilmRequest struct {
	Name       string `json:"name" binding:"required"`
	LocationID uint64 `json:"location_id" binding:"required"`
	IsActive   bool   `json:"is_active"`
}

type UpdateFilmRequest struct {
	Name       *string `json:"name"`
	LocationID *uint64 `json:"location_id"`
	IsActive   *bool   `json:"is_active"`
}

type ListFilmResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type DetailFilmResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	IsActive bool   `json:"is_active"`
}
