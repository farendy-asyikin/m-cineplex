package models

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	IsActive  bool      `gorm:"type:tinyint(1);default:0" json:"is_active"`
	Roles     string    `gorm:"type:varchar(255);not null" json:"roles"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"updated_at"`
}
