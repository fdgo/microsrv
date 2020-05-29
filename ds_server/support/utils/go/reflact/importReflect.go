package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int     `json:"monster_age"`
	Score float32
	Gender string
}

func (m Monster)Print()  {
	fmt.Println("---start---")
	fmt.Println(m)
	fmt.Println("---end---")
}
func (m Monster)GetSum(n1,n2 int) (int ,int) {
	return n1 + n2 ,100
}
func (m Monster)Set(name string, age int,score float32,gender string)  {
	m.Name = name
	m.Age = age
	m.Score = score
	m.Gender = gender
}
func TestStruct(a interface{})  {
	ty := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct{
		fmt.Println("expect struct")
		return
	}
	num := val.NumField()
	fmt.Printf("struct has %d fields\n",num)
	for i := 0; i< num ; i++ {
		fmt.Printf("Field %d:值=%v\n",i,val.Field(i))
		tagVal :=ty.Field(i).Tag.Get("json")
		if tagVal != ""{
			fmt.Printf("Field %d:tag=%v\n",i,tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n",numOfMethod)
	val.Method(1).Call(nil)


	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(40))
	res := val.Method(0).Call(params)
	fmt.Println("res = ",res[0].Int()) //res[1].Int() =100
}

func main()  {
	var a Monster = Monster{
		Name:"jack",
		Age:90,
		Score:45.8,
	}
	TestStruct(a)
}
