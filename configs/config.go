package configs

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (db *sql.DB, err error) {

	dbDriver := "mysql"
	// dbName := "asolwkqwi123"
	// dbUser := "asolwkqwi123"
	// dbPass := "6N@q5KKq6j$6$iS"
	// dbHost := "db4free.net"

	// dbName := "go_miniproject_dts"
	// dbUser := "root"
	// dbPass := ""
	// dbPort := ""
	// dbHost := "localhost"

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "asolwkqwi123"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "asolwkqwi123"
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		dbPass = "6N@q5KKq6j$6$iS"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "db4free.net"
	}

	// db, err = sql.Open("mysql", "asolwkqwi123:6N@q5KKq6j$6$iS@tcp(db4free.net:3306)/asolwkqwi123")
	// db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName)
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		log.Printf("Errors %s open DB", err)
		panic(err)
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return
	}

	// See "Important settings" section.
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(20 * time.Minute)

	return
}
