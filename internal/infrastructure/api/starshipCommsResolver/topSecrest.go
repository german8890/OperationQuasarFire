package starshipcommsresolver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/OperationQuasarFire/internal/pkg/entity"
	"github.com/OperationQuasarFire/internal/pkg/ports"
	entity2 "github.com/OperationQuasarFire/internal/pkg/response"
	"github.com/OperationQuasarFire/internal/pkg/usecase"
)

type topsecretHandler struct {
	topsecretService ports.CommunicationServices
	topsecretUseCase *usecase.TopsecrestUseCase
	repo             ports.DBRepository
}

func newHandler(service ports.CommunicationServices, repo ports.DBRepository) *topsecretHandler {
	orderUseCase := usecase.NewTopsecretUseCase(service, repo)
	return &topsecretHandler{
		topsecretService: service,
		topsecretUseCase: orderUseCase,
	}
}

func (o *topsecretHandler) topsecret(c *gin.Context) {

	var satellite entity.SatelliteRequest
	if err := c.BindJSON(&satellite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Input "})
		return
	}

	entityResponse, err := o.topsecretUseCase.CreateTopSecret(&satellite)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	entityResponse.Result.Details = []entity2.Detail{{InternalCode: strconv.Itoa(http.StatusOK), Message: http.StatusText(http.StatusOK)}}

	c.Set("entityResponse", *entityResponse)
	c.JSON(http.StatusOK, entityResponse)

}
