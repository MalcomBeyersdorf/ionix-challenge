package models

import (
	"time"
)

type Drug struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Approved    bool      `json:"approved"`
	MinDose     int       `json:"min_dose"`
	MaxDose     int       `json:"max_dose"`
	AvailableAt time.Time `json:"available_at"`
}
