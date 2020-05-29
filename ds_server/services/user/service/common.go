package service

import (
	dao "ds_server/services/user/dao"
	rds "ds_server/support/lib/redisex"
	token "ds_server/support/utils/auth"
	"ds_server/support/utils/constex"
	time_ex "ds_server/support/utils/timex"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type UserService struct {
	RedisCache                    *redis.Client
	DsUserBasicinfoDao            *dao.DsUserBasicinfoMgr
	DsSysInfoDao                  *dao.DsSysInfoMgr
	DsUserMemberDepositHistoryDao *dao.DsUserMemberDepositHistoryMgr
	DsUserMemberAgentDao          *dao.DsUserMemberAgentMgr
	DsUserMemberAccountDao        *dao.DsUserMemberAccountMgr
	DsUserMemberClassDao          *dao.DsUserMemberClassMgr
	DsUserAgentClassDao           *dao.DsUserAgentClassMgr
}

func newtoken(uuid, mobie, invcodeself, invcodeagent string) string {
	subject := &token.Subject{uuid, mobie, "", invcodeagent, invcodeself, jwt.StandardClaims{ExpiresAt: constex.JwtCfg.Exptime + time_ex.GetCurrentTimeStamp(), Issuer: "ds.srv.user"}}
	jwt := &token.JwtToken{SigningKey: []byte(constex.JwtCfg.SecretKey)}
	token, _ := jwt.CreateToken(*subject)
	exptime, _ := time.ParseDuration("+" + strconv.FormatInt(constex.JwtCfg.Exptime, 10) + "s")
	rds.RedisInstanceg().Set(constex.REDIS_USER_TOKEN+mobie, token, exptime).Err()
	return token
}
