package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)
}
