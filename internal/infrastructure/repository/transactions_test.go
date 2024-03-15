package repository

import (
	"errors"
	"testing"

	"github.com/OperationQuasarFire/internal/pkg/entity"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestBDRepository_Create_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := BDRepository{db}

	satellites := []entity.Satellite{
		{Name: "satellite1", Distance: 100, Message: []string{"message1", "message2"}},
		{Name: "satellite2", Distance: 200, Message: []string{"message3", "message4"}},
	}

	mock.ExpectPrepare("INSERT INTO Satellites").WillReturnError(errors.New("error preparing statement"))

	request := &entity.SatelliteRequest{Satellites: satellites}

	_, err = repo.Create(request)

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
