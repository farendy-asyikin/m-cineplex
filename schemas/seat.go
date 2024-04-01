package schemas

type ListSeatRequest struct {
	FilmID     uint64 `json:"film_id"`
	LocationID uint64 `json:"location_id"`
}

type ListSeatResponse struct {
	ID       uint64 `json:"id"`
	Row      string `json:"row"`
	Column   string `json:"column"`
	Location string `json:"location"`
}
