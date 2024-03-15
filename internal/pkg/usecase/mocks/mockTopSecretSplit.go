package usecase

import (
	"github.com/OperationQuasarFire/internal/pkg/entity"
	"github.com/OperationQuasarFire/internal/pkg/response"

	"github.com/stretchr/testify/mock"
)

type TopsecretSplitUseCaseMock struct {
	mock.Mock
}

func (m *TopsecretSplitUseCaseMock) CreateTopSecretSplit(satelliteName *string, satelliteData *entity.SatelliteData) (*response.Response, error) {
	args := m.Called(satelliteName, satelliteData)
	return args.Get(0).(*response.Response), args.Error(1)
}

func (m *TopsecretSplitUseCaseMock) GetAll(satelliteName string, satellite *entity.Satellitequery) (*response.Response, error) {
	args := m.Called(satelliteName, satellite)
	return args.Get(0).(*response.Response), args.Error(1)
}

func (m *TopsecretSplitUseCaseMock) PatchtopsecretSplit(satelliteName string, satelliteData *entity.SatelliteData) (*response.Response, error) {
	args := m.Called(satelliteName, satelliteData)
	return args.Get(0).(*response.Response), args.Error(1)
}
