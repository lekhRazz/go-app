package onboardService

import (
	"github.com/labstack/echo"
)

func MobileVerifyHandler(c echo.Context) error {
	return c.String(200, DefaultMessageConfigData.MobileVerifySuccess)
}
