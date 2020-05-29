package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  int
}
func start(w http.ResponseWriter,r *http.Request)  {
	t, _ := template.ParseFiles("view/jsonshow.html")
	t.Execute(w, "11")
}
func show(w http.ResponseWriter,r *http.Request)  {
	us := make([]User,0)
	us = append(us,User{"lily",12})
	us = append(us,User{"tom",16})
	us = append(us,User{"jack",23})
	b,_:=json.Marshal(us)
	w.Header().Set("Content-Type","application/json;charset=utf-8")
	fmt.Fprintln(w,string(b))

}

func main()  {
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))//让文件服务生效
	http.HandleFunc("/",start)
	http.HandleFunc("/show",show)
	http.ListenAndServe(":8081",nil)
}
