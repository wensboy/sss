package auth

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/wensboy/sss/model"
)

const (
	VarKey_AuthEnabled  = "auth::var::enabled"
	ToolKey_AuthSkipper = "auth::tool::skipper"

	VarKey_JwtSecret  = "auth.jwt::var::secret"
	VarKey_UserClaims = "auth.jwt::var::user_claims"
	ToolKey_JwtUtil   = "auth.jwt::tool::jwt_util"
)

type AuthSkipper func(c *echo.Context) bool

func AuthWithJwt(mc model.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			skipper, ok := mc.Get(ToolKey_AuthSkipper).(AuthSkipper)
			if !mc.MustGet(VarKey_AuthEnabled).(bool) || (ok && skipper(c)) {
				return next(c)
			}
			jwtUtil := mc.MustGet(ToolKey_JwtUtil).(JwtUtil)
			token := jwtUtil.Extract(c)
			if token == "" {
				return echo.NewHTTPError(401, "Missing Authorization header")
			}
			claims, err := jwtUtil.Decode(token, mc.MustGet(VarKey_JwtSecret).(string))
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid auth token")
			}
			c.Set(VarKey_UserClaims, claims.UserClaims)
			c.Set(ToolKey_JwtUtil, jwtUtil)
			return next(c)
		}
	}
}

func AuthWithSession(mc model.MiddlewareContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			skipper, ok := mc.Get(ToolKey_AuthSkipper).(AuthSkipper)
			if !mc.MustGet(VarKey_AuthEnabled).(bool) || (ok && skipper(c)) {
				return next(c)
			}
			return next(c)
		}
	}
}
