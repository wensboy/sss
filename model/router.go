package model

import (
	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/server"
)

type RouterEntry interface {
	UseEchoRouter() func(*echo.Group, *server.ServerContext)
}
