package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	FirstScrewUp time.Time `json:"firstScrewUp"`
}
