package config

import (
	"fmt"
	"os"

	"strconv"

	"github.com/joho/godotenv"
)

type EnvironmentalVariable struct {
	PORT                    int
	ELASTICSEARCH_LOG_INDEX string
	ELASTICSEARCH_END_URL   string
	ELASTICSEARCH_USER      string
	ELASTICSEARCH_PASSWORD  string
	DB_USER                 string
	DB_PASSWORD             string
	DB_HOST                 string
	DB_PORT                 string
	DB_NAME                 string
	APP_SECRET_KEY          string
}
type AppConfig struct {
	ADMIN_ROUTE      string
	USER_ROUTE       string
	HEALTH_CHECK_API string
}

type MessageConfig struct {
	INTERNAL_SERVER_ERROR string
	DATA_FETCH            string
	SERVER_STARTED        string
	HEALTH_CHECK_MESSAGE  string
	API_NOT_FOUND         string
}

var DefaultEnvironmentalVariable = EnvironmentalVariable{}

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("error", err)
	}

	DefaultEnvironmentalVariable.PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("error", err)
	}

	DefaultEnvironmentalVariable.ELASTICSEARCH_END_URL = os.Getenv("ELASTICSEARCH_END_URL")
	DefaultEnvironmentalVariable.ELASTICSEARCH_LOG_INDEX = os.Getenv("ELASTICSEARCH_LOG_INDEX")
	DefaultEnvironmentalVariable.ELASTICSEARCH_USER = os.Getenv("ELASTICSEARCH_USER")
	DefaultEnvironmentalVariable.ELASTICSEARCH_PASSWORD = os.Getenv("ELASTICSEARCH_PASSWORD")
	DefaultEnvironmentalVariable.DB_HOST = os.Getenv("DB_HOST")
	DefaultEnvironmentalVariable.DB_PORT = os.Getenv("DB_PORT")
	DefaultEnvironmentalVariable.DB_NAME = os.Getenv("DB_NAME")
	DefaultEnvironmentalVariable.DB_USER = os.Getenv("DB_USER")
	DefaultEnvironmentalVariable.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DefaultEnvironmentalVariable.APP_SECRET_KEY = os.Getenv("APP_SECRET_KEY")
}

var DefaultAppConfigData = AppConfig{
	ADMIN_ROUTE:      "/api/v1/admin",
	USER_ROUTE:       "/api/v1/user",
	HEALTH_CHECK_API: "/api/health-check",
}

var DefaultMessageConfigData = MessageConfig{
	INTERNAL_SERVER_ERROR: "OOPS!!! Something went wrong.",
	DATA_FETCH:            "Data fetched.",
	SERVER_STARTED:        "Server started on PORT",
	HEALTH_CHECK_MESSAGE:  "I'M OKAY SIR!!! ",
	API_NOT_FOUND:         "SORRY !!! API not found",
}
