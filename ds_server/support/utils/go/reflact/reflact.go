package main

import (
	"fmt"
	"reflect"
)

func reflectTest0(b interface{})  {
	rType := reflect.TypeOf(b)
	fmt.Println("rType =",rType)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, rVal的type = %T\n",rVal,rVal)

	iv := rVal.Interface()
	fmt.Printf("iv = %v, iv = %T\n",iv,iv)


	x := iv.(Student0) // x := b.(Student)
	fmt.Println(x.Name)
}
type Student0 struct {
	Name string
	Age int
}
func  main() {
	//var num int = 100
	//fmt.Printf("num=%v---%T\n",num,num)
	//reflectTest01(num)

	stu := Student0{
		Name:"wang",
		Age:45,
	}

	reflectTest0(stu)
}
