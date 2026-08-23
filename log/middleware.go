package log

import (
	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/server"
)

const (
	VarKey_LogEnabled       = "log::var::enabled"
	ToolKey_LogMutateLogger = "log::tool::mutate_logger"
)

func LogWithZap(mc server.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			err := next(c)
			if err != nil {
				return err
			}

		}
	}
}
