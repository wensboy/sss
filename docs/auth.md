# AUTH

## auth type

- [cookie](./auth/cookie.md)
- [sesion](./auth/session.md)
- [jwt](./auth/jwt.md)
- [cas](./auth/cas.md)
- [oauth2](./auth/oauth2.md)

## auth middleware

### auth.jwt

| 上下文 | 变量 | 值类型 | 示例 | 说明 |
|:-----:|:-----:|:-----:|:-----:|:-----:|
| auth.jwt::var::secret | VarKey_JwtSecret | string | random uuid | jwt 密钥 |
| auth.jwt::var::user_claims | VarKey_JwtUserClaims | struct | UserClaims{1, "username"} | 用户负载 |
| auth.jwt::tool::jwt_util | ToolKey_JwtUtil | interface | jwtUtil | jwt 工具, 包含编码token, 解码token, 提取token |

### auth.session

| 上下文 | 变量 | 值类型 | 示例 | 说明 |
|:-----:|:-----:|:-----:|:-----:|:-----:|
| auth.session::tool::store | ToolKey_SessionStore | interface | random store | session store |
| auth.session::var::user_claims | VarKey_SessionUserClaims | struct | UserClaims{1, "username"} | 用户负载 |