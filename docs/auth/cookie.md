# cookie

浏览器 cookie 是最简单的状态保持实现.

## 原理

服务端响应通过响应头: Set-Cookie 来设置客户端 cookie, 后续请求自动带上 cookie.

## 实现

以 golang 为例, 伪代码:

```go
func SetCookie(extractCookieContext func(*handleContext)) func(*handleContext) {
    return func(c *handleContext) {
        next(c)
        cookieCtx := extractCookieContext(c)
        if len(cookieCtx)>0 {
            c.SetCookie(cookieCtx)
        }
    }
}
// 注入
func injectCookieContext(c *handleContext, cookies []string)
// 提取
func extractCookieContext(c *handleContext) []string
```

## 缺陷

- 增加http网络负载
- 安全性差
- 大小限制

## 场景

- 含个性化元信息
- 测试环境下快速校验