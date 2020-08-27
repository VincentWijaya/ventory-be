package contextHelper

import (
	"context"
)

type contextKey int

const (
	UserIDKey contextKey = iota
	SessionCTKey
)

func (ck contextKey) getInt64(ctx context.Context) int64 {
	var (
		result int64
		ok     bool
	)
	if ctx.Value(ck) != nil {
		result, ok = ctx.Value(ck).(int64)
		if !ok {
			return 0
		}
	}
	return result
}

func UserIDFromContext(ctx context.Context) int64 {
	return UserIDKey.getInt64(ctx)
}

func GetSessionFromContext(ctx context.Context) (string, bool) {
	session, ok := ctx.Value(SessionCTKey).(string)
	return session, ok
}
