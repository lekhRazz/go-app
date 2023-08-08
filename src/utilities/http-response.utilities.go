package utilities

import (
	"net/http"
	"sample_go_app/src/config"
	"sample_go_app/src/logger"

	"github.com/labstack/echo"
)

func InternalServerError(ctx echo.Context, err error) error {
	logger.Error(ctx, config.DefaultMessageConfigData.INTERNAL_SERVER_ERROR, err)
	return ctx.String(http.StatusInternalServerError, config.DefaultMessageConfigData.INTERNAL_SERVER_ERROR)
}
