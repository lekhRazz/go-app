package onboardService

import (
	"github.com/labstack/echo"
)

func EmailVerifyHandler(c echo.Context) error {
	return c.String(200, DefaultMessageConfigData.EmailVerifySuccess)
}
