package interfacee
import (
	"fmt"
	"reflect"
	"strconv"
)
type User struct {
	Id   int
	Name string
	Age  int
}
func (u User) Hello() {
	fmt.Println("Hello world.")
}
func main_test() {
	u := User{1, "ok", 12}
	Info(u) //如果是&u结果大不一样
}
func Info(o interface{}) {
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
//-----------------------------------------------------

type List []interface{}

type Person struct{
	name string
	age  int
}

func (p Person)String()string{
	return "(name:"+p.name+" - age:"+strconv.Itoa(p.age)+")"
}

func main_test2(){
	list := make(List,4)
	list[0]=23
	list[1]="China"
	list[2]=&Person{"Tom",35}
	list[3]=Person{"Jack",18}
	for index,element:=range list{
		switch element.(type){
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n",index,element.(int))
		case string:
			fmt.Printf("list[%d] is an string and its value is %s\n",index,element.(string))
		case *Person:
			element.(*Person).name = "Fred"
			fmt.Printf("list[%d] is an Person and its value is %v\n",index,element.(*Person))
		case Person:
			fmt.Printf("list[%d] is an Person and its value is %v\n",index,element.(Person))
		default:
			fmt.Printf("list[%d] is of a different type",index)
		}
	}
	fmt.Println(list)
	//fmt.Println(list[3])
}