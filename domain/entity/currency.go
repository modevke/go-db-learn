package entity

import (
	"time"
	"github.com/gofrs/uuid"
)

type Currency struct {
	ID     uuid.UUID	`gorm:"type:uuid;primary_key;" json:"id"`
	Name   string	`gorm:"size:100;not null;" json:"name"`
	Code   string	`gorm:"size:100;not null;unique" json:"code"`
	Symbol string	`gorm:"size:100;not null;" json:"symbol"`
	Status string	`gorm:"size:100;not null;" json:"status"`

	CreatedAt	time.Time		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt	*time.Time		`json:"deleted_at"`
}
