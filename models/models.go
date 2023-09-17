// Package models : This package contains all the models for the application
package models

import "time"

// URL : This struct represents the URL model
type URL struct {
	ID        uint      `gorm:"primaryKey:autoIncrement"`
	Hash      string    `gorm:"type:string;unique:not null;default:null"`
	URL       string    `gorm:"type:string;not null;default:null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// CreateURLRequest : This struct represents the CreateURLRequest model
type CreateURLRequest struct {
	URL string `json:"url" validate:"required,url"`
}
