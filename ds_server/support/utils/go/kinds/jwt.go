package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/ha666/golibs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.hfjy.com/service/user-service/us_initials/config"
	"net/http"
	"strings"
)

// 中间件，检查用户token
func JWTAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		//region 获取jwt信息
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带token，无权限访问",
			})
			return
		}
		if !strings.HasPrefix(token, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带token参数，无权限访问",
			})
			return
		}
		if len(token) < 128 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "请求未携带超过128位的token参数，无权限访问",
			})
			return
		}

		token = golibs.SubString(token, len("Bearer "), len(token)-len("Bearer "))
		//endregion

		//region 解析jwt信息
		var jwtInfo JwtInfo
		tokenInfo, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return []byte(""), errors.New("签名方法不正确")
			}
			arrs := strings.Split(token.Raw, ".")
			if len(arrs) == 3 {
				b64, err := base64.RawStdEncoding.DecodeString(arrs[1])
				if err != nil {
					return []byte(""), err
				}
				err = json.Unmarshal(b64, &jwtInfo)
				if err != nil {
					return []byte(""), err
				}
				if jwtInfo.AppID < 10000 {
					return []byte(""), errors.New("无效的应用标识")
				}
				if jwtInfo.AppID == config.Config.App.Id {
					return []byte(config.Config.Jwt.Key), nil
				}
			}
			return []byte(""), errors.New("没有找到应用密钥信息")
		})

		if err != nil {
			if !strings.Contains(err.Error(), "没有找到应用密钥信息") {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":    401,
					"message": "解析token失败:" + err.Error(),
				})
				return
			}
		}

		_, ok := tokenInfo.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "解析token失败了",
			})
			return
		}
		//endregion

		//region 验证过期时间
		if jwtInfo.Exp < golibs.Unix() {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    401,
				"message": "jwt信息已过期",
			})
			return
		}
		//endregion

		//region 验证appId及签名

		if jwtInfo.AppID == config.Config.App.Id {

			//region 验证签名
			if tokenInfo.Valid {
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":    401,
					"message": "jwt信息签名错误",
				})
				return
			}
			//endregion

		} else {

			//region 获取新的jwt信息
			loginReqParams := fmt.Sprintf(`{"type":2,"appId":%d,"token":"%s"}`, config.Config.App.Id, token)
			code, body, err := golibs.PostBody(config.Config.Sso.Url+"/login", loginReqParams)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":    401,
					"message": "获取新的jwt信息时出错:" + err.Error(),
				})
				return
			}
			if code != 200 {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":    401,
					"message": fmt.Sprintf("获取新的jwt信息时出错:%d,%s", code, body),
				})
				return
			}
			var response GetJwtResponse
			err = json.Unmarshal([]byte(body), &response)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":    401,
					"message": "获取新的jwt信息解析出错:" + err.Error(),
				})
				return
			}
			if response.Code != 200 {
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":    401,
					"message": fmt.Sprintf("获取jwt出错:%d,%s", response.Code, response.Message),
				})
				return
			}
			//endregion

			//region 返回新的jwt信息
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    1111,
				"message": "",
				"data": map[string]interface{}{
					"token": response.Data.Token,
				},
			})
			return
			//endregion

		}
		//endregion

	}
}

// JWT基本信息
type JwtInfo struct {
	AppID    int    `json:"appId"`
	Exp      int64  `json:"exp"`
	Iat      int64  `json:"iat"`
	Token    string `json:"token"`
	UserCode string `json:"userCode"`
	RoleIds  string `json:"roleIds"`
}

// 从SSO获取新的jwt信息响应结果
type GetJwtResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}
