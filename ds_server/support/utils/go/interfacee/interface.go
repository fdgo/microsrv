package interfacee

import (
	"encoding/json"
	"fmt"
	"strconv"
)
func main() {
	txt := `{"a":1,"b":"2","c":[{"name":"1","group":"2"},{"name":"3","group":"4"}]}`
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(txt), &m); err != nil {
		panic(err)
	}
	//v := reflect.ValueOf(m["c"])
	//count := v.Len()
	//for i := 0; i < count; i++ {
	//	fmt.Println(v.Index(i))
	//}
	print_json(m)
}


func print_json(m map[string]interface{}) string {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			//fmt.Println(k, "is string", vv)
			return vv
		case float64:
			//fmt.Println(k, "is float", int64(vv))
			return strconv.FormatFloat(vv, 'E', -1, 64)
		case int:
			//fmt.Println(k, "is int", vv)
			return  strconv.Itoa(vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_json(vv)
		default:
			//fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
	return ""
}