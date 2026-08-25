package rest

import (
	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/server"
	"github.com/wensboy/sss/rbac"
)

type RbacRouterEntry struct{}

func NewRbacRouterEntry() RbacRouterEntry {
	return RbacRouterEntry{}
}

func (r RbacRouterEntry) UseEchoRouter() func(*echo.Group, *server.ServerContext) {
	return func(g *echo.Group, sctx *server.ServerContext) {
		repo := rbac.NewRbacRepo().SetDB(sctx.DBContext.MustGet("localhost:3306/rbac").DB())
		service := rbac.NewRbacService(repo)
		_ = NewEchoRbacHandler(service)
		// todo: 完善端点路由
	}
}
