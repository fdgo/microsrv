package main

import (
	"fmt"
)

func main()  {
	Cache,_ :=  NewCache("120.27.239.127:6379")
	m := make(map[string]interface{})
	m["name"] = "wang"
	m["age"] = "33"
	if err :=Cache.SetFields("W",m);err!=nil{
		fmt.Println(err)
	}
	fmt.Println("------------------")
	x,y :=Cache.GetAllFields("W")
	fmt.Println(x)
	fmt.Println(y)
}
