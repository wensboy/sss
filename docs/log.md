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

> 中间件

必要: 请求, 链路细节

可选: 除必要外的一切细节

设计: server context -> middleware context -> log middleware

关键: 1) 日志形式 2) 

> 关键逻辑