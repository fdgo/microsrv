package redisex

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}

func standalon_redis_test() {
	fmt.Printf("standalon_redis_test")

	client := redis.NewClient(&redis.Options{
		Addr:     "120.27.239.127:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
		err_handler(err)
	}
	fmt.Printf("ping result: %s\n", pong)

	fmt.Printf("----------------------------------------\n")

	// set / get test
	fmt.Printf("set/get test\n")
	err = client.Set("foo", "bar", 0).Err()
	if err != nil {
		fmt.Printf("try set key[foo] to value[bar] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	err = client.Set("foo1", "bar1", time.Hour*2).Err()
	if err != nil {
		fmt.Printf("try set key[foo1] to value[bar1] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	// get value
	value, err := client.Get("foo").Result()
	if err != nil {
		fmt.Printf("try get key[foo] error[%s]\n", err.Error())
		// err_handler(err)
	}

	fmt.Printf("key[foo]'s value is %s\n", value)

	value, err = client.Get("foo1").Result()
	if err != nil {
		fmt.Printf("try get key[foo1] error[%s]\n", err.Error())
		// err_handler(err)
	}

	fmt.Printf("key[foo1]'s value is %s\n", value)

	value, err = client.Get("foo2").Result()
	if err != nil {
		fmt.Printf("try get key[foo2] error[%s]\n", err.Error())
		// err_handler(err)
	}

	fmt.Printf("key[foo2]'s value is %s\n", value)

	// get ttl
	duration, err := client.TTL("foo").Result()
	if err != nil {
		fmt.Printf("try get ttl of key[foo] error[%s]\n", err.Error())
		err_handler(err)
	}

	fmt.Printf("key[foo]' ttl is [%s] %fs\n",
		duration.String(), duration.Seconds())

	duration, err = client.TTL("foo1").Result()
	if err != nil {
		fmt.Printf("try get ttl of key[foo1] error[%s]\n", err.Error())
		err_handler(err)
	}

	fmt.Printf("key[foo1]' ttl is [%s] %ds\n",
		duration.String(), int64(duration.Seconds()))

	fmt.Printf("----------------------------------------\n")

	// list test
	fmt.Printf("list test\n")

	err = client.RPush("tqueue", "tmsg1").Err()
	if err != nil {
		fmt.Printf("rpush list[tqueue] error[%s]\n", err.Error())
		err_handler(err)
	}

	list_len, err := client.LLen("tqueue").Result()
	if err != nil {
		fmt.Printf("try get len of list[tqueue] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	fmt.Printf("len of list[tqueue] is %d\n", list_len)

	result, err := client.BLPop(time.Second*1, "tqueue").Result()
	if err != nil {
		fmt.Printf("blpop list[tqueue] error[%s]\n", err.Error())
		err_handler(err)
	}
	fmt.Printf("blpop list[tqueue], value[%s]\n", result[1])

	fmt.Printf("----------------------------------------\n")

	fmt.Printf("hmap test\n")

	err = client.HSet("tmap", "1", "f1").Err()
	if err != nil {
		fmt.Printf("try hset map[tmap] field[1] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	err = client.HSet("tmap", "2", "f2").Err()
	if err != nil {
		fmt.Printf("try hset map[tmap] field[2] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	kv_map := make(map[string]interface{})
	kv_map["3"] = "f3"
	kv_map["4"] = "f4"

	err = client.HMSet("tmap", kv_map).Err()
	if err != nil {
		fmt.Printf("try mset map[tmap] field[3] field[4] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	map_len, err := client.HLen("tmap").Result()
	if err != nil {
		fmt.Printf("try get len of map[tmap] error[%s]\n", err.Error())
		err_handler(err)
	}
	fmt.Printf("len of map[tmap] is %d\n", map_len)

	// get map value
	value, err = client.HGet("tmap", "1").Result()
	if err != nil {
		fmt.Printf("try get field[1] value of map[tmap] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	fmt.Printf("field[1] value of map[tmap] is %s\n", value)

	// hgetall
	result_kv, err := client.HGetAll("tmap").Result()
	if err != nil {
		fmt.Printf("try hgetall map[tmap] error[%s]\n", err.Error())
		err_handler(err)
	}

	for k, v := range result_kv {
		fmt.Printf("map[tmap] %s = %s\n", k, v)
	}

	fmt.Printf("----------------------------------------\n")

	fmt.Printf("pubsub test\n")

	pubsub := client.Subscribe("test_channel")

	_, err = pubsub.Receive()
	if err != nil {
		fmt.Printf("try subscribe channel[test_channel] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	// go channel to used to receive message
	ch := pubsub.Channel()

	// publish a message
	err = client.Publish("test_channel", "hello").Err()
	if err != nil {
		fmt.Printf("try publish message to channel[test_channel] error[%s]\n",
			err.Error())
		err_handler(err)
	}

	time.AfterFunc(time.Second*2, func() {
		_ = pubsub.Close()
	})

	// consume message
	for {
		msg, ok := <-ch
		if !ok {
			break
		}

		fmt.Printf("recv message[%s] from channel[%s]\n",
			msg.Payload, msg.Channel)
	}
}

func main() {
	fmt.Println("redis_test!")

	standalon_redis_test()

}