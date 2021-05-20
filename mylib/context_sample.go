package mylib

import (
	"context"
	"fmt"
)

func ContextSample() {
	type favContextKey string // 定义类型favContextKey

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	fmt.Printf("k: %s\n", k)
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)

	k = favContextKey("color")
	//ctx = context.WithValue(context.Background(), k, "Red")
	f(ctx, k)
}

func ContextSample2() {
	context2 := context.Background()
	deadline, ok := context2.Deadline()
	if ok {
		panic(ok)
	}
	fmt.Printf("time: %s\n", deadline)

	ctx, cancel := context.WithCancel(context2)
	fmt.Printf("ctx: %s\n", ctx)
	if cancel != nil {
		fmt.Println("cancel is not nil")
		cancel() // 停止工作
	}

	ctx2, cancel2 := context.WithCancel(context2)
	fmt.Printf("ctx2: %s\n", ctx2)
	if cancel2 != nil {
		fmt.Println("cancel2 is not nil")
		cancel2() // 停止工作
	}
}
