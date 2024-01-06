package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Firstname    string
	Lastname     string
	FirstScrewUp time.Time
}
