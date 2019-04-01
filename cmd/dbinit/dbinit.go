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
	file := logutil.SetupLogging(logger, configs.DBInitLogName)
	defer utils.Close(file, logger)

	// Setup db handle.
	dbConn := dbconnect.GetDBConn()
	db, err := dbConn.Connect(logger)
	if err != nil {
		logger.Panic(err)
	}
	defer utils.Close(db, logger)

	models.CreateTables(db, logger)
	models.InitializeTables(db, logger)

	user := &models.User{}
	user.ID = 123456
	user.Name = "abrdtest321"
	user.Active = true
	db.Where(&models.User{Name: user.Name}).FirstOrCreate(&user)
}
