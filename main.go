package main

import (
	_ "github.com/wensboy/sss/api/docs"

	"github.com/wensboy/sss/test"
)

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
	test.MockLoadCommand("./test/spec/command.json")
	test.MockLoadEnv("./test/spec/.env")
	test.MockLoadConfig("./test/spec", "server.json")
	runner := test.MockRunner()
	runner.Run()
}
