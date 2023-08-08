package customMiddleware

import (
	"fmt"
	"net/http"
	"regexp"
	"sample_go_app/src/config"
	"sample_go_app/src/utilities"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type HttpResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Debug   config.Debug `json:"debug"`
}

func VerifyRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		debug := c.Get("debug").(config.Debug)
		response := &HttpResponse{Debug: debug}

		authHeader := c.Request().Header["X-Api-Key"][0]
		signature, timestamp, err := splitAPIKey(authHeader)

		if err != nil {
			fmt.Println("err in timestamp", err)
		}

		if expired := checkRequestExpiry(timestamp); expired == true {
			response.Message = "Token expired"
			response.Status = http.StatusUnauthorized
			return c.JSON(http.StatusUnauthorized, response)
		}

		if isValidToken := validateToken(signature, timestamp); isValidToken == false {
			response.Message = "Invalid auth token"
			response.Status = http.StatusUnauthorized
			return c.JSON(http.StatusUnauthorized, response)
		}

		return next(c)
	}
}

func splitAPIKey(apiKey string) (singature string, timestamp int64, err error) {
	pattern := regexp.MustCompile("Signature=([A-Za-z0-9]+),Timestamp=([0-9]+)")
	match := pattern.FindAllStringSubmatch(apiKey, -1)
	singature = match[0][1]
	Timestamp, err := strconv.Atoi(match[0][2])
	if err != nil {
		return
	}
	timestamp = int64(Timestamp)
	return
}

func checkRequestExpiry(timestamp int64) bool {

	tokenValidTime := int64(10 * 60)
	if time.Now().Unix()-timestamp < tokenValidTime {
		return false
	}
	return true
}

func validateToken(signature string, timestamp int64) bool {

	envVar := config.DefaultEnvironmentalVariable
	appSecretKey := envVar.APP_SECRET_KEY

	hashPlainText := fmt.Sprintf("%s%v", appSecretKey, timestamp)

	hash := utilities.GenerateHash(hashPlainText)

	if equalHash := utilities.CompareHash([]byte(hash), []byte(signature)); equalHash == true {
		return true
	}
	return false
}
