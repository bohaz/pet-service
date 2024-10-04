package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Mascota struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

// ConnectDatabase conecta a la base de datos y realiza la migración del modelo.
func ConnectDatabase() error {
	dsn := "host=localhost user=postgres password=samba1702 dbname=pets_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrar el esquema
	return db.AutoMigrate(&Mascota{})
}

// GetDB devuelve la instancia de la base de datos
func GetDB() *gorm.DB {
	return db
}
