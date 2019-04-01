package main

import (
	"github.com/go-orm/configs"
	"github.com/go-orm/internal/pkg/dbconnect"
	"github.com/go-orm/internal/pkg/logutil"
	"github.com/go-orm/internal/pkg/models"
	"github.com/go-orm/internal/pkg/utils"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup logging
	logger := log.New()
	file := logutil.SetupLogging(logger, configs.DBActionsLogName)
	defer utils.Close(file, logger)

	// Setup db handle.
	dbConn := dbconnect.GetDBConn()
	db, err := dbConn.Connect(logger)
	if err != nil {
		logger.Panic(err)
	}
	defer utils.Close(db, logger)

	// Get active user details
	var user models.User
	err = user.GetActiveUserName(true, "abrdtest", db)
	if err != nil {
		logger.Warn(err)
	}
	if user.Active {
		logger.Info("Successfully retrieved user details!!!")
		// Update user as inactive
		err = user.UpdateActive(false, db)
		if err != nil {
			logger.Warn(err)
		}
		logger.Info("Successfully updated user as inactive!!!")
	}
}
