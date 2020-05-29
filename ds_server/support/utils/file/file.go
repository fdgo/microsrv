package file

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"strings"
	"path/filepath"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func WriteFile(path string)  {
	f,err := os.Create(path)
	if err != nil{
		fmt.Println("err = ",err)
		return
	}
	defer f.Close()
	var buf string
	for i:=0;i<10;i++{
		buf = fmt.Sprintf("i = %d\n",i)
		_,err := f.WriteString(buf)
		if err != nil{
			fmt.Println("err = ",err)
		}
	}
}
func ReadFile(path string)  {
	f,err := os.Open(path)
	if err !=nil{
		fmt.Println("err = ",err)
		return
	}
	defer f.Close()
	buf := make([]byte,1024*4)
	n,err1:= f.Read(buf)
	if err1 != nil && err1 != io.EOF{
		fmt.Println("err1 = ",err1)
		return
	}
	fmt.Println(string(buf[:n]))
}
func ReadFileLine(path string)  {
	f,err := os.Open(path)
	if err !=nil{
		fmt.Println("err = ",err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for{
		//遇到'\n'结束读取，但是'\n'也读取进来
		buf,err := r.ReadBytes('\n')
		if err !=nil {
			if err == io.EOF{
				break
			}
			fmt.Println("err =",err)
		}
		fmt.Println(string(buf))
	}

}