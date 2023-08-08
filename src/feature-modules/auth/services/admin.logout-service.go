package authService

import (
	"net/http"
	"sample_go_app/src/config"
	logoutDBSvc "sample_go_app/src/feature-modules/auth/db-services/logout-db-svc"

	"github.com/labstack/echo"
)

func AdminLogoutHandler(c echo.Context) error {
	debug := c.Get("debug").(config.Debug)
	sessionInfo := c.Get("sessionInfo").(config.SessionInfo)

	response := &HttpResponse{Debug: debug}

	pool := c.Get("dbConnection").(config.DBConnection).DBConnection
	dbSvcs := logoutDBSvc.NewLogoutDBSvcs(pool)

	err := dbSvcs.LogoutUser(sessionInfo.UserId, sessionInfo.SessionId)
	if err != nil {
		response.Message = "logout failed"
		response.Status = http.StatusBadRequest
		response.Token = ""
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "logout success"
	response.Status = http.StatusOK
	response.Token = ""
	return c.JSON(http.StatusOK, response)
}
