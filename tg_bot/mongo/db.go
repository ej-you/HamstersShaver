package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"
	
	"github.com/google/uuid"
	goMongo "go.mongodb.org/mongo-driver/mongo"
	goMongoOptions "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ej-you/HamstersShaver/tg_bot/mongo/schemas"
	
	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// структура для запросов к mongo
type MongoDB struct {
	dbConnect *goMongo.Database
}
func NewMongoDB() *MongoDB {
	return &MongoDB{
		dbConnect: getMongoClient().Database(settings.MongoDB),
	}
}


// интерфейс для использования струкутры в качестве документа коллекции MongoDB
type MongoDBCollection interface {
	// возвращает название коллекции
	CollectionName() string
	// валидация структуры для вставки документа в коллекцию
	Validate() error
}


// словарь для запросов к БД на получение и обновление документов
type AnyCollectionData map[string]interface{}


// добавление записи data в коллекцию (указана с помощью MongoDBCollection)
func (this MongoDB) InsertOne(data MongoDBCollection) error {
	var err error

	if err = data.Validate(); err != nil {
		return fmt.Errorf("collection %s: insert one: %w", data.CollectionName(), err)
	}

	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := this.dbConnect.Collection(data.CollectionName())
	// пропускаем возвращаемый результат с ID созданного документа
	_, err = collection.InsertOne(ctx, data)

	if err != nil {
		dbErr := customErrors.DBError(fmt.Sprintf("failed to insert document to %s collection", data.CollectionName()))
		return fmt.Errorf("collection %s: insert one: %v: %w", data.CollectionName(), err, dbErr)
	}
	return nil
}

// обновление записи данными updater с переданным id в коллекции collection
func (this MongoDB) UpdateByID(collectionName string, id uuid.UUID, updater AnyCollectionData) error {
	var err error

	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := this.dbConnect.Collection(collectionName)
	// пропускаем возвращаемый результат с ID обновлённого документа и другими данными
	_, err = collection.UpdateByID(ctx, id, AnyCollectionData{"$set": updater})

	if err != nil {
		dbErr := customErrors.DBError(fmt.Sprintf("failed to update document in %s collection", collectionName))
		return fmt.Errorf("collection %s: update one: %v: %w", collectionName, err, dbErr)
	}
	return nil
}


// получение одной записи транзакции, подходящей под фильтр filter
func (this MongoDB) GetTransactionByFilter(filter AnyCollectionData, result *schemas.Transaction) error {
	var err error

	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// поиск транзакции, подходящей под фильтр filter
	collection := this.dbConnect.Collection(result.CollectionName())
	err = collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		mongoErrNoDocuments := goMongo.ErrNoDocuments
		// если не найдено результатов
		if errors.Is(err, mongoErrNoDocuments) {
			dbNotFoundErr := customErrors.DBNotFoundError("no transaction found by given filter")
			return fmt.Errorf("collection %s: get transaction by filter %v: %v: %w", result.CollectionName(), filter, err, dbNotFoundErr)
		}
		// неизвестная ошибка
		dbErr := customErrors.DBError("failed to get transaction")
		return fmt.Errorf("collection %s: get transaction by filter: %v: %w", result.CollectionName(), err, dbErr)
	}
	return nil
}

// получение последней записи транзакции
func (this MongoDB) GetLastTransaction(result *schemas.Transaction) error {
	var err error

	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := this.dbConnect.Collection(result.CollectionName())
	// указываем в опциях поиска сортировку в обратном порядке относительно добавления документов в коллекцию
	findOptions := goMongoOptions.FindOne().SetSort(map[string]int{"$natural": -1})
	err = collection.FindOne(ctx, struct{}{}, findOptions).Decode(&result)

	if err != nil {
		mongoErrNoDocuments := goMongo.ErrNoDocuments
		// если не найдено результатов
		if errors.Is(err, mongoErrNoDocuments) {
			dbNotFoundErr := customErrors.DBNotFoundError("no one transaction exists")
			return fmt.Errorf("collection %s: get last transaction: %v: %w", result.CollectionName(), err, dbNotFoundErr)
		}
		// неизвестная ошибка
		dbErr := customErrors.DBError("failed to get transaction")
		return fmt.Errorf("collection %s: get last transaction: %v: %w", result.CollectionName(), err, dbErr)
	}
	return nil
}
