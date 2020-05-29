package interfacee

import (
	"fmt"
	"reflect"
)

type User6 struct {
	Id   int
	Name string
	Age  int
}

func (u User6) Hello() {
	fmt.Println("Hello world.")
}
func main_test6() {
	u := User6{1, "ok", 12}
	Info6(&u)
}
func Info6(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("XX")
		return
	}
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

}