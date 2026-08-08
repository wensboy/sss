# jwt

json web token 是目前实现浏览器状态最常见的形式.

## 原理

将必要的信息通过服务端加密形成三段式 token 设置到浏览器的 cookie 当中, 后续浏览器携带并在服务端提取.

## 实现

详细参考: [golang-jwt](https://github.com/golang-jwt/jwt)

## 缺陷

- 仍然存在风险
- 过期策略设置
- 某些场景引入过复杂

## 场景

- 安全
- 灵活控制
- 复杂校验体系设计