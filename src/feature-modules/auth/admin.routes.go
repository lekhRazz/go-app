package authRoutes

import (
	authService "sample_go_app/src/feature-modules/auth/services"

	customMiddleware "sample_go_app/src/middleware"

	"github.com/labstack/echo"
)

func InitializeAdminAuthRoutes(group *echo.Group) {

	//@handle login route
	group.POST(loginRoute, authService.AdminLoginHandler)

	group.Use(customMiddleware.Auth)

	//@handle logout route
	group.PUT(logoutRoute, authService.AdminLogoutHandler)
}
