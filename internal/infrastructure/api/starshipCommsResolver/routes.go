package starshipcommsresolver

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"

	repository "github.com/OperationQuasarFire/internal/infrastructure/repository"
	services "github.com/OperationQuasarFire/internal/pkg/service"
)

func RegisterRoutes2(e *gin.Engine, db *sql.DB) {

	cacheExpiration := 5 * time.Minute
	cacheCleanupInterval := 10 * time.Minute
	cacheInstance := cache.New(cacheExpiration, cacheCleanupInterval)

	SatelliteRepository := repository.NewBdRepository(db)
	topsecretService := services.NewService(nil)
	topsecretHandler := newHandler(topsecretService, SatelliteRepository)
	var topsecretsplitService = services.NewServiceSplit(nil)
	topsecretSplitHandler := newSplitHandler(topsecretsplitService, cacheInstance)

	e.POST("/topsecret/", topsecretHandler.topsecret)
	e.POST("/topsecret_split/:satellite_name", topsecretSplitHandler.topsecretSplit)
	e.GET("/topsecret_split/:satellite_name", topsecretSplitHandler.GetsecretSplit)

}
