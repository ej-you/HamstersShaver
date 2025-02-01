package schemas

import (
	"github.com/google/uuid"
)


const transCollection = "transactions"


// информация о завершившейся транзакции закупки для auto функции
type InitTransactionInfo struct {
	Hash 		string 		`bson:"hash" json:"hash" validate:"required"` // хэш первой операции
	Success 	bool 		`bson:"success,omitempty" json:"success,omitempty"` // (если завершена без ошибки)
	Error 		bool		`bson:"error,omitempty" json:"error,omitempty"` // завершена ли от ошибки, не окончившись корректно
	LastTxHash 	string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty"` // хэш последней операции (если завершена без ошибки)
}

// структура для данных об отдельных транзакциях в trading функции
type Transaction struct {
	Type 		string 		`bson:"type" json:"type" validate:"required,oneof=trade auto"`
	ID 			uuid.UUID 	`bson:"_id" json:"id" validate:"required"`
	UserID 		string 		`bson:"userID" json:"userID" validate:"required"`
	JettonCA 	string 		`bson:"jettonCA" json:"jettonCA" validate:"required"`
	Action 		string 		`bson:"action" json:"action" validate:"required,oneof=buy cell"`
	DEX 		string 		`bson:"dex" json:"dex" validate:"required,oneof=Ston.fi Dedust.io"`
	Hash 		string 		`bson:"hash" json:"hash" validate:"required"` // хэш первой операции цепочки транзакций
	Finished 	bool 		`bson:"finished" json:"finished"` // завершена ли транзакция (false по умолчанию)

	UsedJettons string 		`bson:"usedJettons,omitempty" json:"usedJettons,omitempty"` // (для продажи)
	UsedTON 	string 		`bson:"usedTon,omitempty" json:"usedTon,omitempty"` // (для покупки)
	Success 	bool 		`bson:"success,omitempty" json:"success,omitempty"` // (если завершена без ошибки)
	Error 		bool		`bson:"error,omitempty" json:"error,omitempty"` // завершена ли от ошибки, не окончившись корректно
	LastTxHash 	string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty"` // хэш последней операции (если завершена без ошибки)
}
func (this Transaction) DataCollectionName() string {
	return transCollection
}

// структура для данных о транзакциях в контексте auto функции
type TransactionAuto struct {
	Type 		string 		`bson:"type" json:"type" validate:"required,oneof=trade auto"`
	ID 			uuid.UUID 	`bson:"_id" json:"id" validate:"required"`
	UserID 		string 		`bson:"userID" json:"userID" validate:"required"`
	UsedTON 	string 		`bson:"usedTon" json:"usedTon" validate:"required"` // TON для конфигурации
	StopLoss	int 		`bson:"stopLoss" json:"stopLoss" validate:"required,min=1,max=100"` // процент стоп-лосса
	TakeProfit	int 		`bson:"takeProfit" json:"takeProfit" validate:"required,min=1,max=100"` // процент тейк-профита
	JettonCA 	string 		`bson:"jettonCA" json:"jettonCA" validate:"required"`
	DEX 		string 		`bson:"dex" json:"dex" validate:"required,oneof=Ston.fi Dedust.io"`
	Status		string		`bson:"status" json:"status" validate:"required,oneof=init auto"` // какая из двух транзакций в процессе
	Hash 		string 		`bson:"hash" json:"hash" validate:"required"` // хэш первой операции
	Finished 	bool 		`bson:"finished" json:"finished"` // завершена ли транзакция (false по умолчанию)

	Success 	bool 		`bson:"success,omitempty" json:"success,omitempty"` // (если завершена без ошибки)
	Error 		bool		`bson:"error,omitempty" json:"error,omitempty"` // завершена ли от ошибки, не окончившись корректно
	LastTxHash 	string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty"` // хэш последней операции
	
	InitTrans	InitTransactionInfo `bson:"initTrans,omitempty" json:"initTrans,omitempty"` // информация о транзакции закупки
}
func (this TransactionAuto) DataCollectionName() string {
	return transCollection
}

// фильтр для поиска 
type TransactionFilter struct {
	ID 			*uuid.UUID 	`bson:"_id,omitempty" json:"id,omitempty"`
	Hash 		*string 	`bson:"hash,omitempty" json:"hash,omitempty"`
}
func (this TransactionFilter) FilterCollectionName() string {
	return transCollection
}

// структура для обновления данных коллекции транзакций
type TransactionUpdater struct {
	Finished 	*bool		`bson:"finished,omitempty" json:"finished,omitempty"`
	Success 	*bool		`bson:"success,omitempty" json:"success,omitempty"`
	Error 		*bool		`bson:"error,omitempty" json:"error,omitempty"`
	LastTxHash 	*string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty"`

	UsedJettons *string		`bson:"usedJettons,omitempty" json:"usedJettons,omitempty"`
	UsedTON 	*string		`bson:"usedTon,omitempty" json:"usedTon,omitempty"`

	Status		*string		`bson:"status,omitempty" json:"status,omitempty" validate:"oneof=init auto"`
	InitTrans	*InitTransactionInfo `bson:"initTrans,omitempty" json:"initTrans,omitempty"`
}
func (this TransactionUpdater) UpdateCollectionName() string {
	return transCollection
}
