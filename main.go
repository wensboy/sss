package main

import (
	"embed"

	_ "github.com/wensboy/sss/api/docs"

	"github.com/wensboy/sss/test"
)

//go:embed test/spec/*
var specFiles embed.FS

// @title SSS API Definition
// @version 1.0
// @description sss is a simple pre-built service repository that contains endpoint definitions for commonly integrable services.

// @contact.name API support and feedback
// @contact.url -
// @contact.email 2195058149@qq.com

// @license.name MIT
// @license.url https://github.com/wensboy/sss/blob/main/LICENSE

// @BasePath /api/v1
func main() {
	mocker := test.Mocker{}
	mocker.LoadCommand("./test/spec/command.json")
	mocker.LoadEnv("./test/spec/.env")
	mocker.LoadConfig("./test/spec", "server.json")
	runner := mocker.Runner()
	runner.Run()
}
