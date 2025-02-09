package mongo

import (
	"errors"
	"fmt"
	"time"
	"testing"

	"github.com/google/uuid"

	coreValidator "github.com/ej-you/HamstersShaver/rest_api_mongo/core/validator"
	"github.com/ej-you/HamstersShaver/rest_api_mongo/mongo/schemas"
)


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
func TestInsertJetton(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test insert one Jetton document into mongo")
	{
		data := schemas.Jetton{
			ID: uuid.New(),
			Symbol: "DOGS",
			JettonCA: "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS",
			DEX: "Ston.fi",
		}
		// data := schemas.Jetton{
		// 	ID: uuid.New(),
		// 	Symbol: "NOT",
		// 	JettonCA: "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT",
		// 	DEX: "Ston.fi",
		// }
		// data := schemas.Jetton{
		// 	ID: uuid.New(),
		// 	Symbol: "DUST",
		// 	JettonCA: "EQBlqsm144Dq6SjbPI4jjZvA1hqTIP3CvHovbIfW_t-SCALE",
		// 	DEX: "Dedust.io",
		// }

		// валидация перед добавлением в БД
		err := coreValidator.GetValidator().Validate(&data)
		if err != nil {
			ErrorLog(t, err)
			return
		}

		// добавление в БД
		err = NewMongoDB().Insert(data)
		if err != nil {
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully insert one Jetton document")
	}
	logExecTime(t, &startTime)

	t.Logf("Test invalid Jetton data")
	{
		data := schemas.Jetton{
			ID: uuid.New(),
			Symbol: "DOGS",
			JettonCA: "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS",
			DEX: "AAA",
		}

		// валидация перед добавлением в БД
		err := coreValidator.GetValidator().Validate(&data)
		if err == nil { // NOT nil
			ErrorLog(t, fmt.Errorf("no one error"))
			return
		}
		SuccessLog(t, "Successfully got validation error: %v", err)
	}
	logExecTime(t, &startTime)
}

// db.go
func TestInsertTransaction(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test insert one Transaction document into mongo")
	{
		// data := schemas.TransactionCreator{
		// 	Type: "trade",
		// 	ID: uuid.New(),
		// 	UserID: "1601245210",
		// 	JettonCA: "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS",
		// 	Action: "cell",
		// 	DEX: "Ston.fi",
		// 	Hash: "99821a8101a7d25e76811e01e97d606d5caa5f62419867245b8fd1f5f362590b",
		// 	Finished: true,

		// 	UsedJettons: "2000",
		// 	Success: true,
		// 	LastTxHash: "29a301e4d2a05713f4eab6c8f0daa3c58eed15d1d41678068cd50fe46ca7f6a5",
		// }
		data := schemas.TransactionCreator{
			Type: "trade",
			ID: uuid.New(),
			UserID: "1601245210",
			JettonCA: "EQAvlWFDxGF2lXm67y4yzC17wYKD9A0guwPkMs1gOsM__NOT",
			Action: "buy",
			DEX: "Ston.fi",
			Hash: "009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f",
			Finished: true,

			UsedTON: "0.1",
			Success: true,
			LastTxHash: "9a730767d311893fb6081b95663aee4d1d69c82f34a960057f15014e86b187c1",
		}
		
		// добавление в БД
		err := NewMongoDB().Insert(data)
		if err != nil {
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully insert one Transaction document")
	}
	logExecTime(t, &startTime)
}

// db.go
func TestUpdateTransaction(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test update Transaction in mongo")
	{
		filterHash := "99821a8101a7d25e76811e01e97d606d5caa5f62419867245b8fd1f5f362590b"
		updateDataFinished := true
		updateDataError := false

		filter := schemas.TransactionFilter{
			Hash: &filterHash,
		}
		updateData := schemas.TransactionUpdater{
			Finished: &updateDataFinished,
			Error: &updateDataError,
		}

		// обновление записей в БД
		updatedAmount, err := NewMongoDB().Update(filter, updateData)
		if err != nil {
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully update Transaction documents (%d documents)", updatedAmount)
	}
	logExecTime(t, &startTime)
}

// db.go
func TestGetOneByFilterTransaction(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test find Transaction with filter (ID) from mongo")
	{
		filterID := uuid.MustParse("a83df76b-bf3b-4d80-95aa-f17f859e1191")
		filter := schemas.TransactionFilter{
			ID: &filterID,
		}

		var dataFound schemas.Transaction
		err := NewMongoDB().GetOneByFilter(filter, &dataFound)
		if err != nil {
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully got transaction by filter: %#v", dataFound)
	}
	logExecTime(t, &startTime)

	t.Logf("Test find Transaction with filter (hash) from mongo")
	{
		filterHash := "009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f"
		filter := schemas.TransactionFilter{
			Hash: &filterHash,
		}

		var dataFound schemas.Transaction
		err := NewMongoDB().GetOneByFilter(filter, &dataFound)
		if err != nil {
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully got transaction by filter: %#v", dataFound)
		SuccessLog(t, "Transaction UUID: %s", dataFound.ID.String())
	}
	logExecTime(t, &startTime)

	t.Logf("Test handling mongoErrNoDocuments error if found no one Transaction with filter from mongo")
	{
		filterHash := "test-unexisting-hash"
		filter := schemas.TransactionFilter{
			Hash: &filterHash,
		}

		var dataFound schemas.Transaction
		err := NewMongoDB().GetOneByFilter(filter, &dataFound)
		if err == nil || !errors.Is(err, ErrNoDocuments) { // NOT err OR not ErrNoDocuments error
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully got ErrNoDocuments error: %v", err)
	}
	logExecTime(t, &startTime)
}

// db.go
func TestGetManyByFilterJetton(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test get Jetton documents with filter from mongo")
	{
		filterDEX := "Ston.fi"
		filter := schemas.JettonFilter{
			DEX: &filterDEX,
		}

		var dataFound []schemas.Jetton
		err := NewMongoDB().GetManyByFilter(filter, &dataFound)

		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully got jettons: %#v", dataFound)
		}
	}
	logExecTime(t, &startTime)
}

// db.go
func TestDeleteJetton(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test delete Jetton documents with filter from mongo")
	{
		filterJettonCA := "EQCvxJy4eG8hyHBFsZ7eePxrRsUQSFE_jpptRAYBmcG_DOGS"
		filter := schemas.JettonFilter{
			JettonCA: &filterJettonCA,
		}

		// удаление записей из БД
		deletedAmount, err := NewMongoDB().Delete(filter)
		if err != nil {
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully deleted Jetton documents (%d documents)", deletedAmount)
	}
	logExecTime(t, &startTime)
}

// db.go
func TestDeleteTransaction(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test delete Transaction documents with filter from mongo")
	{
		filterID := uuid.MustParse("f55f1804-889d-4b85-87e8-5bd379449e38")
		filter := schemas.TransactionFilter{
			ID: &filterID,
		}

		// удаление записей из БД
		deletedAmount, err := NewMongoDB().Delete(filter)
		if err != nil {
			ErrorLog(t, err)
			return
		}
		SuccessLog(t, "Successfully deleted Transaction documents (%d documents)", deletedAmount)
	}
	logExecTime(t, &startTime)
}
