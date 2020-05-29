package timex

import (
	"fmt"
	"strings"
	"time"
)

//获取系统时间(字符串形式)
func GetCurrentTime() (strCurTime string) {
	return time.Now().Format("2006-01-02 15:04:05")
}

//获取系统时间戳
func GetCurrentTimeStamp() (nCurrTimeStamp int64) {
	return time.Now().Unix()
}

//时间戳转时间
func TimeStampToTimeStr(nTimeStamp int64) (strTime string) {
	tm := time.Unix(nTimeStamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//时间转时间戳
func TimeStrToTimeStamp(strTime string) (nTimeStamp int64) {
	timestamp, _ := time.Parse("2006-01-02 15:04:05", strTime)
	return timestamp.Unix() - 3600*8
}

//两个时间的差值
func DistanceTwoTimes(strTimeFisrt, strTimeSecond string) int64 {
	return TimeStrToTimeStamp(strTimeFisrt) - TimeStrToTimeStamp(strTimeSecond)
}

//----------------------------------------------------------------------------------------------
type TimeEx struct {
	Year   int
	Month  time.Month
	Day    int
	Hour   int
	Minute int
	Second int
}

func (tx *TimeEx) GetBJTime() time.Time {
	// 将服务器UTC转成北京时间
	uTime := time.Now().UTC()
	dur, _ := time.ParseDuration("+8h")
	return uTime.Add(dur)
}
func (tx *TimeEx) TimeToTimeStr(tTime time.Time) (strTime string) {
	return time.Date(tTime.Year()+tx.Year, tTime.Month()+tx.Month, tTime.Day()+tx.Day, tTime.Hour()+tx.Hour,
		tTime.Minute()+tx.Minute, tTime.Second()+tx.Second, 0, tTime.Location()).Format("2006-01-02 15:04:05")
}

//example:
//func Time_Demo() {
//	t := &timepkg.Time{}
//	fmt.Println(t.GetCurrentTime())                                                      //当前时间(字符串形式)
//	timestamp := t.GetCurrentTimeStamp() - 3600                                          //当前时间的前一小时(时间戳形式)
//	before := t.TimeStampToTimeStr(timestamp)                                            //时间戳转化为字符串形式
//	fmt.Println(before)                                                                  //一小时前时间(字符串形式)
//	fmt.Println(t.DistanceTwoTimes(t.GetCurrentTime(), t.TimeStampToTimeStr(timestamp))) //两个时间的差值
//  tx := &TimeEx{Hour:1}                                                                //一小时后
//  fmt.Println(tx.TimeToTimeStr(tx.GetBJTime()))                                        //一小时后时间(字符串形式)
//}

func GetTimeStamp() (string, string) {
	timenow := time.Now()
	ts := timenow.Unix()

	format := "20060102150405.000"
	ss := timenow.Format(format)
	ss = strings.Replace(ss, ".", "", -1)

	return fmt.Sprintf("%v", ts), ss
}
