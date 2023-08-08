package userOnboardRoutes

type DefaultRouteConfig struct {
	BaseRoute         string
	EmailVerifyRoute  string
	MobileVerifyRoute string
	SignUpRoute       string
}

var DefaultRouteConfigData = DefaultRouteConfig{
	BaseRoute:         "/onboarding",
	EmailVerifyRoute:  "/verify/email",
	MobileVerifyRoute: "/verify/mobile",
	SignUpRoute:       "/signup",
}
