package health

import (
	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/server"
)

type HealthRouterEntry struct{}

func NewHealthRouterEntry() HealthRouterEntry {
	return HealthRouterEntry{}
}

func (h HealthRouterEntry) UseEchoRouter() func(*echo.Group, *server.ServerContext) {
	return func(g *echo.Group, sctx *server.ServerContext) {
		logger, err := sctx.Logger.UseZap()
		if err == nil {
			logger.Info("HealthRouterEntry: UseEchoRouter")
		}
		handler := NewEchoHealthHandler()
		g.GET("/ping", handler.Ping)
		g.GET("/healthy", handler.Healthy)
		g.GET("/check", handler.Check)
	}
}
