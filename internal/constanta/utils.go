package constanta

type ContextKey string

const (
	Tx           ContextKey = "tx"
	AuthUserID   ContextKey = "user_id"
	AuthRoleID   ContextKey = "role_id"
	AuthRoleName ContextKey = "role_name"
	AuthRoleCode ContextKey = "role_code"
	IsAdmin      ContextKey = "is_admin"
	Scope        ContextKey = "scope"
	TraceID      ContextKey = "trace_id"
	RequestID    ContextKey = "request_id"
)
