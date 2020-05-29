package randstring

import (
	"github.com/holdno/snowFlakeByGo"
	"math/rand"
	"strconv"
	"time"
)

func Rand8Numstring() string {
	min:=10000000
	max:=99999999
	return strconv.Itoa(GetRandSize(min,max))
}
func Rand6NumString()string{
	 min:=100000
	 max:=999999
	return strconv.Itoa(GetRandSize(min,max))
}
func Rand1NumString()string{
	min:=1
	max:=9
	return strconv.Itoa(GetRandSize(min,max))
}
func  GetRandomString(nSize int) string {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(chars)
	value := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < nSize; i++ {
		value = append(value, bytes[r.Intn(len(bytes))])
	}
	return string(value)
}
func GetRandSize(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return  min + int(rand.Int63n(int64(max)-int64(min)+1))
}
func GetRandAccntPwd() string {
	str := GetRandomString(GetRandSize(7,21))
	if str[0]=='0'||str[0]=='1'||str[0]=='2'||str[0]=='3'||str[0]=='4'||
		str[0]=='5'||str[0]=='6'||str[0]=='7'||str[0]=='8'||str[0]=='9'{
		str = str[1:len(str)]
	}
	str = str[0:len(str)-1]
	return str
}
var idWorker *snowFlakeByGo.Worker
func GetUuid() int64 {
	idWorker, _ = snowFlakeByGo.NewWorker(0)
	return  idWorker.GetId()
}
func GetUuidStr() string {
	return strconv.FormatInt(GetUuid(),10)
}
//example:
//rs := rand6.RandSix{}
//fmt.Println(rs.Rand6NumString())