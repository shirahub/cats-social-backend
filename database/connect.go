package database

import (
	"app/config"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
	"strconv"
)

// ConnectDB connect to db
func ConnectDB() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}

	DB, err = sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s",
			config.Config("DB_HOST"),
			port,
			config.Config("DB_USERNAME"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_NAME"),
			config.Config("DB_PARAMS"),
		))
	if err != nil {
		log.Error(err.Error())
		panic("failed to open db")
	}

	if err = DB.Ping(); err != nil {
		panic("failed to ping db")
	}
}
