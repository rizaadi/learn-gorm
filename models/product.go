package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)"`
	Brand     string `gorm:"not null:type:varchar(191)"`
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Hooks.
// tereksekusi sebelum melakukan create data
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}
	return
}
