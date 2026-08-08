# session

服务端 session 是较为安全的简单的浏览器状态实现.

## 原理

服务端存储状态session并通过Set-Cookie设置session_id到客户端.

## 实现

以 golang 为例, 伪代码:

```go
func SetSession(extractSessionContext func(*handleContext)) func(*HandleContext) {
    return func(c *HandleContext) {
        next(c)
        sessionContext := extractSessionContext(c)
        sessionId := NewSession(sessionContext)
        c.SetCookie([]string("sessionId", sessionId))
    }
}
```

相关参考:
- [gorilla/sessions](https://github.com/gorilla/sessions)

## 缺陷

- session 累积增加存储资源消耗
- 分布式难保证一致性

## 场景

- 安全
- 减轻网络负载
- 复杂会话状态记录