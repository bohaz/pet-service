package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Mascota struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Name  string `json:"name"`
    Breed string `json:"breed"`
    Age   int    `json:"age"`
}

var db *gorm.DB

func main() {
    var err error

    // Conectar a PostgreSQL
    dsn := "host=localhost user=postgres password=samba1702 dbname=pets_db port=5432 sslmode=disable"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    // Migrar el esquema
    db.AutoMigrate(&Mascota{})

    r := gin.Default()

			// Configurar CORS
			r.Use(cors.New(cors.Config{
				AllowOrigins:     []string{"http://localhost:3000"}, // Permitir solicitudes de localhost:3000
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
				AllowHeaders:     []string{"Origin", "Content-Type"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
		}))

    // Rutas CRUD
		r.GET("/mascotas", getMascotas)
    r.POST("/mascotas", addMascota)
    r.GET("/mascotas/:id", getMascotaByID)
    r.PUT("/mascotas/:id", updateMascota)
    r.DELETE("/mascotas/:id", deleteMascota)

    r.Run(":8080")
}

func getMascotas(c *gin.Context) {
	var mascotas []Mascota
	db.Find(&mascotas)
	c.JSON(http.StatusOK, mascotas)
}

func addMascota(c *gin.Context) {
	var newMascota Mascota
	if err := c.ShouldBindJSON(&newMascota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	db.Create(&newMascota)
	c.JSON(http.StatusCreated, newMascota)
}

func getMascotaByID(c *gin.Context) {
	id := c.Param("id")
	var mascota Mascota
	if err := db.First(&mascota, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Mascota no encontrada"})
			return
	}
	c.JSON(http.StatusOK, mascota)
}

func updateMascota(c *gin.Context) {
	id := c.Param("id")
	var mascota Mascota
	if err := db.First(&mascota, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Mascota no encontrada"})
			return
	}

	if err := c.ShouldBindJSON(&mascota); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	db.Save(&mascota)
	c.JSON(http.StatusOK, mascota)
}

func deleteMascota(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&Mascota{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Mascota no encontrada"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mascota eliminada"})
}