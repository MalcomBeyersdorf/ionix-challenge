package repositories

import (
	"regexp"
	"testing"
	"time"
	"vaccine-api/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDrug(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewDrugRepository(db)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO drugs (name, approved, min_dose, max_dose, available_at) VALUES ($1, $2, $3, $4, $5)")).
		WithArgs("Ibuprofen", true, 200, 400, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateDrug(models.Drug{Name: "Ibuprofen", Approved: true, MinDose: 200, MaxDose: 400, AvailableAt: time.Now()})
	require.NoError(t, err)
}

func TestGetDrugByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewDrugRepository(db)
	rows := sqlmock.NewRows([]string{"id", "name", "available_at", "min_dose", "max_dose", "approved"}).
		AddRow(1, "Aspirin", time.Now(), 100, 200, true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, available_at, min_dose, max_dose, approved FROM drugs WHERE id = $1")).
		WithArgs(1).
		WillReturnRows(rows)

	drug, err := repo.GetDrugByID(1)
	require.NoError(t, err)
	assert.Equal(t, "Aspirin", drug.Name)
}

func TestUpdateDrug(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewDrugRepository(db)
	mock.ExpectExec(regexp.QuoteMeta("UPDATE drugs SET name = $1, approved = $2, min_dose = $3, max_dose = $4, available_at = $5 WHERE id = $6")).
		WithArgs("Ibuprofen", true, 200, 400, sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateDrug(1, models.Drug{Name: "Ibuprofen", Approved: true, MinDose: 200, MaxDose: 400, AvailableAt: time.Now()})
	require.NoError(t, err)
}

func TestGetAllDrugs(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewDrugRepository(db)
	rows := sqlmock.NewRows([]string{"id", "name", "available_at", "min_dose", "max_dose", "approved"}).
		AddRow(1, "Aspirin", time.Now(), 100, 200, true).
		AddRow(2, "Ibuprofen", time.Now(), 200, 400, true)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, available_at, min_dose, max_dose, approved FROM drugs")).
		WillReturnRows(rows)

	drugs, err := repo.GetAllDrugs()
	require.NoError(t, err)
	assert.Len(t, drugs, 2)
}

func TestDeleteDrug(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewDrugRepository(db)
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM drugs WHERE id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteDrug(1)
	require.NoError(t, err)
}
