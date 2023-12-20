package repositories

import (
	"database/sql"
	"errors"
	"vaccine-api/models"
)

type VaccinationRepository struct {
	Db *sql.DB
}

func NewVaccinationRepository(db *sql.DB) *VaccinationRepository {
	return &VaccinationRepository{Db: db}
}

func (repo *VaccinationRepository) CreateVaccination(vaccination models.Vaccination, drugRepo *DrugRepository) error {
	drug, err := drugRepo.GetDrugByID(vaccination.DrugID)
	if err != nil {
		return err
	}

	if vaccination.Dose < drug.MinDose || vaccination.Dose > drug.MaxDose {
		return errors.New("the dose is out of the allowed range")
	}

	if vaccination.Date.Before(drug.AvailableAt) {
		return errors.New("the vaccination date is earlier than the drug's available date")
	}

	query := `INSERT INTO vaccinations (name, drug_id, dose, date) VALUES ($1, $2, $3, $4)`
	_, err = repo.Db.Exec(query, vaccination.Name, vaccination.DrugID, vaccination.Dose, vaccination.Date)
	return err
}

func (repo *VaccinationRepository) UpdateVaccination(id int, vaccination models.Vaccination) error {
	query := `UPDATE vaccinations SET name = $1, drug_id = $2, dose = $3, date = $4 WHERE id = $5`
	_, err := repo.Db.Exec(query, vaccination.Name, vaccination.DrugID, vaccination.Dose, vaccination.Date, id)
	return err
}

func (repo *VaccinationRepository) GetAllVaccinations() ([]models.Vaccination, error) {
	var vaccinations []models.Vaccination
	query := `SELECT id, name, drug_id, dose, date FROM vaccinations`

	rows, err := repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var vaccination models.Vaccination
		if err := rows.Scan(&vaccination.ID, &vaccination.Name, &vaccination.DrugID, &vaccination.Dose, &vaccination.Date); err != nil {
			return nil, err
		}
		vaccinations = append(vaccinations, vaccination)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return vaccinations, nil
}

func (repo *VaccinationRepository) DeleteVaccination(id int) error {
	query := `DELETE FROM vaccinations WHERE id = $1`
	_, err := repo.Db.Exec(query, id)
	return err
}
