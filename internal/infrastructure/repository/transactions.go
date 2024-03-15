package repository

import (
	"strings"

	"github.com/OperationQuasarFire/internal/pkg/entity"
)

func (p BDRepository) CreateTaskTotal(request *entity.SatelliteRequest) (string, error) {
	query := "INSERT INTO Satellites (name, distance, message) VALUES ($1, $2, $3)"
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	for _, satellite := range request.Satellites {
		messageString := ""
		if len(satellite.Message) > 0 {
			messageString = strings.Join(satellite.Message, ",")
		}

		_, err = stmt.Exec(satellite.Name, satellite.Distance, messageString)
		if err != nil {
			return "", err
		}
	}

	return "Registros creados en la base de datos", nil
}
