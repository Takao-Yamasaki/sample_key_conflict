package main

import (
	"context"
	"fmt"
	"sample_key_conflict/fuga"
	"sample_key_conflict/hoge"
)

func main() {
	// 空のコンテキストを作成
	ctx := context.Background()
	
	// トレースIDの付加
	ctx = hoge.SetTraceID(ctx, 1)
	ctx = fuga.SetTraceID(ctx, 2)
	
	hogeTraceID := hoge.GetTraceID(ctx)
	fugaTraceID := fuga.GetTraceID(ctx)

	fmt.Println(hogeTraceID, fugaTraceID)
}