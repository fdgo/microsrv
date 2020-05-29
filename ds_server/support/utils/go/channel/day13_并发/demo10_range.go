package main

import "fmt"

func main() {
	/*
	for range：
		数组，切片，map，string，chan
			数组/切片/string--->index，value
			map--->key, value
			chan--->value
	 */
	 ch1 :=make(chan string)
	 go sendData(ch1)


	 for value:= range ch1{//停止条件：通道关闭，显示的调用close()
	 	fmt.Println("从通道中读取的数据：",value)
	 }
}

func sendData(ch1 chan  string){
	for i:=1;i<=10;i++{
		ch1 <- fmt.Sprint("数据",i)
	}
	fmt.Println("写入数据完毕。。")
	close(ch1)
}
/*
练习：使用关闭通道结合for range从通道中读取数据
	一个goroutine写数据，两个goroutine读数据。。
 */
