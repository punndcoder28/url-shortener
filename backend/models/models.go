package models

import "time"

type URL struct {
	ID        uint      `gorm:"primaryKey:autoIncrement"`
	Hash      string    `gorm:"type:string;unique:not null"`
	URL       string    `gorm:"type:string;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type CreateURLRequest struct {
	URL string `json:"url"`
}
