package entity

import (
	"time"
	"github.com/gofrs/uuid"
)

type Country struct {
	ID        uuid.UUID	`gorm:"type:uuid;primary_key;" json:"id"`
	Name      string	`gorm:"size:100;not null;" json:"name"`	
	Iso       string	`gorm:"size:100;not null;unique" json:"iso"`
	Iso3      string	`gorm:"size:100;not null;" json:"iso3"`
	Numcode   string	`gorm:"size:100;not null;" json:"numcode"`
	Phonecode string	`gorm:"size:100;not null;" json:"phonecode"`
	Currency  string	`gorm:"size:100;not null;" json:"currency"`
	Status    string	`gorm:"size:100;not null;" json:"status"`

	CreatedAt	time.Time		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt	*time.Time	`json:"deleted_at"`
}
