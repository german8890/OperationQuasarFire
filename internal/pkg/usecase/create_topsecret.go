package usecase

import (
	"fmt"

	entity "github.com/OperationQuasarFire/internal/pkg/entity"
	"github.com/OperationQuasarFire/internal/pkg/ports"
	entity2 "github.com/OperationQuasarFire/internal/pkg/response"
)

type TopsecrestUseCase struct {
	topsecrestService ports.CommunicationServices
	repo              ports.DBRepository
}

func NewTopsecretUseCase(service ports.CommunicationServices, repo ports.DBRepository) *TopsecrestUseCase {
	return &TopsecrestUseCase{
		topsecrestService: service,
		repo:              repo,
	}
}

func (uc *TopsecrestUseCase) CreateTopSecret(satellite *entity.SatelliteRequest) (*entity2.Response, error) {

	var Position entity2.Position
	var message [][]string
	distance := []float32{}

	uc.repo.Create(satellite)

	for _, satellite := range satellite.Satellites {
		message = append(message, satellite.Message)
		distance = append(distance, float32(satellite.Distance))
	}
	Position.Message = uc.topsecrestService.GetMessage(message...)

	Position.X, Position.Y = uc.topsecrestService.GetLocation(distance...)
	if Position.X == 0 && Position.Y == 0 || Position.Message == "" {
		return nil, fmt.Errorf("the position or the message cannot be determined")
	} else {
		return &entity2.Response{
			Position: entity2.Position{
				X:       Position.X,
				Y:       Position.Y,
				Message: Position.Message,
			},
			Result: entity2.Result{
				Source:  "Find position to message",
				Details: nil,
			},
		}, nil
	}
}
