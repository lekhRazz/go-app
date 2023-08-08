package main

import (
	"fmt"
	"log"
	"net/http"
	"sample_go_app/src/routes"

	customMiddleware "sample_go_app/src/middleware"

	"sample_go_app/src/config"

	"sample_go_app/src/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	envVar := config.DefaultEnvironmentalVariable

	messageConfig := config.DefaultMessageConfigData
	appConfig := config.DefaultAppConfigData

	port := fmt.Sprintf(":%d", envVar.PORT)

	fmt.Println(messageConfig.SERVER_STARTED, port)

	e := echo.New()
	e.Use(customMiddleware.ServerHeader)
	e.Use(customMiddleware.ExtendCustomContext)
	e.Use(customMiddleware.SetDebug)
	e.Use(customMiddleware.SetDBConnection)
	e.Use(customMiddleware.VerifyRequest)

	//@middleware for console logs
	e.Use(middleware.Logger())

	//@middleware for panic recovery i.e. exception handling kind of global catch block
	e.Use(customMiddleware.Recover())

	//@middleware for cors
	e.Use(middleware.CORS())

	//@initiate admin base api route
	adminGroup := e.Group(appConfig.ADMIN_ROUTE)
	routes.InitializeAdminRoutes(adminGroup)

	//@initiate customer base api route
	userRouteGroup := e.Group(appConfig.USER_ROUTE)
	routes.InitializeUserRoutes(userRouteGroup)

	e.GET(appConfig.HEALTH_CHECK_API, func(c echo.Context) error {
		return c.String(http.StatusOK, messageConfig.HEALTH_CHECK_MESSAGE)
	})

	routelevelMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}

	e.Any("/*", func(c echo.Context) error {
		c.Set("HandlerName", "Api-not-found")
		logger.Info(c, "Api not found handler is initiated")
		// debug := c.Get("debug").(config.Debug)

		return c.String(http.StatusNotFound, messageConfig.API_NOT_FOUND)
	}, routelevelMiddleware)

	s := http.Server{
		Addr:    port,
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
