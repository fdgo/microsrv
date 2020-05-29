package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

func main() {
	/*
	模拟聊天记录
	从键盘接收数据，模拟聊天记录，并将内容保存到本地文件中
	 */
	 fileName := "chat.txt"
	 //fmt.Println(filepath.Abs(fileName))
	 r1:=bufio.NewReader(os.Stdin)//创建一个bufio.Reader对象，读取键盘
	 file,_:=os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	 w1 :=bufio.NewWriter(file)//创建一个bufio.Writer对象，写出数据到chat.txt文件中

	 defer file.Close()


	 str:=""//表示读取到数据
	 w1.WriteString(time.Now().Format("2006-01-02"))
	 w1.WriteString("\n")
	 w1.Flush()//刷新缓冲区：将缓冲区中的数据，写入到目标文件中

	 name :=""
	 flag := true
	 content:=""
	 for{
	 	//1.设置name
	 	if flag{
	 		name = "小明"
		}else{
			name = "小红"
		}
		flag = !flag
	 	//2.读取键盘
	 	str,_=r1.ReadString('\n')
	 	if str == "over\n"{
	 		fmt.Println("程序即将推出。。")
	 		w1.WriteString(str)
	 		w1.Flush()
	 		break
		}
	 	//3.拼串
	 	content = name +":"+str//fmt.Sprint(name,":",str)
	 	fmt.Print(content)
	 	//4.将数据写入到文件中
	 	w1.WriteString(time.Now().Format("15:04:05"))
	 	w1.WriteString("\n")
	 	w1.WriteString(content)
	 	w1.Flush()
	 }
}
