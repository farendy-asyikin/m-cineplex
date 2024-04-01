package models

type Seat struct {
	ID         uint64   `gorm:"primary_key;auto_increment" json:"id"`
	Row        string   `gorm:"type:varchar(255);not null" json:"row"`
	Column     string   `gorm:"type:varchar(255);not null" json:"column"`
	LocationID uint64   `gorm:"not null" json:"location_id"`
	Location   Location `gorm:"foreignkey:LocationID" json:"location"`
}
