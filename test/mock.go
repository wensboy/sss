package test

import (
	"html/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/spf13/cast"
	"github.com/wensboy/ss/config"
	"github.com/wensboy/ss/log"
	"github.com/wensboy/ss/server"
	"github.com/wensboy/sss/api"
	"github.com/wensboy/sss/health"
	log3 "github.com/wensboy/sss/log"
	"github.com/wensboy/sss/perf"
)

type Mocker struct{}

func (m *Mocker) LoadCommand(path string) {
	_ = config.InitCommand(path)
}

func (m *Mocker) LoadEnv(paths ...string) {
	err := godotenv.Load(paths...)
	if err != nil {
		panic(err)
	}
}

func (m *Mocker) LoadConfig(dir string, paths ...string) {
	cfg := config.NewConfig().SetDir(dir)
	config.SetGlobalConfig(cfg)
	cfg.Load(paths...)
}

func (m *Mocker) RestServer(basePath string) *server.RestServer {
	restServer := server.NewRestServer()
	restServer.MountModules(
		restServer.MountLogger(nil, nil, func(rs *server.RestServer) {
			log.SetGMutateLogger(restServer.Scontext.GetLogger())
		}),
		restServer.MountMContext(nil, nil, func(rs *server.RestServer) {
			enabled := config.MustLookup[bool](
				config.GConfigSource("server.rest.middleware.log.enabled"),
				config.DefaultSource(true),
			)
			format := config.MustLookup[string](
				config.GConfigSource("server.rest.middleware.log.template"),
			)
			tmplList := config.MustLookupRaw(
				config.GConfigSource("server.rest.middleware.log.templateCtx"),
				config.DefaultSource([]string{}),
			)
			tmpl, _ := template.New("mocker").Parse(format)
			rs.Scontext.MContext.Set(log3.VarKey_LogEnabled, enabled)
			rs.Scontext.MContext.Set(log3.VarKey_LogTemplate, tmpl)
			rs.Scontext.MContext.Set(log3.VarKey_LogTemplateCtx, cast.Must[[]string](cast.ToStringSliceE(tmplList)))
			rs.Scontext.MContext.Set(log3.ToolKey_LogMutateLogger, rs.Scontext.GetLogger())
			rs.Scontext.MContext.Set(log3.ToolKey_LogTemplateCtxHook, func(c *echo.Context, tmplCtx log3.TemplateContext, err error) {
				tmplCtx["timestamp"] = c.Request().Header.Get("X-Request-Start")
				tmplCtx["method"] = c.Request().Method
				tmplCtx["uri"] = c.Request().RequestURI
				_, tmplCtx["statusCode"] = echo.ResolveResponseStatus(c.Response(), err)
			})
			enabled = config.MustLookup[bool](
				config.GConfigSource("server.rest.perf.recover"),
				config.DefaultSource(true),
			)
			rs.Scontext.MContext.Set(perf.VarKey_PerfRecoverEnabled, enabled)
			enabled = config.MustLookup[bool](
				config.GConfigSource("server.rest.perf.lagency"),
				config.DefaultSource(true),
			)
			rs.Scontext.MContext.Set(perf.VarKey_PerfLagencyEnabled, enabled)
		}),
		restServer.MountMiddlewares(server.GLOBAL_MIDDLEWARE, perf.PerfRecover, log3.LogWithZap, perf.PerfLagency),
		restServer.MountRouters(basePath,
			health.NewHealthRouterEntry().UseEchoRouter(),
			api.NewSwaggerRouterEntry().UseEchoRouter(),
			api.NewScalarRouterEntry().UseEchoRouter(),
		),
		restServer.MountConfig(nil, nil, nil),
	)
	return restServer
}

func (m *Mocker) Runner() *server.Runner {
	runner := server.NewRunner().SetServer(m.RestServer("/api"))
	return runner
}
