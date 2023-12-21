package repositories

import (
	"errors"
	"regexp"
	"testing"
	"time"
	"vaccine-api/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockDrugRepository struct {
	Drugs map[int]models.Drug
}

func (m *MockDrugRepository) GetDrugByID(id int) (models.Drug, error) {
	if drug, exists := m.Drugs[id]; exists {
		return drug, nil
	}
	return models.Drug{}, errors.New("drug not found")
}

// func TestCreateVaccination(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	require.NoError(t, err)
// 	defer db.Close()
//
// 	repo := NewVaccinationRepository(db)
// 	drugRepo := &MockDrugRepository{
// 		Drugs: map[int]models.Drug{
// 			1: {ID: 1, MinDose: 100, MaxDose: 200, AvailableAt: time.Now().AddDate(0, 0, -1)},
// 		},
// 	}
// 	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO vaccinations (name, drug_id, dose, date) VALUES ($1, $2, $3, $4)")).
// 		WithArgs("Vaccination 1", 1, 150, sqlmock.AnyArg()).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
//
// 	err = repo.CreateVaccination(models.Vaccination{Name: "Vaccination 1", DrugID: 1, Dose: 150, Date: time.Now()}, drugRepo)
// 	require.NoError(t, err)
// }

func TestUpdateVaccination(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewVaccinationRepository(db)

	mock.ExpectExec(regexp.QuoteMeta("UPDATE vaccinations SET name = $1, drug_id = $2, dose = $3, date = $4 WHERE id = $5")).
		WithArgs("Vaccination Updated", 1, 150, sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateVaccination(1, models.Vaccination{Name: "Vaccination Updated", DrugID: 1, Dose: 150, Date: time.Now()})
	require.NoError(t, err)
}

func TestGetAllVaccinations(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewVaccinationRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "drug_id", "dose", "date"}).
		AddRow(1, "Vaccination 1", 1, 150, time.Now()).
		AddRow(2, "Vaccination 2", 2, 250, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, drug_id, dose, date FROM vaccinations")).
		WillReturnRows(rows)

	vaccinations, err := repo.GetAllVaccinations()
	require.NoError(t, err)
	assert.Len(t, vaccinations, 2)
}

func TestDeleteVaccination(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewVaccinationRepository(db)

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM vaccinations WHERE id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteVaccination(1)
	require.NoError(t, err)
}
