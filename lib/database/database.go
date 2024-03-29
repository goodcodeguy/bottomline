package database

import (
	"fmt"

	"github.com/goodcodeguy/bottomline/lib/logger"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// DB holds the database connection
type DB struct {
	*sqlx.DB
}

type Model struct {
	Id        uint        `json:"id"`
	CreatedAt pq.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at" json:"updated_at"`
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

	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		panic("Error connecting to database")
	}

	return &DB{DB: db}
}
