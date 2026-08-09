package health

import (
	"github.com/labstack/echo/v5"
	"github.com/wensboy/sss/model"
)

type EchoHealthHandler interface {
	Ping(*echo.Context) error
	Healthy(*echo.Context) error
	Check(*echo.Context) error
}

type echoHealthHandler struct {
}

func NewEchoHealthHandler() EchoHealthHandler {
	return &echoHealthHandler{}
}

// Ping godoc
//
//	@Summary		Ping
//	@Description	健康检查 - Ping端点，用于检测服务是否可达，返回 "pong"
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.RestResponse	"成功响应，code=0, msg=\"pong\", data=nil"
//	@Router			/ping [get]
func (h *echoHealthHandler) Ping(c *echo.Context) error {
	return model.UseEchoResponder(c).Success("pong", nil)
}

// Healthy godoc
//
//	@Summary		Healthy
//	@Description	健康检查 - Healthy端点，用于检测服务是否处于健康状态，返回 "healthy"
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.RestResponse	"成功响应，code=0, msg=\"healthy\", data=nil"
//	@Router			/healthy [get]
func (h *echoHealthHandler) Healthy(c *echo.Context) error {
	return model.UseEchoResponder(c).Success("healthy", nil)
}

// Check godoc
//
//	@Summary		Check
//	@Description	健康检查 - Check端点，用于综合检查服务各组件状态，返回 "check"
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.RestResponse	"成功响应，code=0, msg=\"check\", data=nil"
//	@Router			/check [get]
func (h *echoHealthHandler) Check(c *echo.Context) error {
	return model.UseEchoResponder(c).Success("check", nil)
}
