package handler

import (
	"encoding/json"
	"net/http"
	"time"

	hystrix_go "github.com/afex/hystrix-go/hystrix"
	auth "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/auth/proto/auth"
	"github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/plugins/session"
	us "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part7/user-srv/proto/user"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
)

var (
	serviceClient us.UserService
	authClient    auth.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	cl := hystrix.NewClientWrapper()(client.DefaultClient)
	serviceClient = us.NewUserService("fred.micro.shop.srv.user", cl)
	authClient = auth.NewService("fred.micro.shop.srv.auth", cl)
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(ctx, &us.Request{
		UserName: r.Form.Get("userName"),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	if rsp.User.Pwd == r.Form.Get("pwd") {
		response["success"] = true

		// 干掉密码返回
		rsp.User.Pwd = ""
		response["data"] = rsp.User
		log.Logf("[Login] 密码校验完成，生成token...")

		// 生成token
		rsp2, err := authClient.MakeAccessToken(ctx, &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Name,
		})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Logf("[Login] token %s", rsp2.Token)
		response["token"] = rsp2.Token

		// 同时将token写到cookies中
		w.Header().Add("set-cookie", "application/json; charset=utf-8")
		// 过期30分钟
		expire := time.Now().Add(30 * time.Minute)
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)

		// 同步到session中
		sess := session.GetSession(w, r)
		sess.Values["userId"] = rsp.User.Id
		sess.Values["userName"] = rsp.User.Name
		_ = sess.Save(r, w)
	} else {
		response["success"] = false
		response["error"] = &Error{
			Detail: "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Logf("token获取失败")
		http.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(ctx, &auth.Request{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref":     time.Now().UnixNano(),
		"success": true,
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TestSession(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)

	if v, ok := sess.Values["path"]; !ok {
		sess.Values["path"] = r.URL.Query().Get("path")
		log.Logf("path:" + r.URL.Query().Get("path"))
	} else {
		log.Logf(v.(string))
	}

	log.Logf(sess.ID)
	log.Logf(sess.Name())

	w.Write([]byte("OK"))
}
