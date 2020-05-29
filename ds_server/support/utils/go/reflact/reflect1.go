package main

import (
	"fmt"
	"reflect"
)

func reflectTest1(b interface{})  {
	rType := reflect.TypeOf(b)
	fmt.Println("rType =",rType)

	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal = %v, rVal的type = %T\n",rVal,rVal)

	fmt.Printf("kind=%v  kind=%v\n",rVal.Kind(), rType.Kind())

	iv := rVal.Interface()
	fmt.Printf("iv = %v, iv = %T\n",iv,iv)

}
type Student1 struct {
	Name string
	Age int
}
func  main() {
	//var num int = 100
	//fmt.Printf("num=%v---%T\n",num,num)
	//reflectTest01(num)

	stu := Student1{
		Name:"wang",
		Age:45,
	}

	reflectTest1(stu)
}
