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
		&Request{},
	)

	// Create foreign keys
	db.Model(&Request{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE")

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
		var user User
		user.Name = "testuser"
		if err := user.GetActiveUserName(true, "test", db); err != nil {
			db.Create(&user)
			req := Request{
				User: user,
			}
			db.Create(&req)
			logger.Info("Test data initialized successfully.")
		}
	}
	// ====================================================
}
