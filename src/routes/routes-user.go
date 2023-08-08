package routes

import (
	userOnboardRoutes "sample_go_app/src/feature-modules/user-onboard"

	"github.com/labstack/echo"
)

func InitializeUserRoutes(group *echo.Group) {

	//@make subroute for onboard module
	userOnboardRoutesGroup := group.Group(userOnboardRoutes.DefaultRouteConfigData.BaseRoute)
	userOnboardRoutes.InitializeUserOnboardRoute(userOnboardRoutesGroup)
}
