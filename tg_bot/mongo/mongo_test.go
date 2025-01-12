package mongo

import (
	"fmt"
	"time"
	"testing"

	"github.com/google/uuid"

	"github.com/ej-you/HamstersShaver/tg_bot/mongo/schemas"
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
func TestGetMongoClient(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test get mongo connection")
	{
		mongoConn := getMongoClient()
		// если подключение не получится, то случится паника
		SuccessLog(t, "Successfully got mongo connection: %v", mongoConn)
	}
	logExecTime(t, &startTime)
}

// db.go
func TestInsertOne(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test insert one Transaction document into mongo")
	{
		data := schemas.Transaction{
			UserID: "1601245210",
			ID: uuid.New(),
			Hash: "29a301e4d2a05713f4eab6c8f0daa3c58eed15d1d41678068cd50fe46ca7f6a5",
			
			Action: "cell",
			DEX: "Ston.fi",

			JettonCA: "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS",
			UsedJettons: "2000",

			Finished: true,
			Success: true,
		}
		err := NewMongoDB().InsertOne(data)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully insert one Transaction document")
		}
	}
	logExecTime(t, &startTime)
}

// db.go
func TestUpdateByID(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test update Transaction by its ID with given update data in mongo")
	{
		id := uuid.MustParse("27c9fa37-a35c-416a-a85c-93398e2659d4")
		updater := AnyCollectionData{"finished": false}

		err := NewMongoDB().UpdateByID("transactions", id, updater)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully update transaction by ID")
		}
	}
	logExecTime(t, &startTime)
}

// db.go
func TestGetTransactionByFilter(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test find Transaction with filter from mongo")
	{
		filter := AnyCollectionData{"jettonCA": "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"}

		var dataFound schemas.Transaction
		err := NewMongoDB().GetTransactionByFilter(filter, &dataFound)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully got transaction by filter: %v", dataFound)
		}
	}
	logExecTime(t, &startTime)
}

// db.go
func TestGetLastTransaction(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test last Transaction from mongo")
	{
		var dataFound schemas.Transaction
		err := NewMongoDB().GetLastTransaction(&dataFound)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully got last transaction: %v", dataFound)
		}
	}
	logExecTime(t, &startTime)
}
