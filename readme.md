# SSS

so simple services.

| package | description |
|:-------:|:-----------:|
| [auth](./docs/auth.md) | middleware |
| [rbac](./docs/rbac.md) | middleware, handler, service, model |
| health | router, handler, service |
| log | middleware |
| api | router, handler |

如何集成?

一个包可能包含:
- routers
- middlewares
- handlers
- services
- models
- repos

导入指定包使用即可. 通常直接导入 router 和 middleware, 少数情况下需要导入 service, repo 等.

```bash
go get -u github.com/wensboy/sss
```
