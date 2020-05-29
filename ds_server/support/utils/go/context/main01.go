package main

import (
	"context"
)

func get_result(ctx context.Context) <-chan int  {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	result := get_result(ctx)

	for x := range result{

		if x == 1000000{
			break
		}
	}
}
