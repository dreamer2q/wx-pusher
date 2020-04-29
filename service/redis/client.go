package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"wxServ/service/log"
)

var client *redis.Client

func Instance() *redis.Client {
	return client
}

func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Instance().Debugf("redis client: %s", pong)
}

const (
	storeTimeout = 30 * 24 * time.Hour
)

func Store(val interface{}) (key string, err error) {
	t := time.Now().UnixNano()
	key = strconv.FormatInt(t, 10)
	err = client.Set(key, val, storeTimeout).Err()
	return
}

func Load(key string, out interface{}) error {
	ret, err := client.Get(key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(ret, out)
}
