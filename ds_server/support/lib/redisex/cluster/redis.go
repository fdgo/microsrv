package main

import (
	"ds_server/support/utils/constex"
	"github.com/go-redis/redis"
	"strings"
	"sync"
	"time"
)

// NewCache("192.168.60.10:6379,192.168.60.11:6379,192.168.60.12:6379,192.168.60.13:6379,192.168.60.14:6379,192.168.60.15:6379")
func NewCache(servers string) (*Cache, error) {
	opt := new(redis.ClusterOptions)
	opt.Addrs = strings.Split(servers, ",")

	c := new(Cache)
	c.client = redis.NewClusterClient(opt)

	return c, nil
}

var redislock sync.Mutex
var redisInstarnce *Cache

type Cache struct {
	client *redis.ClusterClient
}

func RedisClusterInstanceg() (*Cache) {
	if redisInstarnce != nil {
		return redisInstarnce
	}
	redislock.Lock()
	defer redislock.Unlock()
	if redisInstarnce != nil {
		return redisInstarnce
	}
	return newcache()
}
func newcache() (*Cache) {
	opt := new(redis.ClusterOptions)
	opt.Addrs = strings.Split( constex.RedisCluterCfg.Ip, ",")
	c := new(Cache)
	c.client = redis.NewClusterClient(opt)
	return c
}

//----------------------------------------------------------------------------

func (c *Cache) Close() error {
	return c.client.Close()
}

//----------------------------------------------------------------------------

func (c *Cache) Exists(key string) (int64, error) {
	return c.client.Exists(key).Result()
}

//----------------------------------------------------------------------------

func (c *Cache) Del(key string) error {
	return c.client.Del(key).Err()
}

//----------------------------------------------------------------------------

func (c *Cache) Expire(key string, duration time.Duration) error {
	_, err := c.client.Expire(key, duration).Result()
	if err != nil {
		return err
	}

	return nil
}

//----------------------------------------------------------------------------

func (c *Cache) GetKey(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *Cache) SetKey(key string, value string) error {
	return c.client.Set(key, value, 0).Err()
}

func (c *Cache) SetKeyAndExpire(key string, value string, expiration time.Duration) error {
	return c.client.Set(key, value, expiration).Err()
}

func (c *Cache) Append(key string, value string) error {
	return c.client.Append(key, value).Err()
}

//----------------------------------------------------------------------------

func (c *Cache) FieldExist(key string, field string) (bool, error) {
	return c.client.HExists(key, field).Result()
}

func (c *Cache) GetField(key string, field string) (string, error) {
	value, err := c.client.HGet(key, field).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cache) GetAllFields(key string) (map[string]string, error) {
	return c.client.HGetAll(key).Result()
}

func (c *Cache) SetField(key string, field string, value string) error {
	return c.client.HSet(key, field, value).Err()
}

func (c *Cache) SetFields(key string, fields map[string]interface{}) error {
	return c.client.HMSet(key, fields).Err()
}

func (c *Cache) DelField(key string, field string) error {
	return c.client.HDel(key, field).Err()
}

func (c *Cache) HLen(key string) (int64, error) {
	return c.client.HLen(key).Result()
}

//----------------------------------------------------------------------------

func (c *Cache) Incr(key string) (int64, error) {
	return c.client.Incr(key).Result()
}

func (c *Cache) Decr(key string) (int64, error) {
	return c.client.Decr(key).Result()
}

func (c *Cache) IncrField(key string, field string) (int64, error) {
	return c.client.HIncrBy(key, field, 1).Result()
}

func (c *Cache) DecrField(key string, field string) (int64, error) {
	return c.client.HIncrBy(key, field, -1).Result()
}

//----------------------------------------------------------------------------

func (c *Cache) Push(key string, value string) error {
	return c.client.LPush(key, value).Err()
}

func (c *Cache) Pop(key string) (string, error) {
	return c.client.LPop(key).Result()
}

func (c *Cache) Range(key string, start int, end int) ([]string, error) {
	return c.client.LRange(key, int64(start), int64(end)).Result()
}

func (c *Cache) LLen(key string) (int64, error) {
	return c.client.LLen(key).Result()
}

func (c *Cache) LAll(key string) ([]string, error) {
	n, err := c.client.LLen(key).Result()
	if err != nil {
		return nil, err
	}

	if n < 1 {
		return []string{}, nil
	}

	return c.client.LRange(key, 0, n-1).Result()
}

//----------------------------------------------------------------------------

func (c *Cache) Publish(channel string, message string) error {
	return c.client.Publish(channel, message).Err()
}

//----------------------------------------------------------------------------
