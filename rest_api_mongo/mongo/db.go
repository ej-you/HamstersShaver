package mongo

import (
	"context"
	"fmt"
	"time"
	
	goMongo "go.mongodb.org/mongo-driver/mongo"

	"github.com/ej-you/HamstersShaver/rest_api_mongo/settings"
)


const contextTimeout = 5*time.Second
// псевдоним для исключения лишних импортов модуля mongo-driver
var ErrNoDocuments error = goMongo.ErrNoDocuments


// структура для запросов к mongo
type MongoDB struct {
	dbConnect *goMongo.Database
}
func NewMongoDB() *MongoDB {
	return &MongoDB{
		dbConnect: getMongoClient().Database(settings.MongoDB),
	}
}


// для создания записей в коллекциях
type MongoDBCreator interface {
	CreatorCollectionName() string
}
// для получения данных из коллекций
type MongoDBData interface {
	DataCollectionName() string
}
// для обновления данных в коллекциях
type MongoDBUpdater interface {
	UpdateCollectionName() string
}
// для поиска данных в коллекциях
type MongoDBFilter interface {
	FilterCollectionName() string
}


// добавление записи data в коллекцию
func (this MongoDB) Insert(data MongoDBCreator) error {
	var err error

	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	collection := this.dbConnect.Collection(data.CreatorCollectionName())
	// пропускаем возвращаемое значение с ID созданного документа
	_, err = collection.InsertOne(ctx, data)

	if err != nil {
		return fmt.Errorf("insert: collection %s: %w", data.CreatorCollectionName(), err)
	}
	return nil
}


// получение одной записи, подходящей под фильтр filter (result - указатель на структуру)
func (this MongoDB) GetOneByFilter(filter MongoDBFilter, result MongoDBData) error {
	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	// поиск записей, подходящих под фильтр filter
	collection := this.dbConnect.Collection(filter.FilterCollectionName())
	err := collection.FindOne(ctx, filter).Decode(result)

	if err != nil {
		return fmt.Errorf("get one by filter: collection %s: %w", filter.FilterCollectionName(), err)
	}
	return nil
}

// получение всех записей, подходящих под фильтр filter (result - указатель на срез структур, предполагается []MongoDBData)
func (this MongoDB) GetManyByFilter(filter MongoDBFilter, results any) error {
	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	// поиск записей, подходящих под фильтр filter
	collection := this.dbConnect.Collection(filter.FilterCollectionName())
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("get many by filter: collection %s: %w", filter.FilterCollectionName(), err)
	}

	// декодируем ответ в results
	if err = cursor.All(ctx, results); err != nil {
		return fmt.Errorf("get many by filter: collection %s: %w", filter.FilterCollectionName(), err)
	}
	return nil
}

// обновление записей коллекции под фильтром filter данными updateData (возвращает кол-во изменённых записей)
func (this MongoDB) Update(filter MongoDBFilter, updateData MongoDBUpdater) (int64, error) {
	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	collection := this.dbConnect.Collection(filter.FilterCollectionName())
	// обновляем и получаем кол-во обновлённых записей
	result, err := collection.UpdateMany(ctx, filter, map[string]any{"$set": updateData})

	if err != nil {
		return 0, fmt.Errorf("update: collection %s: %w", filter.FilterCollectionName(), err)
	}
	return result.ModifiedCount, nil
}

// удаление записей из коллекции под фильтром filter (возвращает кол-во удалённых записей)
func (this MongoDB) Delete(filter MongoDBFilter) (int64, error) {
	// контекст для запроса к mongo
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	collection := this.dbConnect.Collection(filter.FilterCollectionName())
	// удаляем и получаем кол-во удалённых записей
	result, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return 0, fmt.Errorf("delete: collection %s: %w", filter.FilterCollectionName(), err)
	}
	return result.DeletedCount, nil
}
