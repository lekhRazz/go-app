package customMiddleware

import (
	"fmt"
	"net/http"
	"sample_go_app/src/config"
	"sample_go_app/src/utilities"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		debug := c.Get("debug").(config.Debug)
		response := &HttpResponse{Debug: debug}

		authToken := c.Request().Header["Authorization"][0]

		if len(authToken) == 0 {
			response.Message = "Token Required"
			response.Status = http.StatusUnauthorized
			return c.JSON(http.StatusUnauthorized, response)
		}

		pool := c.Get("dbConnection").(config.DBConnection).DBConnection

		tokenParsed, err := utilities.VerifyJwtToken(authToken, pool)

		if err != nil {
			fmt.Println("jwt error msg", err)

			if err == jwt.ErrTokenExpired {
				response.Message = "JWT token expired"
				response.Status = http.StatusUnauthorized
				return c.JSON(http.StatusUnauthorized, response)
			}

			if err == jwt.ErrSignatureInvalid {
				response.Message = "Invalid JWT token"
				response.Status = http.StatusUnauthorized
				return c.JSON(http.StatusUnauthorized, response)
			}

			response.Message = "JWT token authorization failed"
			response.Status = http.StatusUnauthorized
			return c.JSON(http.StatusUnauthorized, response)
		}

		c.Set("sessionInfo", tokenParsed)
		return next(c)
	}
}
