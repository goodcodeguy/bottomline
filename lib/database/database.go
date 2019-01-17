package database

import (
	"fmt"

	"github.com/goodcodeguy/bottomline/lib/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB holds the database connection
type DB struct {
	*gorm.DB
}

var log = logger.New("bottomline.database")

// Config is a configuration object for the database
type Config struct {
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
}

// Open provides initialized database
func Open(cfg Config) *DB {

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBSSLMode)

	log.Infof("Opening connection to database (%s)", cfg.DBHost)

	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		panic("failed to connect database")
	}
	db.Exec("SET search_path TO bottomline")
	db.LogMode(true)

	return &DB{DB: db}
}
