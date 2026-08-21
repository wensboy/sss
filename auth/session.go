package auth

const (
	ToolKey_SessionStore     = "auth.session::tool::store"
	VarKey_SessionUserClaims = "auth.session::var::user_claims"

	Store_Cookie = "cookie"
	Store_Mem    = "mem"
	Store_redis  = "redis"
	Store_sqlite = "sqlite"
)
