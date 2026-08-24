package test

import (
	"fmt"
	"html/template"

	"github.com/joho/godotenv"
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
			tmpl, _ := template.New("mocker").Parse(format)
			rs.Scontext.MContext.Set(log3.VarKey_LogEnabled, enabled)
			rs.Scontext.MContext.Set(log3.VarKey_LogTemplate, tmpl)
			rs.Scontext.MContext.Set(log3.VarKey_LogTemplateCtx, log3.TemplateContext{})
			rs.Scontext.MContext.Set(log3.ToolKey_LogMutateLogger, rs.Scontext.GetLogger())
			enabled = config.MustLookup[bool](
				config.GConfigSource("server.rest.perf.recover"),
				config.DefaultSource(true),
			)
			rs.Scontext.MContext.Set(perf.VarKey_PerfRecoverEnabled, enabled)
			fmt.Printf("%+v\n", rs.Scontext.MContext.MustGet(perf.VarKey_PerfRecoverEnabled).(bool))
		}),
		restServer.MountMiddlewares(server.GLOBAL_MIDDLEWARE, perf.PerfRecover(restServer.Scontext.MContext), log3.LogWithZap(restServer.Scontext.MContext)),
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
