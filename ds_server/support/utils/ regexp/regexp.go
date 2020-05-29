package regexp

import (
	"regexp"
)

func VerifyUserAccount(userName string) bool {
	regular := "^[a-zA-Z][0-9a-zA-Z]{6,20}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(userName)
}

func VerifyPasswd(passwd string) bool {
	regular := "^[0-9a-zA-Z]{6,16}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(passwd)
}

func VerifyMobile(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func VerifyEmail(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyIDcard(idCard string) bool {
	pattern := `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(idCard)
}

//	vfy := &Veryfy{}
//	if !vfy.VerifyUserAccount("a16t@0") {
//		fmt.Println(errcode.ERROR_COMMON_REGEXP_USERACCOUNT_MSG)
//		return
//	}
//	if !vfy.VerifyPasswd("a1Maat@") {
//		fmt.Println(errcode.ERROR_COMMON_REGEXP_PASSWORD_MSG)
//		return
//	}
//	if !vfy.VerifyMobile("15882140525") {
//		fmt.Println(errcode.ERROR_COMMON_REGEXP_MOBILE_MSG)
//		return
//	}
//	if !vfy.VerifyEmail("wmh@aliyun.com") {
//		fmt.Println(errcode.ERROR_COMMON_REGEXP_EMAIL_MSG)
//		return
//	}
//	fmt.Println("all goes well!")