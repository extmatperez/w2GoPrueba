package main

import (
	"context"
	"fmt"
	"time"
)

func saludar(ctx context.Context) {
	fmt.Println(ctx.Value("Saludo"))
}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "Saludo", "Hola es un saludo en context")
	//deadline := time.Now().Add(1 * time.Second)
	//ctx, _ = context.WithDeadline(ctx, deadline)

	ctx, _ = context.WithTimeout(ctx, time.Second*5)
	<-ctx.Done()

	fmt.Println(ctx.Err().Error())
}
