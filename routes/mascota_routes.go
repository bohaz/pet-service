package routes

import (
	"net/http"

	"github.com/bohaz/pet-service/models"

	"github.com/gin-gonic/gin"
)

func RegisterMascotaRoutes(r *gin.Engine) {
	r.GET("/mascotas", getMascotas)
	r.POST("/mascotas", addMascota)
	r.GET("/mascotas/:id", getMascotaByID)
	r.PUT("/mascotas/:id", updateMascota)
	r.DELETE("/mascotas/:id", deleteMascota)
}

func getMascotas(c *gin.Context) {
	var mascotas []models.Mascota
	models.GetDB().Find(&mascotas)
	c.JSON(http.StatusOK, mascotas)
}

func addMascota(c *gin.Context) {
	var newMascota models.Mascota
	if err := c.ShouldBindJSON(&newMascota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.GetDB().Create(&newMascota)
	c.JSON(http.StatusCreated, newMascota)
}

func getMascotaByID(c *gin.Context) {
	id := c.Param("id")
	var mascota models.Mascota
	if err := models.GetDB().First(&mascota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Mascota no encontrada"})
		return
	}
	c.JSON(http.StatusOK, mascota)
}

func updateMascota(c *gin.Context) {
	id := c.Param("id")
	var mascota models.Mascota
	if err := models.GetDB().First(&mascota, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Mascota no encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&mascota); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.GetDB().Save(&mascota)
	c.JSON(http.StatusOK, mascota)
}

func deleteMascota(c *gin.Context) {
	id := c.Param("id")
	if err := models.GetDB().Delete(&models.Mascota{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Mascota no encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mascota eliminada"})
}
