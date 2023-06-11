package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
	"github.com/vanneeza/e-commerce-lite/utils/pkg"
)

func InitDB() (*sql.DB, error) {
	Host := pkg.GetEnv("DB_HOST")
	Port := pkg.GetEnv("DB_PORT")
	User := pkg.GetEnv("DB_USER")
	Password := pkg.GetEnv("DB_PASSWORD")
	DbName := pkg.GetEnv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
	db, err := sql.Open("postgres", connStr)
	helper.PanicError(err)

	err = db.Ping()
	helper.PanicError(err)

	log.Println("connection to database successfuly")
	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Printf("error closing database connection : %s", err)

	} else {
		log.Println("database connection closed")
	}
}
