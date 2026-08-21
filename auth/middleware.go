package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v5"
	"github.com/wensboy/ss/server"
)

const (
	VarKey_AuthEnabled  = "auth::var::enabled"
	ToolKey_AuthSkipper = "auth::tool::skipper"
)

type AuthSkipper func(c *echo.Context) bool

func AuthWithJwt(mc server.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			skipper, ok := mc.Get(ToolKey_AuthSkipper).(AuthSkipper)
			if !mc.MustGet(VarKey_AuthEnabled).(bool) || (ok && skipper(c)) {
				return next(c)
			}
			jwtUtil := mc.MustGet(ToolKey_JwtUtil).(JwtUtil)
			token := jwtUtil.Extract(c)
			if token == "" {
				return echo.NewHTTPError(401, "missing header Authorization")
			}
			claims, err := jwtUtil.Decode(token, mc.MustGet(VarKey_JwtSecret).(string))
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid auth token")
			}
			c.Set(VarKey_JwtUserClaims, claims.UserClaims)
			c.Set(ToolKey_JwtUtil, jwtUtil)
			return next(c)
		}
	}
}

func AuthWithSession(mc server.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			skipper, ok := mc.Get(ToolKey_AuthSkipper).(AuthSkipper)
			if !mc.MustGet(VarKey_AuthEnabled).(bool) || (ok && skipper(c)) {
				return next(c)
			}
			store := mc.MustGet(ToolKey_SessionStore).(sessions.Store)
			session, err := store.Get(c.Request(), "sessionId")
			if err != nil || session.IsNew {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
			}
			userClaims, ok := session.Values[VarKey_SessionUserClaims]
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid empty claim session")
			}
			c.Set(VarKey_SessionUserClaims, userClaims)
			c.Set(ToolKey_SessionStore, store)
			return next(c)
		}
	}
}
