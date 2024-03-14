package ports

import (
	"github.com/OperationQuasarFire/internal/pkg/entity"
)

type CommunicationServices interface {
	GetLocation(...float32) (x, y float32)
	GetMessage(...[]string) (msg string)
}

type DBRepository interface {
	CreateTaskTotal(request *entity.SatelliteRequest) (string, error)
}
