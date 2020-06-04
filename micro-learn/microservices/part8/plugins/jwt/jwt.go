package jwt

import "github.com/wangmhgo/microservice-project/go-micro-learn/microservices/part8/basic"

// Jwt 配置 接口
type Jwt struct {
	SecretKey string `json:"secretKey"`
}

// init 初始化Redis
func init() {
	basic.Register(initJwt)
}

func initJwt() {

}