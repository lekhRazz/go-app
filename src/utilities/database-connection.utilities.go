package utilities

import (
	"database/sql"
	"fmt"
	"sample_go_app/src/config"
	"time"

	_ "github.com/lib/pq"
)

func getDbConnectionUrl(envVar config.EnvironmentalVariable) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", envVar.DB_USER, envVar.DB_PASSWORD, envVar.DB_HOST, envVar.DB_PORT, envVar.DB_NAME)
}

func GetDatabaseConnectionPool() *sql.DB {
	envVar := config.DefaultEnvironmentalVariable

	dbConnectionUrl := getDbConnectionUrl(envVar)
	db, err := sql.Open("postgres", dbConnectionUrl)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("DB Connected!")

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Second * 20)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Second * 20)
	return db
}

func GetDatabasePersistentConnection(db *sql.DB) *sql.DB {
	// db.Begin()
	db.SetConnMaxLifetime(time.Second * 40)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	return db
}
