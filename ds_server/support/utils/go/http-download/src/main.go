package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)
func welcome01(w http.ResponseWriter,r *http.Request)  {
	h := r.Header
	fmt.Fprintln(w,h["Accept-Language"])

	w.Header().Set("Content-Type","text/html;charset=utf-8")
	fmt.Fprintln(w,"123<b>yes</b>")

	r.ParseForm()//先解析成form
	fmt.Fprintln(w,r.Form)
	fmt.Fprintln(w,r.FormValue("name"))
}
type Stu struct {
	Name string
	Age int
}
func welcome02(w http.ResponseWriter,r *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, Stu{Name:"wangminghui",Age:33})

	//    {{.Name}}
	//    {{.Age}}
}
func welcome03(w http.ResponseWriter,r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	m := make(map[string]interface{})
	m["user"] = Stu{Name:"wmh888",Age:89}
	m["money"] = 789
	t.Execute(w,m)

	//    {{.user.Name}}
	//    {{.user.Age}}
	//    {{.money}}
}
func welcome04(w http.ResponseWriter,r *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	time1 := time.Date(2018,1,2,3,4,5,0,time.Local)
	t.Execute(w, time1)
    //	  完整时间:{{.}}<br/>
	//    年:{{.Year}}<br/>
	//    月:{{.Month}}<br/>
	//    {{.Format "2006-01-02 15:04:05"}}
}

//-----------------------------------------------------
func myfun( t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func welcome05(w http.ResponseWriter,r *http.Request)  {
	mf := template.FuncMap{"mf":myfun}
	t:= template.New("index.html").Funcs(mf)
	t, _ = t.ParseFiles("view/index.html")
	time1 := time.Date(2018,1,2,3,4,5,0,time.Local)
	t.Execute(w, time1)
	//	  完整时间:{{.}}<br/>
	//    年:{{.Year}}<br/>
	//    月:{{.Month}}<br/>
	//    {{.Format "2006-01-02 15:04:05"}}<br/>
	//    自定义函数:{{mf .}}  点为参数
}
//-----------------------------------------------------
func mystring( str string) string {
	return str + " good!!!"
}
func welcome06(w http.ResponseWriter,r *http.Request)  {
	mf := template.FuncMap{"mstr":mystring}
	t:= template.New("index.html").Funcs(mf)
	t, _ = t.ParseFiles("view/index.html")
	str := "wangminghui"
	t.Execute(w, str)
    //自定义函数:{{mstr .}}<br/>
}
//--------------------------------------------------------------
func welcome07(w http.ResponseWriter,r *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, "11")
	//{{if .}}
	//    ififif
	//{{else}}
	//    elseelse
	//{{end}}
}

func welcome08(w http.ResponseWriter,r *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, "11")
	////gt大于的意思</br>
	//{{ $n:=678}}
	//{{if gt $n 456}}
	//    执行if
	//{{else}}
	//    执行else
	//{{end}}
}

func welcome09(w http.ResponseWriter,r *http.Request)  {
	t, _ := template.ParseFiles("view/index.html")
	//arr := []string{"first","second"}
	//t.Execute(w, arr)
	//{{range .}}
	//    {{.}}<br/>
	//{{end}}

	m := map[string]string{"key1":"value1","key2":"value2"}
	t.Execute(w,m)
	//{{range $k,$v := . }}
	//    {{ $k }} <br/>
	//    {{ $v }} <br/>
	//{{end}}
}


func welcome10(w http.ResponseWriter,r *http.Request)  {
	t, _ := template.ParseFiles("view/index.html","view/head.html","view/foot.html")
	t.ExecuteTemplate(w,"layout",nil)
}
func main() {
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))//让文件服务生效
	http.HandleFunc("/abc",welcome01)
	http.HandleFunc("/def",welcome02)
	http.HandleFunc("/ghi",welcome03)
	http.HandleFunc("/jkl",welcome04)
	http.HandleFunc("/mno",welcome05)
	http.HandleFunc("/pqr",welcome06)
	http.HandleFunc("/stu",welcome07)
	http.HandleFunc("/vwx",welcome08)
	http.HandleFunc("/yza",welcome09)
	http.HandleFunc("/bcd",welcome10)
	http.ListenAndServe(":8081",nil)
}
