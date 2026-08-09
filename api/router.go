package api

import (
	"github.com/labstack/echo/v5"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
	"github.com/wensboy/ss/server"
)

type SwaggerRouterEntry struct{}

func NewSwaggerRouterEntry() SwaggerRouterEntry {
	return SwaggerRouterEntry{}
}

func (s SwaggerRouterEntry) UseEchoRouter() func(*echo.Group, *server.ServerContext) {
	return func(g *echo.Group, sctx *server.ServerContext) {
		g.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}

type ScalarRouterEntry struct{}

func NewScalarRouterEntry() ScalarRouterEntry {
	return ScalarRouterEntry{}
}

func (s ScalarRouterEntry) UseEchoRouter() func(*echo.Group, *server.ServerContext) {
	return func(g *echo.Group, sctx *server.ServerContext) {
		apiHandler := NewEchoApiHandler()
		g.GET("/scalar", apiHandler.ScalarWrapHandler)
	}
}
