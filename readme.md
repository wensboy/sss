# SSS

so simple services.

| package | description |
|:-------:|:-----------:|
| [auth](./docs/auth.md) | middleware |
| [rbac](./docs/rbac.md) | middleware, handler, service, model |
| health | router, handler, service |
| log | middleware |
| api | router, handler |

> 如何集成

一个包可能包含:
- routers
- [middlewares](./docs/middleware.md)
- handlers
- services
- models
- repos

```bash
go get -u github.com/wensboy/sss
```

导入指定包使用即可. 通常直接导入 router 和 middleware, 少数情况下需要导入 service, repo 等. 当前由于设计复杂度问题, router, service, repo层能够集成到不同的项目中. middleware 实现只能集成到 echo 框架的项目当中.

> 打开 API 参考文档

```bash
# 安装 swag cli
go install github.com/swaggo/swag/cmd/swag@latest
# 生成 openapi 文档(在包含 main.go 的目录下执行)
swag init [-o <target>]
# 启动开发测试
go run ./main.go
```

针对 **restful api**: 

```go
api.NewSwaggerRouterEntry().UseEchoRouter() // swagger ui
api.NewScalarRouterEntry().UseEchoRouter() // scalar ui
```

> 与 ss 的关系

ss 是封装层, 用于对常见的第三方库做封装. sss 依赖 ss 的 api 来实现服务, sss 为应用服务层, 包含服务相关的部分, 只关注调用和业务逻辑, 不关注如何封装使用三方库.

注意: 由于 ss 倾向于三方封装, 因此部分封装设计对某些三方库存在强依赖, 例如: rest 处理复用器采用 Echo, db orm 使用的 xorm. 由于 sss 自身与 ss 解耦实现, 因此只有相关部分强依赖 ss, 其余部分不依赖 ss. sss 仓库自身可直接魔改使用.