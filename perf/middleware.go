package perf

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/log"
	"github.com/wensboy/ss/server"
	"github.com/wensboy/sss/model"
	"go.uber.org/zap"
)

const (
	VarKey_PerfLinkScan        = "perf::var::link_scan"
	VarKey_PerfLinkScanEnabled = "perf::var::link_scan_enabled"

	VarKey_PerfRecoverEnabled = "perf::var::recover_enabled"
)

// PerfLinkScan 用于计算链路scan并设置到上下文.
func PerfLinkScan(mc server.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if !mc.MustGet(VarKey_PerfLinkScanEnabled).(bool) {
				return next(c)
			}
			start := time.Now()
			err := next(c)
			if err != nil {
				return err
			}
			scan := time.Since(start)
			c.Set(VarKey_PerfLinkScan, scan)
			return nil
		}
	}
}

// PerfRecover 用于捕获并处理链路中的panic.
// todo: 1) 变更 logger 类型 2) 优化 error 处理
func PerfRecover(mc server.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if mc.MustGet(VarKey_PerfRecoverEnabled).(bool) {
				defer func() {
					if err := recover(); err != nil {
						logger, _ := log.GetGMutateLogger().UseZap()
						logger.Error("panic recover", zap.Error(err.(error)))
						model.UseEchoResponder(c).Err(http.StatusInternalServerError, -1, "internal server panic")
					}
				}()
			}
			return next(c)
		}
	}
}
