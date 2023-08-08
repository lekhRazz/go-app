package routes

import (
	authRoutes "sample_go_app/src/feature-modules/auth"

	"github.com/labstack/echo"
)

func InitializeAdminRoutes(group *echo.Group) {

	//@make subroute for auth module
	adminAuthRouteGroup := group.Group("/auth")
	authRoutes.InitializeAdminAuthRoutes(adminAuthRouteGroup)

}
