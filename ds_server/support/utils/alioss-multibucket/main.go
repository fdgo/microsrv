package main

import (
	"log"
	"github.com/go-test/alioss-multibucket/global"
	"fmt"
)

func init()  {
	err := global.GetConfig()
	if err != nil {
		log.Printf("get globalconfig faild! err: #%v", err)
	}
}
func main()  {
	ret := PutObjectFromFileWithUrl("hf-learning-test", "win7_32/100/win7_32_100.txt", "E:/x/win7_32/100/win7_32_100.txt")//阿里oss上传
	if !ret {
		fmt.Println("上传失败")
		return
	}
	ret = PutObjectFromFileWithUrl("hf-learning-test", "win7_32/200/win7_32_200.txt", "E:/x/win7_32/200/win7_32_200.txt")//阿里oss上传
	if !ret {
		fmt.Println("上传失败")
		return
	}
}
