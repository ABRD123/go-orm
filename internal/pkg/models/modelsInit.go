package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/go-orm/configs"
)

// CreateTables creates the DB Tables.
func CreateTables(db *gorm.DB, logger *log.Logger) {
	// Create tables
	db.AutoMigrate(
		&User{},
	)

	logger.Info("Tables created successfully.")
}

// InitializeTables initializes tables with default data.
func InitializeTables(db *gorm.DB, logger *log.Logger) {
	// ====================================================
	// ====================================================
	// Enter dummy data for testing purposes below.
	// ====================================================
	// ====================================================
	if !configs.Live {
		user := User{}
		user.ID = 123456
		user.Name = "abrdtest"
		user.Active = true
		db.Where(User{Name: user.Name}).FirstOrCreate(&user)

		logger.Info("Test data initialized successfully.")
	}
	// ====================================================
}
