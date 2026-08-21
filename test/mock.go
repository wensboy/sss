package test

import (
	"github.com/joho/godotenv"
	"github.com/wensboy/ss/config"
	"github.com/wensboy/ss/server"
	"github.com/wensboy/sss/api"
	"github.com/wensboy/sss/health"
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
		restServer.MountLogger(nil, nil, nil),
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
