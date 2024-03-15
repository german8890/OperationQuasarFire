package starshipcommsresolver

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes2(t *testing.T) {

	router := gin.Default()
	db, err := sql.Open("sqlite3", "user:password@/dbname")
	if err != nil {
		t.Fatalf("Error al crear la instancia de sql.DB: %v", err)
	}
	defer db.Close()

	RegisterRoutes2(router, db)

	reqBody := []byte(`{"key": "value"}`)
	req, err := http.NewRequest("POST", "/topsecret/", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Error al crear la solicitud de prueba: %v", err)
	}

	respRecorder := httptest.NewRecorder()

	router.ServeHTTP(respRecorder, req)
	assert.Equal(t, http.StatusNotFound, respRecorder.Code, "Código de estado esperado: 404 Not Found")
	routes := router.Routes()
	assert.True(t, routeExists(routes, "POST", "/topsecret/"), "La ruta '/topsecret/' debería estar registrada en el enrutador")
}

func routeExists(routes gin.RoutesInfo, method, path string) bool {
	for _, route := range routes {
		if route.Method == method && route.Path == path {
			return true
		}
	}
	return false
}
