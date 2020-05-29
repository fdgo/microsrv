package main

import (
	"fmt"
	"reflect"
)


func main()  {
	var num int = 23
	reflect3(&num)
	fmt.Println("***",num)

	//等价
	//var abc int = 666
	//fn := reflect.ValueOf(&abc)
	//fn.Elem().SetInt(888)
	//fmt.Println(abc)
}

func reflect3(b interface{})  {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(300)
}