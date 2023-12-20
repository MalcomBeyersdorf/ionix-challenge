package models

import "time"

type Vaccination struct {
	Date   time.Time `json:"date"`
	Name   string    `json:"name"`
	ID     int       `json:"id"`
	DrugID int       `json:"drug_id"`
	Dose   int       `json:"dose"`
}
