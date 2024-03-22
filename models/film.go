package models

import "time"

type Film struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	LocationID uint64    `gorm:"not null" json:"location_id"`
	Location   Location  `gorm:"foreignkey:LocationID" json:"location"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"updated_at"`
}
