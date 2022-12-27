package hoge

import "context"

// 独自構造体型の値をキーとして用いる
type traceIDKey struct{}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})
	
	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}