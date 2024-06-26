package models

import (
	"time"
)

type BaseModel struct {
	ID        int        `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}
