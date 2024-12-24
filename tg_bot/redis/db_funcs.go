package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
)


// установка строкового значения в кэш
func SetString(key, value string) error {
	// если передана пустая строка
	if len(value) == 0 {
		redisErr := customErrors.RedisError("failed to set string value")
	    return fmt.Errorf("set %q value to %q: %w", key, value, fmt.Errorf("empty string was given: %w", redisErr))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// установка ключа-значения в redis
	err := GetRedisClient().Set(ctx, key, value, 0).Err()
	if err != nil {
	    redisErr := customErrors.RedisError("failed to set string value")
	    return fmt.Errorf("set %q value to %q: %w", key, value, fmt.Errorf("%v: %w", err, redisErr))
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
	    redisErr := customErrors.RedisError("failed to get string value")
	    return "", fmt.Errorf("get %q value: %w", key, fmt.Errorf("%v: %w", err, redisErr))
	}
	return value, nil
}


// установка значения среза строк в кэш
func SetStringSlice(key string, value []string) error {
	// если передан пустой список
	if len(value) == 0 {
	    redisErr := customErrors.RedisError("failed to set string slice value")
	    return fmt.Errorf("set %q value to %v: %w", key, value, fmt.Errorf("empty slice was given: %w", redisErr))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// сериализация среза строк в строковое значение
	marshaledValue, err := json.Marshal(value)
	if err != nil {
		internalErr := customErrors.InternalError("failed to marshal string slice")
	    return fmt.Errorf("set %q value to %v into redis: %w", key, value, fmt.Errorf("%v: %w", err, internalErr))
	}

	// установка ключа-значения в redis
	err = GetRedisClient().Set(ctx, key, marshaledValue, 0).Err()
	if err != nil {
	    redisErr := customErrors.RedisError("failed to set string slice value")
	    return fmt.Errorf("set %q value to %v: %w", key, value, fmt.Errorf("%v: %w", err, redisErr))
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
		redisErr := customErrors.RedisError("failed to get string slice value")
	    return unmarshaledValue, fmt.Errorf("get %q value: %w", key, fmt.Errorf("%v: %w", err, redisErr))
	}

	// десериализация строкового значения обратно в срез строк
	if err = json.Unmarshal([]byte(value), &unmarshaledValue); err != nil {
	    internalErr := customErrors.InternalError("failed to unmarshal string slice")
	    return unmarshaledValue, fmt.Errorf("get %q value from redis: %w", key, fmt.Errorf("%v: %w", err, internalErr))
	}
	return unmarshaledValue, nil
}


// удаление ключа из кэша
func DeleteKey(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// удаление значения из redis по ключу (первое возвращаемое значение - кол-во удалённых ключей)
	if _, err := GetRedisClient().Del(ctx, key).Result(); err != nil {
	    redisErr := customErrors.RedisError("failed to delete key")
	    return fmt.Errorf("delete %q key: %w", key, fmt.Errorf("%v: %w", err, redisErr))
	}
	return nil
}
