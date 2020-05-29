package middleware

import (
	"gitee.com/ha666/golibs"
	"gitee.com/ha666/logs"
	"github.com/gin-gonic/gin"
	"gitlab.hfjy.com/service/user-service/us_cache"
	"gitlab.hfjy.com/service/user-service/us_models"
	"gitlab.hfjy.com/service/user-service/worker/api_statistics_worker"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	_ver = "1.0"
)

var (
	appDB = us_models.App{}
)

// 中间件，用户中心开放平台(User Open Platform)
func UOPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		//logs.Info("path:%s,form:%+v", c.Request.URL.Path, c.Request.PostForm)

		start := time.Now()
		time.Sleep(time.Millisecond * 3)

		//region 验证appId参数
		appId := c.DefaultPostForm("appId", "")
		if len(appId) != 5 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带正确的appId，无权限访问",
			})
			return
		}
		appIdInt, err := strconv.Atoi(appId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带正确的appId，解析appId出错:" + err.Error(),
			})
			return
		}
		appSecret, err := us_cache.GetAppSecret(appIdInt)
		if err != nil {
			logs.Error("【UOPMiddleware】获取密钥失败%d,err:%s", appIdInt, err.Error())
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "密钥获取失败，请检查后重试",
			})
			return
		}
		if len(appSecret) != 32 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带正确的appId，请检查后重试",
			})
			return
		}
		//endregion

		//region 验证ver参数
		ver := c.DefaultPostForm("ver", "")
		if ver != _ver {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带正确的ver，无权限访问",
			})
			return
		}
		//endregion

		//region 验证time参数
		timeVal := c.DefaultPostForm("time", "")
		if len(timeVal) != 10 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带正确的time，无权限访问",
			})
			return
		}
		if !golibs.IsNumber(timeVal) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带正确格式的time，无权限访问",
			})
			return
		}
		timeUnix, err := strconv.ParseInt(timeVal, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "解析参数time失败:" + err.Error(),
			})
			return
		}
		if timeUnix < golibs.Unix()-300 || timeUnix > golibs.Unix()+300 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "time参数时间间隔超过5分钟",
			})
			return
		}
		//endregion

		//region 验证签名
		if !golibs.ApiSignIsValid(c.Request.PostForm, appSecret) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "签名错误",
			})
			return
		}
		//endregion

		c.Next()

		code := c.GetInt("code")
		if code > 0 {
			t := golibs.Since(start)
			api := strings.Replace(c.Request.URL.String(), "/api/open/", "", -1)
			m := us_models.ApiStatisticsModel{
				CurrentDate:   golibs.GetDate(time.Now()),
				AppId:         appIdInt,
				ApiName:       api,
				IsSuccess:     code == 200,
				StatusCode:    code,
				ConsumingTime: int(t),
			}
			api_statistics_worker.ApiStatisticsChannel <- golibs.ToJson(&m)
		} else {
			logs.Error("【UOPMiddleware】没有找到code")
		}
	}
}
