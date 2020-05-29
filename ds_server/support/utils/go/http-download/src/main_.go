package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)
//----------------------------------------------------------------------
func hello01(w http.ResponseWriter,r *http.Request)  {
	t,_:= template.ParseFiles("view/file.html")
	t.Execute(w,nil)
}
func upload(w http.ResponseWriter,r *http.Request)  {
	fileName := r.FormValue("name")
	file,fileHeader,_:= r.FormFile("file")
	b,_ := ioutil.ReadAll(file)
	f := "D:/"+ fileName + fileHeader.Filename[strings.LastIndex(	fileHeader.Filename,"."):]
	err := ioutil.WriteFile(f,b,0777)
	fmt.Println(err)
	t,_:=template.ParseFiles("view/success.html")
	t.Execute(w,nil)
}
//----------------------------------------------------------------------
func hello02(w http.ResponseWriter,r *http.Request)  {
	t,_:= template.ParseFiles("view/downloadfile.html")
	t.Execute(w,nil)
}
func download(w http.ResponseWriter,r *http.Request)  {
	filename := r.FormValue("filename")
	f,err := ioutil.ReadFile("/opt/workspace/src/microservice.tar.gz")//+filename
	if err !=nil{
		fmt.Fprintln(w,"文件下载失败!",err)
		return
	}
	h := w.Header()
	h.Set("Content-Type","application/octet-stream")
	h.Set("Content-Disposition","attachment;filename="+filename)
	w.Write(f)
}


func main() {
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))//让文件服务生效
	http.HandleFunc("/aaa",hello01)
	http.HandleFunc("/upload",upload)



	http.HandleFunc("/bbb",hello02)
	http.HandleFunc("/download",download)
	http.ListenAndServe(":8081",nil)
}
