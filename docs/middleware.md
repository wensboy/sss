# MIDDLEWARE

中间件本质上就是一个集成特殊功能的可嵌入流程函数. 在大多数 golang 生态的 web framework 中, 中间件通常通过一种洋葱模型形式集成到流程. 在其他语言中可能被成为拦截器. 中间件是链路关联的, 中间件的处理和当前链路的上下文相关, 同时也作为当前上下文的一部分. 为了避免中间件的形式多样, 应该尝试尽可能统一中间件形式. 当前中间件主要为 echo 框架实现.

在 ss 中, 一个注入的执行顺序为: `config -> server -> server context -> router -> handler -> service -> repo`. server context 通常包含所有需要的配置信息, 避免频繁调用获取全局变量, 通常以指针的形式引用唯一数据源. 中间件在 router 中被注册, 考虑 echo 中的中间件形式:

```go
func <middleware_name>(...) echo.MiddlewareFunc {
    return func(next echo.HandleFunc) echo.HandleFunc {
        return func(c *echo.Context) error {
            // 实际中间件逻辑
        }
    }
}
// group.Use(m1,m2,m3...).Any(..., h) -> fh := m1(m2(m3(h)))
// <middleware_name> 封装的本质是捕获上下文从而影响中间件行为
```

一个中间件至少需要:

- 相关的变量: 中间件处理
- 工具注入: 后续链路

**相关的变量**

1. 不同中间件变量上下文独立. 例如: 

```go
func <middleware_name>(<middleware_context>) echo.MiddlewareFunc {
    return func(next echo.HandleFunc) echo.HandleFunc {
        return func(c *echo.Context) error {
            // 实际中间件逻辑
        }
    }
}
```

2. 统一上下文(上下文是一种实现了set, get, has方法的map结构, 被方法接口化).

- kv 结构, 入参只读, 无并发安全问题
- key 规范为: `<domain>::<kind>::<id>`
- value 走断言类提取, 例如: `value.(T), cast.To[T](value)`


复用问题: server context 统一注册

**工具注入**

1. 注入函数. 例如:

```go
func <middleware_name>(<middleware_context>, <inject_tools_func>) echo.MiddlewareFunc {
    return func(next echo.HandleFunc) echo.HandleFunc {
        return func(c *echo.Context) error {
            <inject_tools_func>(c)
            // 实际中间件逻辑
        }
    }
}
```

注意: 1) 注入函数一定要在 next 执行前调用 2) 注入函数本质上是一种约定, key能够统一管理

2. 统一上下文(和相关变量一致, 工具也是kv结构).

sss 中中间件定义形式:

```go
func <middleware-name>(<middleware-context>) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c *echo.Context) error {
            // 中间件逻辑
        }
    }
}
```

由于形式一致, 通常会在相关的主索引文件中按照如下格式记录必要的参数:

```md
| 上下文 | 变量 | 值类型 | 示例 | 说明 |
|:-----:|:-----:|:-----:|:-----:|:-----:|
| <context-value> | <varible> | <value-type> | <eg> | <description> |
```

例如:
| 上下文 | 变量 | 值类型 | 示例 | 说明 |
|:-----:|:-----:|:-----:|:-----:|:-----:|
| auth.jwt::tool::jwt_util | ToolKey_JwtUtil | interface | JwtUtil | jwt工具, 包含生成token, 解压负载, 提取token |