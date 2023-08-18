package cache

import (
	"encoding/json"

	"github.com/AllanCordeiro/person-st/application/domain"
	"github.com/gomodule/redigo/redis"
)

type RedisInstance struct {
	Conn redis.Conn
}

var pool *redis.Pool

func NewRedisInstance(conn redis.Conn) *RedisInstance {
	return &RedisInstance{Conn: conn}
}

func (r *RedisInstance) Get(key string) (*domain.Person, error) {
	var data []byte
	conn := pool.Get()
	defer conn.Close()
	data, err := redis.Bytes(conn.Do("GET", key))
	//data, err := redis.Bytes(r.Conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	var person domain.Person
	err = json.Unmarshal(data, &person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *RedisInstance) Set(person domain.Person) error {
	var data []byte
	data, err := json.Marshal(person)
	if err != nil {
		return err
	}

	_, err = redis.Bytes(r.Conn.Do("SET", person.Id, data))
	if err != nil {
		return err
	}
	return nil
}
