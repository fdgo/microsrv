package redisex

import (
	"ds_server/support/utils/cfg/config"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
	"strings"
	"sync"
)

var (
	client *redis.Client
	mutex  sync.RWMutex
)

// redisex redisex 配置
type RedisCfg struct {
	Enabled  bool           `json:"enabled"`
	Conn     string         `json:"conn"`
	Password string         `json:"password"`
	DBNum    int            `json:"dbNum"`
	Timeout  int            `json:"timeout"`
	Sentinel *RedisSentinel `json:"sentinel"`
}
type RedisSentinel struct {
	Enabled bool   `json:"enabled"`
	Master  string `json:"master"`
	XNodes  string `json:"nodes"`
	nodes   []string
}

// Nodes redisex 哨兵节点列表
func (s *RedisSentinel) GetNodes() []string {
	if len(s.XNodes) != 0 {
		for _, v := range strings.Split(s.XNodes, ",") {
			v = strings.TrimSpace(v)
			s.nodes = append(s.nodes, v)
		}
	}
	return s.nodes
}
func initRedis() {
	log.Log("[initRedis] 初始化Redis...")
	c := config.C()
	cfg := &RedisCfg{}
	err := c.App("redis", cfg)
	if err != nil {
		log.Logf("[initRedis] %s", err)
	}
	if !cfg.Enabled {
		log.Logf("[initRedis] 未启用redis")
		return
	}
	// 加载哨兵模式
	if cfg.Sentinel != nil && cfg.Sentinel.Enabled {
		log.Log("[initRedis] 初始化Redis，哨兵模式...")
		initSentinel(cfg)
	} else { // 普通模式
		log.Log("[initRedis] 初始化Redis，普通模式...")
		initSingle(cfg)
	}
	log.Log("[initRedis] 初始化Redis，检测连接...")
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Logf("[initRedis] 初始化Redis，检测连接Ping... %s", pong)
}
func initSentinel(redisConfig *RedisCfg) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.Sentinel.Master,
		SentinelAddrs: redisConfig.Sentinel.GetNodes(),
		DB:            redisConfig.DBNum,
		Password:      redisConfig.Password,
	})

}
func initSingle(redisConfig *RedisCfg) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Conn,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DBNum,    // use default DB
	})
}

func RedisInstanceg() *redis.Client {
	if client != nil {
		return client
	}
	mutex.Lock()
	defer mutex.Unlock()
	if client != nil {
		return client
	}
	initRedis()
	return client
}
