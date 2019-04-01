package dbconnect

import (
	_ "github.com/go-sql-driver/mysql" // MySql Driver
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/go-orm/configs"
)

// GetDBConn is getter for DBConn - make sure the env vars are set properly.
func GetDBConn() *DBConn {
	var dbPassword string
	if configs.Live {
		dbPassword = configs.DBProdPassword
	} else {
		dbPassword = configs.DBBetaPassword
	}

	d := DBConn{
		DBUser:      configs.DBUsername,
		DBPassword:  dbPassword,
		DBAuthority: configs.DBAuthority,
		DBName:      configs.DBName}
	return &d
}

// DBConn handles the connection vars for the DB.
type DBConn struct {
	DBUser      string // database username
	DBPassword  string // database password
	DBAuthority string // ex. "127.0.0.1:3306"
	DBName      string // ex. "goorm"
}

// Connect to DB using DBConn data and return pointer to the DB or error if error occurs.
//
// NOTE: The caller is required to close the DB connection when finished with it.
//  Returns pointer to DB and Error.
func (dd *DBConn) Connect(logger *log.Logger) (*gorm.DB, error) {
	dbConnect := dd.DBUser + ":" + dd.DBPassword + "@tcp(" + dd.DBAuthority + ")/" + dd.DBName + "?parseTime=true"
	db, err := gorm.Open("mysql", dbConnect)
	if err != nil {
		logger.Error(err)
	}
	db.SetLogger(logger)
	return db, err
}
