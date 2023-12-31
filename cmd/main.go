package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/AllanCordeiro/person-st/infra/cache"
	"github.com/AllanCordeiro/person-st/infra/database"
	"github.com/AllanCordeiro/person-st/infra/queue"
	"github.com/AllanCordeiro/person-st/infra/webserver"
	"github.com/AllanCordeiro/person-st/infra/worker"
)

func main() {
	dbURL := os.Getenv("DBURL")
	if dbURL == "" {
		dbURL = "postgres://rinha:rinha123@db/rinhadb?sslmode=disable"
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "redis:"+redisPort)
		},
	}

	//TODO: await rabbitMQ starting up
	time.Sleep(15 * time.Second)
	rabbitMQ, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	defer rabbitMQ.Close()

	mqChannel, err := rabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer mqChannel.Close()

	personDB := database.NewPersonDB(db)
	personCache := cache.NewRedisInstance(pool)
	personQueue := queue.NewRabbitMQImpl(mqChannel)
	personQueue.QueueDeclare("person.created")

	personDB.Warmup()

	worker := worker.NewCreatePersonWorker(personDB, *personQueue)
	go worker.Run()

	webserver.Serve(personDB, personCache, personQueue)
}
