package entity

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid,primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func DefaultBase() Base {
	t := time.Now()
	return Base{
		CreatedAt: t,
		UpdatedAt: t,
		DeletedAt: time.Time{},
	}
}

func UpdateBase() Base {
	return Base{
		UpdatedAt: time.Now(),
	}
}
