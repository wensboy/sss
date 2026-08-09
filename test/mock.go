package test

import (
	"github.com/joho/godotenv"
	"github.com/wensboy/ss/config"
	"github.com/wensboy/ss/server"
	"github.com/wensboy/sss/api"
	"github.com/wensboy/sss/health"
)

func MockLoadCommand(path string) {
	_ = config.InitCommand(path)
}

func MockLoadEnv(paths ...string) {
	err := godotenv.Load(paths...)
	if err != nil {
		panic(err)
	}
}

func MockLoadConfig(dir string, paths ...string) {
	cfg := config.NewConfig().SetDir(dir)
	config.SetGlobalConfig(cfg)
	cfg.Load(paths...)
}

func MockRestServer() *server.RestServer {
	restServer := server.NewRestServer()
	mockMountContext := func(s *server.RestServer) {
		s.Scontext.MountLogger()
	}
	restServer.MountModules(
		mockMountContext,
		restServer.MountRouters("/api/v1",
			health.NewHealthRouterEntry().UseEchoRouter(),
			api.NewSwaggerRouterEntry().UseEchoRouter(),
			api.NewScalarRouterEntry().UseEchoRouter(),
		),
		restServer.MountConfig(nil, nil, nil),
	)
	return restServer
}

func MockRunner() *server.Runner {
	runner := server.NewRunner().SetServer(MockRestServer())
	return runner
}
