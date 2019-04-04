package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-orm/configs"
	"github.com/go-orm/internal/pkg/dbconnect"
	"github.com/go-orm/internal/pkg/logutil"
	"github.com/go-orm/internal/pkg/models"
	"github.com/go-orm/internal/pkg/utils"
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
}
