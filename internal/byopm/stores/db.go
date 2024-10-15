package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notes struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Title   string    `gorm:"not null;unique"`
	Content string    `gorm:"not null"`
}
