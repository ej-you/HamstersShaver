package state_machine

import (
	"errors"
	"fmt"
	"time"

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


var stateMachine UserStateMachine


// state_machine.go
func TestUserStateMachineStatus(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test set userTelegramID for stateMachine")
	{
		err := stateMachine.SetStatus("smth")
		if err != nil {
			SuccessLog(t, "Successfully got error while setting status for stateMachine: %v", err)
		} else {
			ErrorLog(t, errors.New("There is no error but there should be one"))
		}

		userTelegramID := "123456789"

		stateMachine.SetUserTelegramID(userTelegramID)

		if err = stateMachine.SetStatus("smth"); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set status %q for stateMachine", "smth")
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test success StatusEquals for stateMachine")
	{
		isEquals, err := stateMachine.StatusEquals("smth")
		if err != nil {
			ErrorLog(t, err)
		}
		if isEquals == true {
			SuccessLog(t, "Successfully compare status %q for stateMachine", "smth")
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test failed StatusEquals for stateMachine")
	{
		isEquals, err := stateMachine.StatusEquals("shit")
		if err != nil {
			ErrorLog(t, err)
		}
		if isEquals == false {
			SuccessLog(t, "Successfully compare status %q for stateMachine with status %q", "shit", "smth")
		}
	}
	logExecTime(t, &startTime)
}

// state_machine.go
func TestUserStateMachinePrepareTrans(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test set Action for stateMachine")
	{
		action := "buy"

		if err := stateMachine.SetAction(action); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set Action %q for stateMachine", action)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test set DEX for stateMachine")
	{
		dex := "stonfi"

		if err := stateMachine.SetDEX(dex); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set DEX %q for stateMachine", dex)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test set JettonsAmount for stateMachine")
	{
		jettonsAmount := "150"

		if err := stateMachine.SetJettonsAmount(jettonsAmount); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set JettonsAmount %q for stateMachine", jettonsAmount)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test set Slippage for stateMachine")
	{
		slippage := "20"

		if err := stateMachine.SetSlippage(slippage); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set Slippage %q for stateMachine", slippage)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test set JettonCA for stateMachine")
	{
		jettonCA := "EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O"

		if err := stateMachine.SetJettonCA(jettonCA); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully set JettonCA %q for stateMachine", jettonCA)
		}
	}
	logExecTime(t, &startTime)
}

// state_machine_transactions.go
func TestUserStateMachineTransactionPreparations(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test get new transaction preparation")
	{
		transInfo, err := stateMachine.GetNewTransactionPreparation()
		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully get new transaction preparation: %v", transInfo)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test clear transaction preparation")
	{
		err := stateMachine.ClearNewTransactionPreparation()
		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully clear transaction preparation")
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test get new transaction preparation (after cleaning)")
	{
		_, err := stateMachine.GetNewTransactionPreparation()
		if err != nil {
			SuccessLog(t, "Successfully got error while getting new transaction preparation: %v", err)
		} else {
			ErrorLog(t, errors.New("There is no error but there should be one"))
		}
	}
	logExecTime(t, &startTime)
}

// state_machine_transactions.go
func TestUserStateMachineTransactionUUIDs(t *testing.T) {
	startTime = time.Now()

	t.Logf("Test add new transaction UUIDs for stateMachine")
	{
		var err error

		firstUUID := "17cc93aa-872e-42a0-9689-2c8fa7b831ae"
		secondUUID := "482e1192-9345-4481-b66c-8c413c637097"

		if err = stateMachine.AddPendingTransaction(firstUUID); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully add new transaction UUID to empty slice: %v", firstUUID)
		}

		if err = stateMachine.AddPendingTransaction(secondUUID); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully add new transaction UUID to existing slice: %v", secondUUID)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test get transaction UUIDs from stateMachine (after adding)")
	{
		pendingTransactions, err := stateMachine.GetPendingTransactions()
		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully get transaction UUIDs: %v", pendingTransactions)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test delete transaction UUIDs from stateMachine")
	{
		var err error
		
		firstUUID := "17cc93aa-872e-42a0-9689-2c8fa7b831ae"
		secondUUID := "482e1192-9345-4481-b66c-8c413c637097"
		nonexistantUUID := "27806a35-d247-4b1d-bf28-1bac625631ac"

		if err = stateMachine.DeletePendingTransaction(nonexistantUUID); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully ignore deleting transaction UUID that is not in slice: %v", nonexistantUUID)
		}

		if err = stateMachine.DeletePendingTransaction(firstUUID); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully delete transaction UUID from slice: %v", firstUUID)
		}

		if err = stateMachine.DeletePendingTransaction(secondUUID); err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully delete last transaction UUID from slice: %v", secondUUID)
		}
	}
	logExecTime(t, &startTime)

	t.Logf("Test get transaction UUIDs from stateMachine (after deleting)")
	{
		pendingTransactions, err := stateMachine.GetPendingTransactions()
		if err != nil {
			ErrorLog(t, err)
		} else {
			SuccessLog(t, "Successfully get transaction UUIDs: %v", pendingTransactions)
		}
	}
	logExecTime(t, &startTime)
}
