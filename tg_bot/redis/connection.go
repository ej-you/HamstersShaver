package redis


import (
	"context"
	"sync"
	"time"

	goRedis "github.com/redis/go-redis/v9"

	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)

var redisClient *goRedis.Client
var once sync.Once


// получение клиента redis
func GetRedisClient() *goRedis.Client {
	once.Do(func() {
		settings.InfoLog.Printf("Connect to redis on %s...", settings.RedisAddr)

		// создание нового клиента
		redisClient = goRedis.NewClient(&goRedis.Options{
			Addr: settings.RedisAddr,
			Password: settings.RedisPassword,
			DB: 0,
		})

		// контекст для выполнения запросов к redis
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// проверка подключения
		pong, err := redisClient.Ping(ctx).Result()
		settings.DieIf(err)

		settings.InfoLog.Printf("Successfully connected to redis: PING - %s", pong)		
	})

	return redisClient
}
