package redis

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/ej-you/HamstersShaver/tg_bot/settings"
)


// установка строкового значения в кэш
func SetString(key, value string) error {
	// если передана пустая строка
	if len(value) == 0 {
		emptyStringError := errors.New("empty string was given")
	    settings.ErrorLog.Println("Failed to set string value in Redis:", emptyStringError)
	    return emptyStringError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// установка ключа-значения в redis
	err := GetRedisClient().Set(ctx, key, value, 0).Err()
	if err != nil {
	    settings.ErrorLog.Println("Failed to set string value in Redis:", err)
	    return err
	}
	return nil
}

// получение строкового значения из кэша
func GetString(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// получение значения из redis по ключу
	value, err := GetRedisClient().Get(ctx, key).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
		    settings.ErrorLog.Println("Failed to get string value from Redis:", err)
		}
	    return "", err
	}
	return value, nil
}


// установка значения среза строк в кэш
func SetStringSlice(key string, value []string) error {
	// если передан пустой список
	if len(value) == 0 {
		emptySliceError := errors.New("empty slice was given")
	    settings.ErrorLog.Println("Failed to set string slice value in Redis:", emptySliceError)
	    return emptySliceError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// сериализация среза строк в строковое значение
	marshaledValue, err := json.Marshal(value)
	if err != nil {
	    settings.ErrorLog.Println("Failed to marshal string slice to set it in Redis:", err)
		return err
	}

	// установка ключа-значения в redis
	err = GetRedisClient().Set(ctx, key, marshaledValue, 0).Err()
	if err != nil {
	    settings.ErrorLog.Println("Failed to set string slice value in Redis:", err)
	    return err
	}
	return nil
}

// получение значения среза строк из кэша
func GetStringSlice(key string) ([]string, error) {
	var unmarshaledValue []string

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// получение значения из redis по ключу
	value, err := GetRedisClient().Get(ctx, key).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
		    settings.ErrorLog.Println("Failed to get string slice value from Redis:", err)
		}
	    return unmarshaledValue, err
	}

	// десериализация строкового значения обратно в срез строк
	if err = json.Unmarshal([]byte(value), &unmarshaledValue); err != nil {
	    settings.ErrorLog.Println("Failed to unmarshal value to string slice after getting it from Redis:", err)
		return unmarshaledValue, err
	}
	return unmarshaledValue, nil
}


// удаление ключа из кэша
func DeleteKey(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// удаление значения из redis по ключу (первое возвращаемое значение - кол-во удалённых ключей)
	if _, err := GetRedisClient().Del(ctx, key).Result(); err != nil {
	    settings.ErrorLog.Println("Failed to delete key from Redis:", err)
	    return err
	}
	return nil
}
