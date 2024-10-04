package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS configura las reglas de CORS
func SetupCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
}
