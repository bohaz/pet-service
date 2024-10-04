package main

import (
	"log"

	"github.com/bohaz/pet-service/config"
	"github.com/bohaz/pet-service/models"
	"github.com/bohaz/pet-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	err := models.ConnectDatabase()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Inicializar Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(config.SetupCORS())

	// Registrar rutas
	routes.RegisterMascotaRoutes(r)

	// Ejecutar el servidor
	r.Run(":8080")
}
