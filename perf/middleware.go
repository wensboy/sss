package perf

import (
	"time"

	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/server"
)

const (
	VarKey_PerfEnabled  = "perf::var::enabled"
	VarKey_PerfLinkScan = "perf::var::link_scan"
)

// PerfLinkScan 用于计算链路scan并设置到上下文.
func PerfLinkScan(mc server.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if !mc.MustGet(VarKey_PerfEnabled).(bool) {
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
