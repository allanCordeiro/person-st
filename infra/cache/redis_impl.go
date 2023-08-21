package cache

import (
	"encoding/json"

	"github.com/AllanCordeiro/person-st/application/domain"
	"github.com/gomodule/redigo/redis"
)

type RedisInstance struct {
	Pool *redis.Pool
}

func NewRedisInstance(pool *redis.Pool) *RedisInstance {
	return &RedisInstance{Pool: pool}
}

func (r *RedisInstance) Get(key string) (*domain.Person, error) {
	var data []byte
	conn := r.Pool.Get()
	defer conn.Close()
	data, err := redis.Bytes(conn.Do("GET", key))
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

func (r *RedisInstance) GetNickname(nickname string) error {
	conn := r.Pool.Get()
	defer conn.Close()
	_, err := redis.Bytes(conn.Do("GET", nickname))
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisInstance) Set(person domain.Person) error {
	var data []byte
	conn := r.Pool.Get()
	defer conn.Close()
	data, err := json.Marshal(person)
	if err != nil {
		return err
	}

	_, err = redis.Bytes(conn.Do("SET", person.Id, data))
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisInstance) SetNickname(nickname string) error {
	conn := r.Pool.Get()
	defer conn.Close()

	_, err := redis.Bytes(conn.Do("SET", nickname, ""))
	if err != nil {
		return err
	}
	return nil
}
