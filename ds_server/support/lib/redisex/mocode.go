package redisex

//import (
//	"encoding/json"
//	"fmt"
//	"strconv"
//	"time"
//
//	"github.com/go-redis/redis"
//)
//func err_handler(err error) {
//	fmt.Printf("err_handler, error:%s\n", err.Error())
//	panic(err.Error())
//}
//
//func client() *redis.Client {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "115.207.209.117:6379",
//		Password: "",
//		DB:       0,
//	})
//	_, err := client.Ping().Result()
//	if err != nil {
//		fmt.Printf("ping error[%s]\n", err.Error())
//		err_handler(err)
//	}
//	return client
//}
//func main() {
//	rcli := client()
//	mc := &MobileCodeTime{
//		Mobile:"13182140525",
//		Code:"999000",
//		Time:GetSysTimeStamp(),
//	}
//	code,msg := CodeToRedisLimit(rcli, mc,3,25)
//	fmt.Println(code,"------",msg)
//	rcli.Close()
//}
//type MobileCodeTime struct {
//	Mobile string
//	Code string
//	Time int64
//}
//func GetSysTimeStamp() (nSysTimeStamp int64) {
//	return time.Now().Unix()
//}
//func MobCodTimeToJson(mbcode *MobileCodeTime) string {
//	ret,err := json.Marshal(mbcode)
//	if err!=nil{
//		fmt.Println("err = ",err)
//		return "nil"
//	}
//	return string(ret)
//}
//func JsonToMobCodTime(strsrc string )(mct *MobileCodeTime){
//	mct = new(MobileCodeTime)
//	err := json.Unmarshal([]byte(strsrc),mct)
//	if err!=nil{
//		fmt.Println("err = ",err)
//		return nil
//	}
//	return
//}
//func CodeToRedisLimit(rcli *redis.Client,mbcodetime *MobileCodeTime, ntimes int, ntimediff int64) (code string, msg string) {
//	defer func()  {
//		if p := recover(); p != nil {
//			return
//		}
//	}()
//	list_len, err := rcli.LLen(mbcodetime.Mobile).Result()
//	if err != nil {
//		err_handler(err)
//		return "nil","query existed code err! "+err.Error()
//	}
//	listlen ,_ := strconv.Atoi(strconv.FormatInt(list_len, 10))
//	if listlen < ntimes {
//		err :=rcli.LPush(mbcodetime.Mobile, MobCodTimeToJson(mbcodetime)).Err()
//		if err != nil {
//			err_handler(err)
//			return "nil","save code err! " + err.Error()
//		}
//		return GetCodeFromRedis(rcli,mbcodetime.Mobile),"ok"
//	} else {
//		mct, err := rcli.LIndex(mbcodetime.Mobile, 0).Result()
//		if err != nil {
//			err_handler(err)
//			return "nil","query latest code err! " + err.Error()
//		}
//		if GetSysTimeStamp()- JsonToMobCodTime(mct).Time < ntimediff {
//			return "nil", fmt.Sprintf("你已经发送了%d次，请一小时后再试...",ntimes)
//		}else {
//			lLen ,_ := strconv.Atoi(strconv.FormatInt(list_len, 10))
//			for n:=0; n< lLen;n++ {
//				err :=rcli.LPop(mbcodetime.Mobile).Err()
//				if err != nil {
//					err_handler(err)
//					return "nil","delete history code err! " + err.Error()
//				}
//			}
//			err :=rcli.LPush(mbcodetime.Mobile, MobCodTimeToJson(mbcodetime)).Err()
//			if err != nil {
//				err_handler(err)
//				return "nil","save code err! " + err.Error()
//			}
//			return GetCodeFromRedis(rcli,mbcodetime.Mobile),"ok"
//		}
//	}
//}
//func GetCodeFromRedis(rcli *redis.Client,strmobile string ) string {
//	mct, err := rcli.LIndex(strmobile, 0).Result()
//	if err != nil {
//		err_handler(err)
//	}
//	return JsonToMobCodTime(mct).Code
//}
