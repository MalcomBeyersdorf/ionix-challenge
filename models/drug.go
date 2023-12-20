package models

import (
	"time"
)

type Drug struct {
	AvailableAt time.Time `json:"available_at"`
	Name        string    `json:"name"`
	ID          int       `json:"id"`
	MinDose     int       `json:"min_dose"`
	MaxDose     int       `json:"max_dose"`
	Approved    bool      `json:"approved"`
}
