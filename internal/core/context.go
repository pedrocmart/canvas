package core

import (
	"context"
	"time"
)

func GetCtxStringVal(ctx context.Context, key ContextKey) string {
	ctxValue := ctx.Value(key)

	if ctxValue != nil {
		val, ok := ctxValue.(string)
		if ok {
			return val
		}
	}

	return ""
}

func GetCtxTimeVal(ctx context.Context, key ContextKey) time.Time {
	ctxValue := ctx.Value(key)

	if ctxValue != nil {
		val, ok := ctxValue.(time.Time)
		if ok {
			return val
		}
	}

	return time.Time{}
}
