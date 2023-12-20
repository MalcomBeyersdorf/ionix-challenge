package repositories

import (
	"database/sql"
	"vaccine-api/models"
)

type DrugRepository struct {
	Db *sql.DB
}

func NewDrugRepository(db *sql.DB) *DrugRepository {
	return &DrugRepository{Db: db}
}

func (repo *DrugRepository) GetDrugByID(id int) (models.Drug, error) {
	var drug models.Drug
	query := `SELECT id, name, available_at, min_dose, max_dose, approved FROM drugs WHERE id = $1`
	err := repo.Db.QueryRow(query, id).Scan(&drug.ID, &drug.Name, &drug.AvailableAt, &drug.MinDose, &drug.MaxDose, &drug.Approved)
	if err != nil {
		return models.Drug{}, err
	}
	return drug, nil
}

func (repo *DrugRepository) CreateDrug(drug models.Drug) error {
	query := `INSERT INTO drugs (name, approved, min_dose, max_dose, available_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := repo.Db.Exec(query, drug.Name, drug.Approved, drug.MinDose, drug.MaxDose, drug.AvailableAt)
	return err
}

func (repo *DrugRepository) UpdateDrug(id int, drug models.Drug) error {
	query := `UPDATE drugs SET name = $1, approved = $2, min_dose = $3, max_dose = $4, available_at = $5 WHERE id = $6`
	_, err := repo.Db.Exec(query, drug.Name, drug.Approved, drug.MinDose, drug.MaxDose, drug.AvailableAt, id)
	return err
}

func (repo *DrugRepository) GetAllDrugs() ([]models.Drug, error) {
	var drugs []models.Drug
	query := `SELECT id, name, available_at, min_dose, max_dose, approved FROM drugs`

	rows, err := repo.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var drug models.Drug
		if err := rows.Scan(&drug.ID, &drug.Name, &drug.AvailableAt, &drug.MinDose, &drug.MaxDose, &drug.Approved); err != nil {
			return nil, err
		}
		drugs = append(drugs, drug)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return drugs, nil
}

func (repo *DrugRepository) DeleteDrug(id int) error {
	query := `DELETE FROM drugs WHERE id = $1`
	_, err := repo.Db.Exec(query, id)
	return err
}
