package persistance

import (
	"github.com/garyburd/redigo/redis"
	"os"
)

type Store interface {
	Get(key string) (value int64, ok bool)
	Set(key string, value int64)
}

type RealStore struct{}

func (store RealStore) Get(key string) (value int64, ok bool) {

	client, err := redis.Dial("tcp", os.Getenv("WB_DB_HOST"))
	if err != nil {
		ok = false
	}
	defer client.Close()

	value, err = redis.Int64(client.Do("GET", key))
	ok = err == nil
	return
}

func (store RealStore) Set(key string, value int64) {
	client, err := redis.Dial("tcp", os.Getenv("WB_DB_HOST") + ":6379")
	if err != nil {
		return
	}
	defer client.Close()

	_, err = client.Do("SET", key, value)
	if err != nil {
		return
	}
}
