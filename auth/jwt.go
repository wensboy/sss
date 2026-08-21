package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

const (
	VarKey_JwtSecret     = "auth.jwt::var::secret"
	VarKey_JwtUserClaims = "auth.jwt::var::user_claims"
	ToolKey_JwtUtil      = "auth.jwt::tool::jwt_util"
)

type UserClaims struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type JwtClaims struct {
	UserClaims
	jwt.RegisteredClaims
}

type JwtUtil interface {
	Encode(*JwtClaims, jwt.SigningMethod, string) (string, error)
	Decode(string, string) (*JwtClaims, error)
	Extract(*echo.Context) string
}

type jwtUtil struct{}

func NewJwtUtil() JwtUtil {
	return jwtUtil{}
}

func (j jwtUtil) Encode(claims *JwtClaims, method jwt.SigningMethod, secret string) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString([]byte(secret))
}

func (j jwtUtil) Decode(token string, secret string) (*JwtClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*JwtClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func (j jwtUtil) Extract(c *echo.Context) string {
	token := c.Request().Header.Get("Authorization")
	if token != "" {
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
	}
	return token
}
