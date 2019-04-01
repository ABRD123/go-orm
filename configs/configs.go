package configs

import (
	"os"
	"strconv"
	"strings"
)

// ***************************
// *** Database - go-orm ***
// ***************************

// DBAuthority database authority ex. "127.0.0.1:3306"
var DBAuthority = "127.0.0.1:3306"

// DBInitLogName is the filename to use for logging.
const DBInitLogName = "dbinit.log"

// DBActionsLogName is the filename to use for logging.
const DBActionsLogName = "dbactions.log"

// DBName name of the go-orm database.
const DBName = "goorm"

// DBProdPassword Prod password
const DBProdPassword = ""

// DBBetaPassword Beta password
const DBBetaPassword = "goorm"

// DBUsername User Name
const DBUsername = "goorm"

// Live - DB Live Flag
const Live = false

// *****************************
// *** Common ***
// *****************************

// ProdLogDebug is used to force debug level logs in production.
// If true and Live then logs at Debug level else if false and Live then Info level.
var ProdLogDebug = GetBoolEnvVar("GO_DAEMONS_PROD_LOG_DEBUG", false)

// LogPath is the absolute directory where the HelloWorld log files are located at.
const LogPath = "/var/log/goorm"

// GetBoolEnvVar will get the environment variable for envVarName and attempt to cast it to a bool.
// If it fails or does not exist then uses the defaultValue.
func GetBoolEnvVar(envVarName string, defaultValue bool) bool {
	result := defaultValue
	value, exists := os.LookupEnv(envVarName)
	if exists {
		tempValue, err := strconv.ParseBool(strings.ToLower(value))
		if err == nil {
			result = tempValue
		}
	}
	return result
}
