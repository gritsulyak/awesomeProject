package v1

import (
	"context"

	"github.com/labstack/echo/v4"

	model "github.com/gritsulyak/awesomeProject/internal/model"
)

type SatelliteService interface {
	GetSatelliteByName(ctx context.Context, name string) (*model.Satellite, error)
}

type Controller struct {
	service SatelliteService
}

func NewController(g *echo.Group, s SatelliteService) {
	ctl := &Controller{service: s}
	g.GET("/:name", ctl.Get)
}
