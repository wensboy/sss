# LOG

sss 日志主要包含如下考虑:
- 服务细节: 配置加载过程, 服务挂组件挂载过程...
- 中间件: 请求链路...
- 关键逻辑: 指标...

> 服务细节

必要: 配置加载情况, 组件挂载情况, 服务相关参数和状态

可选: 除必要外的一切细节

设计: logger -> server context -> internal log

关键: 1) 日志形式 2) 调用管理

时机: 服务挂载logger后 ~ 服务启动前

> 中间件

必要: 请求, 链路细节

可选: 除必要外的一切细节

设计: server context -> middleware context -> log middleware

关键: 1) 日志形式 2) 日志字段值源

时机: 完整链路结束处理后

> 关键逻辑

必要: 服务指标

可选: 除必要外的一切细节

设计: server context -> handler -> service -> repo

关键: 1) 日志形式 2) 指标 3) 输出

时机: 指标中间件注册后 ~ 指标中间件返回前

**sss 日志当前实际实现思路:**

考虑日志的msg部分, 采用 text/template 结合自定义 template context 进行处理来高效构建. 例如:

```go
func main() {
    format := `{{.timestamp}} - {{.httpMethod}} {{.uri}} {{.statusCode}} {{.lagency}}`
    tContext := map[string]any{
        "timestamp": time.Now(),
        "httpMethod": "GET",
        "uri": "http://localhost:3000/api/ping",
        "statusCode": 200,
        "lagency": 123.4,
    }
    t, _ := template.New("example").Parse(format)
    t.Execute(os.Stdout, tContext)
}
```

所有日志均按照如下优先级适配:

1. template + template context: 适用于灵活控制的日志场景
2. structed logger api: 适用于结构化日志收集
3. raw logger api: 适用于调试和开发

