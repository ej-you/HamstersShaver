package mongo

import (
	"context"
	"sync"
	"time"

	goMongo "go.mongodb.org/mongo-driver/mongo"
	goMongoOptions "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)


var mongoClient *goMongo.Client
var once sync.Once


// получение клиента mongo
func getMongoClient() *goMongo.Client {
	var err error

	once.Do(func() {
		settings.InfoLog.Printf("Connect to mongo on %s...", settings.MongoAddr)

		// контекст для подключения к mongo
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// создание нового клиента
		options := goMongoOptions.Client().ApplyURI(settings.MongoAddr)
		mongoClient, err = goMongo.Connect(ctx, options)
		settings.DieIf(err)

		// проверка подключения
		err = mongoClient.Ping(ctx, nil)
		settings.DieIf(err)

		settings.InfoLog.Printf("Successfully connected to mongo")		
	})

	return mongoClient
}
