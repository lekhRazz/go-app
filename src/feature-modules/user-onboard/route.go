package userOnboardRoutes

import (
	onboardService "sample_go_app/src/feature-modules/user-onboard/services"

	"github.com/labstack/echo"
)

func InitializeUserOnboardRoute(group *echo.Group) {

	//@handle login route
	group.POST(DefaultRouteConfigData.SignUpRoute, onboardService.SignupHandler)

	//@handle email verify otp route
	group.POST(DefaultRouteConfigData.EmailVerifyRoute, onboardService.EmailVerifyHandler)

	//@handle mobile verify otp route
	group.POST(DefaultRouteConfigData.MobileVerifyRoute, onboardService.MobileVerifyHandler)
}
