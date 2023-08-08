package onboardService

import (
	"github.com/labstack/echo"
)

func SignupHandler(c echo.Context) error {
	return c.String(200, DefaultMessageConfigData.SignUpSuccess)
}
