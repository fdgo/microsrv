package auth

import (
	"github.com/micro/micro/plugin"
	retpkg "microservice/jzapi/basic/return_pkg"
	timex "microservice/jzapi/basic/time_ex"
	"microservice/jzapi/lib/limit"
	"microservice/jzapi/lib/token"
	"net/http"
	"time"
)

func JWTAuthWrapper(token *token.JwtToken) plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO 从配置中心动态获取白名单URL
			if r.URL.Path == "/base/send/code" ||
				r.URL.Path == "/base/judge/code" ||
				r.URL.Path == "/user/regist/single/quick" ||
				r.URL.Path == "/user/regist/single/mobile" ||
				r.URL.Path == "/user/regist/single/account" ||
				r.URL.Path == "/user/multi/login" ||
				r.URL.Path == "/user/login/single/guest" ||
				r.URL.Path == "/user/login/single/mobile" ||
				r.URL.Path == "/user/login/single/account" ||

				r.URL.Path == "/admin/channels/list" ||
				r.URL.Path == "/admin/access/list" ||
				r.URL.Path == "/admin/roles/list" ||
				r.URL.Path == "/admin/roles/add" ||
				r.URL.Path == "/admin/roles/edit" ||
				r.URL.Path == "/admin/roles/delete" ||
				r.URL.Path == "/admin/access/delete" ||
				r.URL.Path == "/admin/access/add" ||
				r.URL.Path == "/admin/users/list" ||
				r.URL.Path == "/admin/users/add" ||
				r.URL.Path == "/admin/users/delete" ||
				r.URL.Path == "/admin/login" {
				h.ServeHTTP(w, r)
				return
			}
			istokenok, msg, sub := token.Decode(r, w)
			if !istokenok {
				retpkg.NorMalResponse(http.StatusUnauthorized, msg, nil, r, w)
				return
			}
			r.Header.Set("X-Head-Username", sub.DeviceId)
			r.Header.Set("X-Head-Accountid", sub.AccountId)
			r.Header.Set("X-Head-Timestamp", timex.TimeStampToTimeStr(sub.ExpiresAt))
			h.ServeHTTP(w, r)
		})
	}
}

var (
	rl = limit.New(2000, time.Second)
)

func LimitWrapper() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rl.Limit() {
				retpkg.NorMalResponse(http.StatusOK, "请求频率太高！", nil, r, w)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
