package repository

import (
	"app/config"
	"database/sql"
	"fmt"
	"os"
	"github.com/rs/zerolog"
	"github.com/gofiber/fiber/v2/log"
	"github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	_ "github.com/lib/pq"
	"strconv"
)

func NewPostgresConnection() *sql.DB {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	address := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USERNAME"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
		config.Config("DB_PARAMS"),
	)

	db, err := sql.Open("postgres", address)
	if err != nil {
		log.Error(err.Error())
		panic("failed to open db")
	}

	dbLogLevel := sqldblogger.LevelInfo
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	db = sqldblogger.OpenDriver(
		address, db.Driver(), loggerAdapter, sqldblogger.WithMinimumLevel(dbLogLevel),
	)

	if err = db.Ping(); err != nil {
		panic("failed to ping db")
	}

	return db
}

func CloseDbConn(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}