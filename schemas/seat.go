package schemas

type ListSeatResponse struct {
	ID         uint64 `json:"id"`
	Row        string `json:"row"`
	Column     string `json:"column"`
	LocationID uint64 `json:"location_id"`
}
