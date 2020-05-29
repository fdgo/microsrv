package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	/*
	遍历文件夹
	 */
	 dirName := "/home/ruby/文档/pro"
	 //fileInfos,_:=ioutil.ReadDir(dirName)
	 ////fmt.Println(len(fileInfos))
	 //for i,fi:=range fileInfos{
	 //	//fileName := dirName+"/"+fi.Name()
	 //	fileName := path.Join(dirName,fi.Name())
	 //	fmt.Println(i,fileName)
	 //	if fi.IsDir(){
	 //		//fmt.Println("\t",fileName)
	 //		fileInfos2,_:=ioutil.ReadDir(fileName)
	 //		for j,fi2:=range fileInfos2{
	 //			fmt.Println("\t",j,fi2.Name())
		//	}
		//}
	 //}
	 listDir(dirName)
}

func listDir(dirName string){
	fileInfos,_:=ioutil.ReadDir(dirName)
	for _,fi:=range fileInfos{
		fileName := dirName+"/"+fi.Name()
		fmt.Println(fileName)
		if fi.IsDir(){
			listDir(fileName)
		}
	}
}
