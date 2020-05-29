package token

import (
	glocfg "ds_server/support/utils/cfg"
	"ds_server/support/utils/cfg/common"
	"ds_server/support/utils/cfg/config"
	stringex "ds_server/support/utils/stringex"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-plugins/config/source/grpc"
	"net/http"
	"strings"
	"time"
)

//--------------------------------获取jwt的配置signkey--------------------------------
var (
	jwtName = "jwt"
	cfg     = &jwtCfg{}
)

type jwtCfg struct {
	common.Jwt
}

func InitConfig(address string, prefix string) *jwtCfg {
	source := grpc.NewSource(
		grpc.WithAddress(address),
		grpc.WithPath(prefix),
	)
	glocfg.Init(config.WithSource(source))
	err := config.C().App(jwtName, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

//-----------------------------------------------------------------------------------
var (
	TokenExpired     error  = errors.New("Token过期!")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("token错误!")
	TokenInvalid     error  = errors.New("token错误!")
	SignKey          string = "test"
)

type Subject struct {
	Uuid         string `json:"uuid"`
	Mobile       string `json:"mobile"`
	UserName     string `json:"username"`
	InvCodeAgent string `json:"invcodeagent"`
	InvCodeSelf  string `json:"Invcodeself"`
	jwt.StandardClaims
}
type JwtToken struct {
	SigningKey []byte
}

func (j *JwtToken) CreateToken(claims Subject) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
func (j *JwtToken) ParseToken(tokenString string) (*Subject, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Subject{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Subject); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
func (j *JwtToken) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Subject{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Subject); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

//Decode 解码
func (j *JwtToken) Decode(r *http.Request, w http.ResponseWriter) (bool, string, *Subject) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return false, "请求未携带token，无权限访问", nil
	}
	if !strings.HasPrefix(token, "Bearer ") {
		return false, "请求未携带token，无权限访问", nil
	}
	if stringex.Length(token) < 128 {
		return false, "请求未携带超过128位的token参数，无权限访问", nil
	}
	token = stringex.SubString(token, stringex.Length("Bearer "), stringex.Length(token)-stringex.Length("Bearer "))
	sub, err := j.ParseToken(token)
	if sub == nil {
		return false, err.Error(), nil
	}
	return true, "token验证成功!", sub
}

//func CreateTokenString(subject Subject) (tokenstring string) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, subject)
//	tokenstring, _ = token.SignedString([]byte(conval.TokenKey))
//	return
//}
//func SaveTokenToCache(subject *Subject, tokenstring string) error {
//	//保存
//	if err := redisex.Redis().Set(conval.TokenIDKeyPrefix+subject.Account, tokenstring, conval.TokenExpiredDate).Err(); err != nil {
//		return fmt.Errorf("[saveTokenToCache] 保存token到缓存发生错误，err:" + err.Error())
//	}
//	return nil
//}
//func GetTokenFromCache(subject *Subject) (token string, err error) {
//	// 获取
//	tokenCached, err := redisex.Redis().Get(conval.TokenIDKeyPrefix + subject.Account).Result()
//	if err != nil {
//		return token, fmt.Errorf("[getTokenFromCache] token不存在 %s", err)
//	}
//	return string(tokenCached), nil
//}
//func DelTokenFromCache(subject *Subject) (err error) {
//	//保存
//	if err = redisex.Redis().Del(conval.TokenIDKeyPrefix + subject.Account).Err(); err != nil {
//		return fmt.Errorf("[delTokenFromCache] 清空token 缓存发生错误，err:" + err.Error())
//	}
//	return nil
//}
