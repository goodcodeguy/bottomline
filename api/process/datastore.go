package process

import (
	"github.com/goodcodeguy/bottomline/config"
	"github.com/goodcodeguy/bottomline/lib/database"
	"github.com/goodcodeguy/bottomline/lib/logger"
)

var log = logger.New("bottomline.process")
var cfg = config.Configuration()
var db *database.DB = initializeDatabase()

func initializeDatabase() *database.DB {

	dbConfig := database.Config{
		DBHost:    cfg.DBHost,
		DBName:    cfg.DBName,
		DBUser:    cfg.DBUser,
		DBPass:    cfg.DBPass,
		DBSSLMode: cfg.DBSSLMode,
	}

	return database.Open(dbConfig)
}
