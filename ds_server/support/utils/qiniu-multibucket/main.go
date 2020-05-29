//package main
//
//import (
//	"log"
//	"github.com/go-test/qiniu-multibucket/global"
//)
//func init()  {
//	err := global.GetConfig()
//	if err != nil {
//		log.Printf("get globalconfig faild! err: #%v", err)
//	}
//	//mysql.InitMysql(global.GloConfig.Sqlconfig.MysqlConfig.User, global.GloConfig.Sqlconfig.MysqlConfig.Password, global.GloConfig.Sqlconfig.MysqlConfig.Server, global.GloConfig.Sqlconfig.MysqlConfig.Db)
//}
//func main()  {
//	UploadFile(global.GloConfig.Sqlconfig.QiniuossConfig.AccessKey,
//		global.GloConfig.Sqlconfig.QiniuossConfig.SecretKey,
//		global.GloConfig.Sqlconfig.QiniuossConfig.Bucket,"E:/update/win7_32/100100/win7_32_100.txt","win7_32/100/win7_32_100.txt")
//}

package main

import (
	"log"
	"os"
	"net/http"
	"io"
	"github.com/go-test/qiniu-multibucket/global"
)

func init() {
	err := global.GetConfig()
	if err != nil {
		log.Printf("get globalconfig faild! err: #%v", err)
	}
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//EnableProcessPrivileges([]string{"qiniu_multibucket.exe"})
	//UploadFile(global.GloConfig.Sqlconfig.QiniuossConfig.AccessKey,
	//	global.GloConfig.Sqlconfig.QiniuossConfig.SecretKey,
	//	global.GloConfig.Sqlconfig.QiniuossConfig.Bucket, "E:/update/win7_32/100100/Postman.zip", "win7_32/100100/Postman.zip")
	//
	//mac := qbox.NewMac("nym-Mch3cG-ndSeGCsrdqG6qe4PWFHbFXv3LWNTo", "p5UqJCNiZoXUyBtMvyVIp0Bz8T9iVXVpl5-u2jev")
	//domain := "http://ogbpegup0.bkt.clouddn.com"
	//key := "win7_32/100100/Postman.zip"
	//deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	//privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	//fmt.Println(privateAccessURL)

	//filename := "D:/win7_32_100100.txt"
	DownLoad("http-download://ogbpegup0.bkt.clouddn.com/win7_32/100100/Postman.zip?e=1534675185&token=nym-Mch3cG-ndSeGCsrdqG6qe4PWFHbFXv3LWNTo:oThec5qRFLmc4Zm78nadgylSprI=")
}
func DownLoad(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("F:/workspace/src/github.com/go-test/qiniu-multibucket/Postman.zip")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)

}

