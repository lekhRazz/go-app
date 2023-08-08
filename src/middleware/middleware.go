package customMiddleware

import (
	"context"
	"fmt"

	"sample_go_app/src/config"
	"sample_go_app/src/utilities"

	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
}

func getOrigin(origin []string) string {
	if len(origin) == 0 || origin == nil {
		return ""
	}
	return origin[0]
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func ExtendCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		return next(cc)
	}
}

func SetDebug(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		uuid, err := utilities.GenerateUUID()
		if err != nil {
			fmt.Println("error", err)
		}
		debug := config.Debug{DebugId: uuid, UserAgent: c.Request().Header["User-Agent"][0], Origin: getOrigin(c.Request().Header["Origin"]), Host: c.Request().Host, OriginalUrl: c.Request().RequestURI, Method: c.Request().Method, Ip: c.RealIP()}
		c.Set("debug", debug)
		return next(c)
	}
}

func SetDBConnection(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		dbConnection := utilities.GetDatabaseConnectionPool()
		c.Set("dbConnection", config.DBConnection{DBConnection: dbConnection})
		return next(c)
	}
}

// @Get retrieves data from the context.
func (ctx CustomContext) Get(key string) interface{} {
	val := ctx.Context.Get(key)
	if val != nil {
		return val
	}
	return ctx.Request().Context().Value(key)
}

// @Set saves data in the context.
func (ctx CustomContext) Set(key string, val interface{}) {
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), key, val)))
}
