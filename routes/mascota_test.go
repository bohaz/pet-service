package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bohaz/pet-service/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Esta función se ejecutará antes de cada prueba para configurar el contexto de la base de datos.
func setupTestDatabase() {
    models.ConnectDatabase()

    // Limpiar la tabla antes de cada prueba
    db := models.GetDB()
    db.Exec("TRUNCATE TABLE mascotas RESTART IDENTITY")
}

// Test para obtener todas las mascotas
func TestGetMascotas(t *testing.T) {
	setupTestDatabase()

	// Crear un nuevo contexto de Gin para simular una solicitud HTTP
	router := gin.Default()
	RegisterMascotaRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mascotas", nil)

	router.ServeHTTP(w, req)

	// Verificar el estado de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Verificar el cuerpo de la respuesta (ejemplo básico)
	expected := "[]" // Si no hay mascotas en la base de datos
	assert.Equal(t, expected, w.Body.String())
}

func TestAddMascota(t *testing.T) {
	setupTestDatabase()

	// Crear un nuevo contexto de Gin para simular una solicitud HTTP
	router := gin.Default()
	RegisterMascotaRoutes(router)

	// Simular una solicitud POST para agregar una mascota
	w := httptest.NewRecorder()
	reqBody := `{"name": "Rex", "breed": "Labrador", "age": 3}`
	req, _ := http.NewRequest("POST", "/mascotas", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	// Verificar el estado de la respuesta
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verificar que la respuesta contenga los datos de la mascota creada
	expected := `{"id":1,"name":"Rex","breed":"Labrador","age":3}`
	assert.JSONEq(t, expected, w.Body.String())
}
