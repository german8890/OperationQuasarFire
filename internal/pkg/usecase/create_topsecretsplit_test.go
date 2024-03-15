package usecase_test

import (
	"testing"

	mock1 "github.com/OperationQuasarFire/internal/infrastructure/api/starshipCommsResolver/mock"
	"github.com/OperationQuasarFire/internal/pkg/entity"
	"github.com/OperationQuasarFire/internal/pkg/response"
	"github.com/OperationQuasarFire/internal/pkg/usecase"
	mock2 "github.com/OperationQuasarFire/internal/pkg/usecase/mocks"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTopsecretSplitUseCase_CreateTopSecretSplit_Success(t *testing.T) {
	mockService := new(mock1.CommunicationServicesMock)
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)

	usecase := usecase.NewTopsecretSplitUseCase(mockService, cache)

	satelliteName := "kenobi"
	satelliteData := &entity.SatelliteData{
		Distance: 100,
		Message:  []string{"mensaje1"},
	}

	expectedResponse := &response.Response{
		Result: response.Result{
			Source:  "Message and distance added successfully",
			Details: nil,
		},
	}

	mockService.On("GetLocation", mock.Anything).Return(float32(300.0), float32(400.0))
	mockService.On("GetMessage", mock.Anything).Return("Mensaje decodificado")

	response, _ := usecase.CreateTopSecretSplit(&satelliteName, satelliteData)

	assert.Equal(t, expectedResponse, response)
}

func TestTopsecretSplitUseCase_GetAll_Success(t *testing.T) {
	// Crear un nuevo mapa de datos para cada test
	var satelliteDataMap = make(map[string]*entity.SatelliteData)
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)

	mockService := new(mock1.CommunicationServicesMock)
	mockTopsecretsplitUseCase := new(mock2.TopsecretSplitUseCaseMock)
	usecase := usecase.NewTopsecretSplitUseCase(mockService, cache)

	satelliteName := "kenobi"

	existingData := &entity.SatelliteData{
		Distance: 100,
		Message:  []string{"original message"},
	}

	// Mock the existing data in the map
	satelliteDataMap[satelliteName] = existingData

	// Crear una instancia de Satellitequery con los valores necesarios
	satelliteQuery := &entity.Satellitequery{
		// Asigna los valores apropiados a los campos según la definición de Satellitequery
	}

	// Crear una instancia de response.Response con los valores necesarios
	expectedResponse := (*response.Response)(nil)

	// Mock the CommunicationServices methods
	mockService.On("GetLocation", mock.Anything).Return(float32(300.0), float32(400.0))
	mockService.On("GetMessage", mock.Anything).Return("Mensaje decodificado")

	// Mock the method call and return the expected response
	mockTopsecretsplitUseCase.On("GetAll", satelliteName, satelliteQuery).Return(expectedResponse, nil)

	// Call the method to be tested
	response, _ := mockTopsecretsplitUseCase.GetAll(satelliteName, satelliteQuery)
	expectresponseuse, _ := usecase.GetAll(satelliteName, satelliteQuery)
	assert.Equal(t, expectresponseuse, response)

}
