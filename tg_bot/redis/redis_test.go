package redis

import (
	"errors"
	"time"
	"fmt"

	"testing"
)


// обязательно указать эту переменную окружения перед командой запуска тестов (go test)
// ENV_FILE_PATH=../../


// маркеры успеха и неудачи
const (
    successMarker = "\u2713"
    failedMarker  = "\u2717"
)


var startTime time.Time

func logExecTime(t *testing.T, startTime *time.Time) {
	endTime := time.Now()
	t.Logf("\t\tExec time: %v", endTime.Sub(*startTime))
	*startTime = endTime
}

func SuccessLog(t *testing.T, format string, a ...any) {
	t.Logf("\t%s\t%s", successMarker, fmt.Sprintf(format, a...))
}

func ErrorLog(t *testing.T, err error) {
	t.Logf("\t%s\tFailed: %v", failedMarker, err)
}


// connection.go
func TestGetRedisClient(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test get redis connection")
	{
		redisConn := GetRedisClient()
		// если подключение не получится, то случится паника
		SuccessLog(t, "Successfully got redis connection: %v", redisConn)
	}
	logExecTime(t, &startTime)
}

// db_funcs.go
func TestStringCache(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test set string value in Redis")
	{
		err := SetString("user:123:status", "home")

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set string value to %q", "home")
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test get string value from Redis")
	{
		value, err := GetString("user:123:status")

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully got string value: %s", value)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test delete key with string value from Redis")
	{
		err := DeleteKey("user:123:status")

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully removed string value")
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test get removed string value from Redis")
	{
		value, err := GetString("user:123:status")

		if err != nil {
			SuccessLog(t, "Successfully got error while getting non-existent key: %v", err)
		} else {
			ErrorLog(t, errors.New("There is no error but there should be one. Value: " + value))
		}
	}
	logExecTime(t, &startTime)
}

// db_funcs.go
func TestStringSliceCache(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test set string slice value in Redis")
	{
		testStringSlice := []string{"banana", "strawberry", "apple", "tangerine"}

		err := SetStringSlice("user:123:transactions", testStringSlice)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set string slice value to %v", testStringSlice)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test get string slice value from Redis")
	{
		value, err := GetStringSlice("user:123:transactions")

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully got string slice value: %s", value)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test delete key with string slice value from Redis")
	{
		err := DeleteKey("user:123:transactions")

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully removed string slice value")
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test get removed string slice value from Redis")
	{
		value, err := GetStringSlice("user:123:transactions")

		if err != nil {
			SuccessLog(t, "Successfully got error while getting non-existent key: %v", err)
		} else {
			ErrorLog(t, errors.New(fmt.Sprintf("There is no error but there should be one. Value: %v", value)))
		}
	}
	logExecTime(t, &startTime)
}

// db_funcs.go
func TestSetEmptyValuesCache(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test set empty string value in Redis")
	{
		err := SetString("user:123:slippage", "")

		if err != nil {
			SuccessLog(t, "Successfully got error while setting empty string value: %v", err)
		} else {
			ErrorLog(t, errors.New("There is no error but there should be one."))
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test set empty string slice value in Redis")
	{
		err := SetStringSlice("user:123:transactions", []string{})

		if err != nil {
			SuccessLog(t, "Successfully got error while setting empty string slice value: %v", err)
		} else {
			ErrorLog(t, errors.New("There is no error but there should be one."))
		}
	}
	logExecTime(t, &startTime)
}
