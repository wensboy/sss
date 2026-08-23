package log

import (
	"html/template"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/log"
	"github.com/wensboy/ss/server"
	"go.uber.org/zap"
)

const (
	VarKey_LogEnabled       = "log::var::enabled"
	VarKey_LogTemplate      = "log::var::template"
	VarKey_LogTemplateCtx   = "log::var::template_ctx"
	ToolKey_LogMutateLogger = "log::tool::mutate_logger"
)

func LogWithZap(mc server.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			err := next(c)
			if mc.MustGet(VarKey_LogEnabled).(bool) {
				mlogger := mc.MustGet(ToolKey_LogMutateLogger).(*log.MutateLogger)
				tmpl := mc.MustGet(VarKey_LogTemplate).(*template.Template)
				tmplCtx := mc.MustGet(VarKey_LogTemplateCtx).(TemplateContext)
				logger, _ := mlogger.UseZap()
				var str strings.Builder
				if err := tmpl.Execute(&str, tmplCtx); err != nil {
					logger.Error("failed to execute log template", zap.Error(err))
				} else {
					logger.Info(str.String())
				}
			}
			return err
		}
	}
}
