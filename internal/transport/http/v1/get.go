package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (ctl *Controller) Get(ctx echo.Context) error {
	name := ctx.Param("name")
	satellite, err := ctl.service.GetSatelliteByName(ctx.Request().Context(), name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, satellite)
}
