package usecase

import (
	"errors"

	entity "github.com/OperationQuasarFire/internal/pkg/entity"
	"github.com/OperationQuasarFire/internal/pkg/ports"
	entity2 "github.com/OperationQuasarFire/internal/pkg/response"

	"github.com/patrickmn/go-cache"
)

var satelliteDataMap = make(map[string]*entity.SatelliteData)

type TopsecretSplitUseCase struct {
	topsecretsplitService ports.CommunicationServices
	cache                 *cache.Cache
}

func NewTopsecretSplitUseCase(topsecretsplitService ports.CommunicationServices, cache *cache.Cache) *TopsecretSplitUseCase {
	return &TopsecretSplitUseCase{
		topsecretsplitService: topsecretsplitService,
		cache:                 cache,
	}
}

func (uc *TopsecretSplitUseCase) CreateTopSecretSplit(satelliteName *string, satelliteData *entity.SatelliteData) (*entity2.Response, error) {

	uc.cache.Set(*satelliteName, satelliteData, cache.DefaultExpiration)
	count := uc.cache.ItemCount()
	if count > 3 {
		return nil, errors.New("There are enough positions to track the message and the location of the satellites.")
	}

	satelliteDataMap[*satelliteName] = satelliteData

	return &entity2.Response{
		Result: entity2.Result{
			Source:  "Message and distance added successfully",
			Details: nil,
		},
	}, nil
}
func (uc *TopsecretSplitUseCase) GetAll(satelliteName string, satellite *entity.Satellitequery) (*entity2.Response, error) {
	// Obtener los datos de la caché
	cachedData, err := uc.GetCachedData(satelliteName)
	if err != nil {
		return nil, err
	}

	var Position entity2.Position

	// Verificar si existen datos en la caché para el satélite dado
	if cachedData == nil {
		return nil, errors.New("Satellite data not found in cache")
	}

	// Verificar si hay suficientes posiciones de satélite en la caché
	count := uc.cache.ItemCount()
	if count < 3 {
		return nil, errors.New("Insufficient satellite positions in cache")
	}

	// Recopilar todos los mensajes de los satélites
	var allMessages [][]string
	for _, data := range satelliteDataMap {
		allMessages = append(allMessages, data.Message)
	}

	// Calcular la posición basada en los datos de los satélites
	Position.Message = uc.topsecretsplitService.GetMessage(allMessages...)
	alllocation := []float32{}
	for _, data := range satelliteDataMap {
		alllocation = append(alllocation, data.Distance)
	}
	Position.X, Position.Y = uc.topsecretsplitService.GetLocation(alllocation...)

	// Limpiar el mapa de datos de satélites después de obtener la posición
	// Verificar si se encontraron coordenadas y mensaje
	if Position.X == 0 && Position.Y == 0 || Position.Message == "" {
		return nil, errors.New("Insufficient satellite positions")
	}

	return &entity2.Response{
		Position: Position,
		Result: entity2.Result{
			Source:  "Find position to message",
			Details: nil,
		},
	}, nil
}

func (uc *TopsecretSplitUseCase) GetCachedData(satelliteName string) (*entity.SatelliteData, error) {

	cachedData, found := uc.cache.Get(satelliteName)
	if !found {
		return nil, errors.New("Satellite data not found in cache")
	}

	satelliteData, ok := cachedData.(*entity.SatelliteData)
	if !ok {
		return nil, errors.New("Invalid data type in cache")
	}

	return satelliteData, nil
}
