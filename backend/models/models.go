package models

import "time"

type URL struct {
	ID        uint      `gorm:"primaryKey:autoIncrement"`
	Hash      string    `gorm:"type:string;unique:not null"`
	ShortHash string    `gorm:"type:string;unique;not null:index"`
	URL       string    `gorm:"type:string;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
