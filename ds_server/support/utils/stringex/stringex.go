package stringex

import (
	"ds_server/support/utils/bytesex"
	"fmt"
	"bytes"
	rd "crypto/rand"
	"encoding/binary"
	"encoding/json"
	"math/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"
	"github.com/holdno/snowFlakeByGo"
	"ds_server/support/utils/sign/md5"
)
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// string()
func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// []byte()
func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func Base64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Rand8Numstring() string {
	min:=10000000
	max:=99999999
	return strconv.Itoa(GetRandNum(min,max))
}
func Rand6NumString()string{
	min:=100000
	max:=999999
	return strconv.Itoa(GetRandNum(min,max))
}
func Rand1NumString()string{
	min:=1
	max:=9
	return strconv.Itoa(GetRandNum(min,max))
}
func  GetRandomString(nSize int) string {
	chars := "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(chars)
	value := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < nSize; i++ {
		value = append(value, bytes[r.Intn(len(bytes))])
	}
	return string(value)
}
func GetRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return  min + int(rand.Int63n(int64(max)-int64(min)+1))
}
func GetRandAccntPwd() string {
	str := GetRandomString(GetRandNum(7,21))
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



func StrJoin(tag string, str...string )string  {
	slice := make([]string,0)
	slice = append(slice,str...)
	return strings.Join(slice,tag)
}
//正向截取，得到全部
func StrSplit(tag string, str string)(int,[]string)  {
	slice := make([]string,0)
	stringArr := strings.Split(str,tag)
	for _,v := range stringArr{
		slice = append(slice,v)
	}
	return len(slice),slice
}
//反向截取一次
func StrRevSplit(tag string, str string) (string,string)  {
	index := strings.LastIndex(str,tag)
	if index == -1{
		return "NULL","NULL"
	}
	front := str[:index]
	back := str[index+1:]
	return front,back
}
//截取某一段
func StrSub(str string, pos, length int) string {
	runes := []rune(str)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
//字符串逆序输出
func StrRev(s string) string {
	str := []rune(s)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
//删除空格、换行、空格等字符
func StrStrim(s string) string {
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	return s
}

//ret := StrJoin("@","abc","def","mnp","234")
//fmt.Println(ret) //abc@def@mnp@234
//---------------------------------------------------
//accout, slice := StrSplit("@",ret)
//fmt.Println(accout,slice) // 4 [abc def mnp 234]
//---------------------------------------------------
//src := "/opt/hello/56.html"
//front,back := StrRevSplit(".",src)
//fmt.Println(front,back) //  /opt/hello/56  html
//-----------------------------------------------------
//src := "/opt/hello/56.html"
//ret := StrSub(src,2,4)
//fmt.Println(ret)//  pt/h
//-----------------------------------------------------
//	fmt.Println(StrRev("abc123")) //321cba
//-----------------------------------------------------
//fmt.Println(StrStrim("abc 345  xyz"))  //abc345xyz

//使用 utf8.RuneCountInString()统计字符串长度
func Length(str string) int {
	return utf8.RuneCountInString(str)
}

//字符串转成Unicode编码
func String2Unicode(s string) string {
	json := ""
	for _, r := range s {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16)
		}
	}
	return json
}

//Unicode编码转成字符串
func Unicode2String(s string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(s, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

//html编码
func HTMLEncode(s string) string {
	html := ""
	for _, r := range s {
		html += "&#" + strconv.Itoa(int(r)) + ";"
	}
	return html
}

//获取一个Guid
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rd.Reader, b); err != nil {
		return ""
	}
	return  md5.Md5(Base64(b))
}
//把IP地址转成数字
func GetIPNums(s string) (ipNum uint32, err error) {
	if strings.EqualFold(s, "") {
		return ipNum, errors.New("ipAddress is null")
	}
	items := strings.Split(s, ".")
	if len(items) != 4 {
		return ipNum, errors.New("ipAddress is error")
	}
	item0, err := strconv.Atoi(items[0])
	if err != nil {
		return ipNum, errors.New("ipAddress is error0")
	}
	item1, err := strconv.Atoi(items[1])
	if err != nil {
		return ipNum, errors.New("ipAddress is error1")
	}
	item2, err := strconv.Atoi(items[2])
	if err != nil {
		return ipNum, errors.New("ipAddress is error2")
	}
	item3, err := strconv.Atoi(items[3])
	if err != nil {
		return ipNum, errors.New("ipAddress is error3")
	}
	return uint32(item0<<24 | item1<<16 | item2<<8 | item3), nil
}

//获取IP地址，不包含端口号
func GetIPAddressNotPort(ip_address string) string {
	if !strings.Contains(ip_address, ":") {
		return ip_address
	}
	start := strings.Index(ip_address, ":")
	if start <= 2 {
		return ip_address
	}
	return StrSub(ip_address, 0, start)
}

//判断是否是淘宝用户名
func IsTaobaoNick(s string) bool {
	if len(s) < 2 {
		return false
	}
	return regexp.MustCompile(`(^[\\u4e00-\\u9fa5\\w_—\\-，。…·〔〕（）！@￥%……&*？、；‘“]*$)`).MatchString(s)
}

//判断是否是淘宝用户名（子帐号）
func IsSubTaobaoNick(s string) bool {
	if len(s) < 2 {
		return false
	}
	return regexp.MustCompile(`(^[\\u4e00-\\u9fa5\\w_—\\-，。…·〔〕（）！@￥%……&*？、；‘“:]*$)`).MatchString(s)
}

//判断是否版本号
func IsVersion(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile(`(^[0-9.]*$)`).MatchString(s)
}

//判断是否网址
func IsUrl(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile(`(^[a-zA-z]+://[^\s]*$)`).MatchString(s)
}

//是否数字
func IsNumber(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile(`(^[0-9]*$)`).MatchString(s)
}

//是否多数字(用逗号间隔)
func IsMultipNumber(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile(`(^[0-9,]*$)`).MatchString(s)
}

//判断是否由字母、数字、下划线组成
func IsLetterOrNumber(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile(`(^[A-Za-z0-9_]*$)`).MatchString(s)
}

//判断是否由字母、数字、下划线组成
func IsLetterOrNumber1(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile(`(^[A-Za-z0-9_-]*$)`).MatchString(s)
}

//判断是否由汉字、字母、数字、下划线组成
func IsHanOrLetterOrNumber(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile("^[A-Za-z0-9_\u4e00-\u9fa5]*$").MatchString(s)
}

// 是否IPv4地址
func IsIPAddress(ip string) bool {
	matched, err := regexp.MatchString("(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)", ip)
	if err != nil {
		return false
	}
	return matched
}

// 是否内网IP地址
func IsIntranetIP(s string) bool {
	matched, err := regexp.MatchString(`^((192\.168|172\.([1][6-9]|[2]\d|3[01]))(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){2}|10(\.([2][0-4]\d|[2][5][0-5]|[01]?\d?\d)){3})$`, s)
	if err != nil {
		return false
	}
	return matched
}

//是否email
func IsEmail(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$").MatchString(s)
}

//是否手机号
func IsMobile(s string) bool {
	if len(s) < 1 {
		return false
	}
	return regexp.MustCompile("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$").MatchString(s)
}

/*
判断字符串是否全中文字符
*/
func IsAllChineseChar(s string) bool {
	for _, r := range s {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

//是否utf-8编码字符串
func IsUtf8(s string) bool {
	count := 0
	for _, v := range s {
		if int(v) > 65530 {
			count++
		}
	}
	return count == 0
}

const zip_offset int = 19968

//压缩md5或guid
func ZipMd5(md5String string) (zip_string string, err error) {
	if len(md5String) != 16 && len(md5String) != 32 {
		return "", errors.New("源md5值长度不对")
	}
	var md5_bytes []byte
	switch len(md5String) {
	case 16:
		md5_bytes = getHexBytes(md5String + "00")
	case 32:
		md5_bytes = getHexBytes(md5String + "0")
	}
	var data bytes.Buffer
	var total int = 0
	for index := 0; index < len(md5_bytes); index++ {
		switch index % 3 {
		case 0:
			intByte, err := strconv.Atoi(fmt.Sprintf("%d", md5_bytes[index]))
			if err != nil {
				return zip_string, err
			}
			total = intByte
			break
		case 1:
			intByte, err := strconv.Atoi(fmt.Sprintf("%d", md5_bytes[index]))
			if err != nil {
				return zip_string, err
			}
			total += intByte << 4
			break
		case 2:
			intByte, err := strconv.Atoi(fmt.Sprintf("%d", md5_bytes[index]))
			if err != nil {
				return zip_string, err
			}
			total += (intByte << 8) + zip_offset
			uni_string := fmt.Sprintf("\\u%x", total)
			chinese_string, err := Unicode2String(uni_string)
			if err != nil {
				return zip_string, err
			}
			data.WriteString(chinese_string)
			total = 0
			break
		}
	}
	zip_string = data.String()
	return
}

//解压缩汉字，结果为guid或md5值
func UnZipMd5(zip_string string) (md5_string string, err error) {
	if len(zip_string) != 18 && len(zip_string) != 33 {
		return "", errors.New("源zip值长度不对")
	}
	var data bytes.Buffer
	unicode_string := String2Unicode(zip_string)
	unicode_strings := strings.Split(unicode_string, "\\u")
	for _, unicode_alue := range unicode_strings {
		if strings.EqualFold(unicode_alue, "") {
			continue
		}
		dec, err := strconv.ParseInt(unicode_alue, 16, 32)
		if err != nil {
			return md5_string, err
		}
		dec -= int64(zip_offset)
		data.WriteString(ten_value_to_char(dec & 15))
		data.WriteString(ten_value_to_char((dec >> 4) & 15))
		data.WriteString(ten_value_to_char((dec >> 8) & 15))
	}
	switch len(zip_string) {
	case 18:
		md5_string = StrSub(data.String(), 0, 16)
	case 33:
		md5_string = StrSub(data.String(), 0, 32)
	}
	return
}

func getHexBytes(str string) []byte {
	result := []byte(str)
	for index := 0; index < len(result); index++ {
		if result[index] < 58 {
			result[index] -= 48
		} else {
			result[index] -= 55
		}
	}
	return result
}

func ten_value_to_char(ten int64) string {
	switch ten {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		return strconv.FormatInt(ten, 10)
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	}
	return ""
}

//获取当前内网IP
func GetCurrentIntranetIP() string {
	ip_address := ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("获取当前内网IP出错：", err)
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if strings.HasPrefix(ipnet.IP.String(), "10.") || strings.HasPrefix(ipnet.IP.String(), "192.") {
					ip_address = ipnet.IP.String()
					break
				}
			}
		}
	}
	if len(ip_address) < 9 {
		fmt.Println("获取当前内网IP出错：没有找到IP")
		return ""
	}
	return ip_address
}

//序列化为json
//func ToJson(data interface{}) string {
//	b, _ := json.Marshal(data)
//	return string(b)
//}

//格式化为Json字符串
func FormatJson(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		return err.Error()
	}
	return out.String()
}

//序列化-->zlib压缩-->混淆-->base64
func ToJsonZipConfusedBase64(obj interface{}) string {
	b, _ := json.Marshal(obj)
	b, _ = bytesex.ZlibZipBytes(b)
	b = bytesex.ConfusedTwo(b)
	return Base64(b)
}

//序列化-->混淆-->base64
func ToJsonConfusedBase64(obj interface{}) string {
	b, _ := json.Marshal(obj)
	b = bytesex.ConfusedTwo(b)
	return Base64(b)
}

//混淆-->zlib压缩-->base64
func ToConfusedZipBase64(str string) string {
	b := bytesex.ConfusedTwo([]byte(str))
	b, _ = bytesex.ZlibZipBytes(b)
	return Base64(b)
}

//混淆-->base64
func ToConfusedBase64(str string) string {
	b := bytesex.ConfusedTwo([]byte(str))
	return Base64(b)
}

//生成简单随机密码，短时间内会重复
func GetSimplePwd(lenth int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	runes := []rune(chars)
	len := len(runes)
	time.Sleep(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	pwd := NewStringBuilder()
	for index := 0; index < lenth; index++ {
		pwd.Append(string(runes[rand.Intn(len)]))
	}
	return pwd.ToString()
}

//生成随机密码，短时间内会重复
func GetPwd(lenth int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&*+-=?@^_~'"
	runes := []rune(chars)
	len1 := len(runes)
	time.Sleep(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	pwd := NewStringBuilder()
	for index := 0; index < lenth; index++ {
		pwd.Append(string(runes[rand.Intn(len1)]))
	}
	return pwd.ToString()
}

type StringBuilder struct {
	buf bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{buf: bytes.Buffer{}}
}

func (this *StringBuilder) Append(obj interface{}) *StringBuilder {
	this.buf.WriteString(fmt.Sprintf("%v", obj))
	return this
}

func (this *StringBuilder) ToString() string {
	return this.buf.String()
}

func Json2map(jsstr string) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsstr), &m)
	if err != nil {
		return
	}
	for key, value := range m {
		switch data := value.(type) {
		case string:
			fmt.Printf("map[%s]的值类型string,value = %s\n", key, data)
		case []string:
			fmt.Printf("map[%s]的值类型[]string,value = %s\n", key, data)
		case bool:
			fmt.Printf("map[%s]的值类型bool,value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]的值类型float64,value = %f\n", key, data)
		case []interface{}:
			fmt.Printf("map[%s]的值类型[]interface{},value = %v\n", key, data)
		}
	}
	result, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))
}