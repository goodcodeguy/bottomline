package database

import (
	"database/sql"
	"fmt"

	"github.com/goodcodeguy/bottomline/lib/logger"
	_ "github.com/lib/pq" // sql driver
)

// DB holds the database connection
type DB struct {
	*sql.DB
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

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Criticalf("Error when connecting to database: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Criticalf("Error when pinging database: %s", err.Error())
	}

	return &DB{DB: db}
}

// Query Row queries a row
func (db *DB) QueryRow(qry string, args ...interface{}) *sql.Row {
	log.Infof("Query Row: %s", qry)
	return db.DB.QueryRow(qry, args...)
}

// Query does a query
func (db *DB) Query(qry string, args ...interface{}) (*sql.Rows, error) {
	log.Infof("Query: %s", qry)
	rows, err := db.DB.Query(qry, args...)
	if err != nil {
		log.Criticalf("Error Executing Query: %s\n%s", qry, err.Error())
	}
	return rows, err
}

// Exec does a single execute
func (db *DB) Exec(qry string, args ...interface{}) error {
	log.Infof("Exec Query: %s", qry)
	_, err := db.DB.Exec(qry, args...)
	return err
}
