package main

import (
	"context"
	"fmt"
)

func main()  {
	ctx := context.WithValue(context.Background(),"trace_id",12345633)
	ret,ok := ctx.Value("trace_id").(int)
	if !ok{
		ret = 34589
	}
	fmt.Printf("ret:%d\n",ret)

	ctx = context.WithValue(ctx,"user_id","fred")
	result,b := ctx.Value("user_id").(string)
	if !b {
		result = "hello world!"
	}

	fmt.Printf("result:%s\n",result)
}

