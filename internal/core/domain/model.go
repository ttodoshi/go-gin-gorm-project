package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	UUID      string         `json:"uuid" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Body      string         `json:"body" gorm:"not null"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime;not null"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}

func (p *Post) BeforeCreate(_ *gorm.DB) (err error) {
	p.UUID = uuid.NewString()
	return
}
