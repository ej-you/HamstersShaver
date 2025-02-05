package schemas

import (
	"github.com/google/uuid"
)


const transactionsCollection = "transactions"


// информация о завершившейся транзакции закупки для auto функции
type InitTransactionInfo struct {
	Hash 		string 		`bson:"hash,omitempty" json:"hash,omitempty" example:"009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f" description:"Хэш первой операции"`
	LastTxHash 	string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty" example:"9a730767d311893fb6081b95663aee4d1d69c82f34a960057f15014e86b187c1" description:"Хэш последней операции (если транзакция завершена без ошибки)"`
}

// структура для создания записей об отдельных транзакциях в trading функции
type TransactionCreator struct {
	Type 		string 		`bson:"type" json:"type" validate:"required,oneof=trade auto" example:"trade" description:"Тип транзакции" $ref:"TypesEnum" readOnly:"true"`
	ID 			uuid.UUID 	`bson:"_id" json:"id" validate:"required" example:"715c0b81-bf1b-46c4-bf08-5c137cc6ec4d" description:"UUID записи" readOnly:"true"`
	UserID 		string 		`bson:"userID" json:"userID" validate:"required" example:"1601245210" description:"ID юзера"`
	JettonCA 	string 		`bson:"jettonCA" json:"jettonCA" validate:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"Мастер-адрес монеты (jetton_master)"`
	Action 		string 		`bson:"action" json:"action" validate:"required,oneof=buy cell" example:"buy" description:"Продажа/покупка монет" $ref:"ActionsEnum"`
	DEX 		string 		`bson:"dex" json:"dex" validate:"required,oneof=Ston.fi Dedust.io" example:"Ston.fi" description:"DEX-биржа" $ref:"DEXesEnum"`
	Hash 		string 		`bson:"hash" json:"hash" validate:"required" example:"009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f" description:"Хэш первой операции цепочки транзакций"`
	Finished 	bool 		`bson:"finished" json:"finished" example:"true" description:"Завершена ли транзакция (false по умолчанию)" default:"false"`

	UsedJettons string 		`bson:"usedJettons,omitempty" json:"usedJettons,omitempty" example:"2000" description:"Количество монет (для action == cell)"`
	UsedTON 	string 		`bson:"usedTon,omitempty" json:"usedTon,omitempty" example:"5.5" description:"Количество TON (для action == buy)"`
	Success 	bool 		`bson:"success,omitempty" json:"success,omitempty" example:"true" description:"true, если завершена без ошибки"`
	Error 		bool		`bson:"error,omitempty" json:"error,omitempty" example:"true" description:"Завершена ли от ошибки, не окончившись корректно"`
	LastTxHash 	string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty" example:"9a730767d311893fb6081b95663aee4d1d69c82f34a960057f15014e86b187c1" description:"Хэш последней операции (если завершена без ошибки)"`
}
func (this TransactionCreator) CreatorCollectionName() string {
	return transactionsCollection
}

// структура для создания записей о транзакциях в контексте auto функции
type TransactionAutoCreator struct {
	Type 		string 		`bson:"type" json:"type" validate:"required,oneof=trade auto" example:"auto" description:"Тип транзакции" $ref:"TypesEnum" readOnly:"true"`
	ID 			uuid.UUID 	`bson:"_id" json:"id" validate:"required" example:"715c0b81-bf1b-46c4-bf08-5c137cc6ec4d" description:"UUID записи" readOnly:"true"`
	UserID 		string 		`bson:"userID" json:"userID" validate:"required" example:"1601245210" description:"ID юзера"`
	UsedTON 	string 		`bson:"usedTon" json:"usedTon" validate:"required" example:"5.5" description:"Количество TON для конфигурации"`
	StopLoss	int 		`bson:"stopLoss" json:"stopLoss" validate:"required,min=1,max=100" example:"20" description:"Процент стоп-лосса" minimum:"1", maximum:"100"`
	TakeProfit	int 		`bson:"takeProfit" json:"takeProfit" validate:"required,min=1,max=100" example:"10" description:"Процент тейк-профита" minimum:"1", maximum:"100"`
	JettonCA 	string 		`bson:"jettonCA" json:"jettonCA" validate:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"Мастер-адрес монеты (jetton_master)"`
	Action 		string 		`bson:"action" json:"action" validate:"required,oneof=buy cell" example:"buy" description:"Продажа/покупка монет" $ref:"ActionsEnum"`
	DEX 		string 		`bson:"dex" json:"dex" validate:"required,oneof=Ston.fi Dedust.io" example:"Ston.fi" description:"DEX-биржа" $ref:"DEXesEnum"`
	Status		string		`bson:"status" json:"status" validate:"required,oneof=init auto" example:"init" description:"Какая из двух транзакций в процессе" $ref:"StatusesEnum" readOnly:"true"`
	Hash 		string 		`bson:"hash" json:"hash" validate:"required" example:"009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f" description:"Хэш первой операции"`
	Finished 	bool 		`bson:"finished" json:"finished" example:"true" description:"Завершена ли транзакция (false по умолчанию)" default:"false"`

	Success 	bool 		`bson:"success,omitempty" json:"success,omitempty" example:"true" description:"true, если транзакция завершена без ошибки"`
	Error 		bool		`bson:"error,omitempty" json:"error,omitempty" example:"true" description:"Завершена ли транзакция от ошибки, не окончившись корректно"`
	LastTxHash 	string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty" example:"9a730767d311893fb6081b95663aee4d1d69c82f34a960057f15014e86b187c1" description:"Хэш последней операции (если транзакция завершена без ошибки)"`
	
	InitTrans	*InitTransactionInfo `bson:"initTrans,omitempty" json:"initTrans,omitempty" description:"Информация о транзакции закупки (первой транзакции)"`
}
func (this TransactionAutoCreator) CreatorCollectionName() string {
	return transactionsCollection
}

// структура для получения записей о транзакциях
type Transaction struct {
	// обязательные поля для всех записей
	Type 		string 		`bson:"type" json:"type" example:"auto" description:"Тип транзакции" $ref:"TypesEnum"`
	ID 			uuid.UUID 	`bson:"_id" json:"id" example:"715c0b81-bf1b-46c4-bf08-5c137cc6ec4d" description:"UUID записи"`
	UserID 		string 		`bson:"userID" json:"userID" example:"1601245210" description:"ID юзера"`
	JettonCA 	string 		`bson:"jettonCA" json:"jettonCA" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"Мастер-адрес монеты (jetton_master)"`
	Action 		string 		`bson:"action" json:"action" example:"buy" description:"Продажа/покупка монет" $ref:"ActionsEnum"`
	DEX 		string 		`bson:"dex" json:"dex" example:"Ston.fi" description:"DEX-биржа" $ref:"DEXesEnum"`
	Hash 		string 		`bson:"hash" json:"hash" example:"009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f" description:"Хэш первой операции"`
	Finished 	bool 		`bson:"finished" json:"finished" example:"true" description:"Завершена ли транзакция"`
	// обязательные поля для auto функции
	UsedTON 	string 		`bson:"usedTon,omitempty" json:"usedTon,omitempty" example:"5.5" description:"Количество TON для конфигурации (тип auto) или для покупки (тип trade)"`
	Status		string		`bson:"status,omitempty" json:"status,omitempty" example:"init" description:"Какая из двух транзакций в процессе (тип auto)" $ref:"StatusesEnum"`
	StopLoss	int 		`bson:"stopLoss,omitempty" json:"stopLoss,omitempty" example:"20" description:"Процент стоп-лосса (тип auto)" minimum:"1", maximum:"100"`
	TakeProfit	int 		`bson:"takeProfit,omitempty" json:"takeProfit,omitempty" example:"10" description:"Процент тейк-профита (тип auto)" minimum:"1", maximum:"100"`
	// необязательные поля
	UsedJettons string 		`bson:"usedJettons,omitempty" json:"usedJettons,omitempty" example:"2000" description:"Количество монет для продажи (тип trade)"`
	Success 	bool 		`bson:"success,omitempty" json:"success,omitempty" example:"true" description:"true, если транзакция завершена без ошибки"`
	Error 		bool		`bson:"error,omitempty" json:"error,omitempty" example:"true" description:"Завершена ли транзакция от ошибки, не окончившись корректно"`
	LastTxHash 	string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty" example:"9a730767d311893fb6081b95663aee4d1d69c82f34a960057f15014e86b187c1" description:"Хэш последней операции (если транзакция завершена без ошибки)"`
	InitTrans	*InitTransactionInfo `bson:"initTrans,omitempty" json:"initTrans,omitempty" description:"Информация о транзакции закупки - первой транзакции (тип auto)"`
}
func (this Transaction) DataCollectionName() string {
	return transactionsCollection
}

// структура для обновления данных коллекции транзакций
type TransactionUpdater struct {
	Action 		*string		`bson:"action,omitempty" json:"action,omitempty" validate:"omitempty,oneof=buy cell" example:"buy" description:"Продажа/покупка монет" $ref:"ActionsEnum"`
	Hash 		*string		`bson:"hash,omitempty" json:"hash,omitempty" example:"009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f" description:"Хэш первой операции"`
	Finished 	*bool		`bson:"finished,omitempty" json:"finished,omitempty" example:"true" description:"Завершена ли транзакция (false по умолчанию)"`

	Success 	*bool		`bson:"success,omitempty" json:"success,omitempty" example:"true" description:"true, если транзакция завершена без ошибки"`
	Error 		*bool		`bson:"error,omitempty" json:"error,omitempty" example:"true" description:"Завершена ли транзакция от ошибки, не окончившись корректно"`
	LastTxHash 	*string		`bson:"lastTxHash,omitempty" json:"lastTxHash,omitempty" example:"9a730767d311893fb6081b95663aee4d1d69c82f34a960057f15014e86b187c1" description:"Хэш последней операции (если транзакция завершена без ошибки)"`

	UsedJettons *string		`bson:"usedJettons,omitempty" json:"usedJettons,omitempty" example:"2000" description:"Количество монет (для action == cell)"`
	UsedTON 	*string		`bson:"usedTon,omitempty" json:"usedTon,omitempty" example:"5.5" description:"Количество TON (для action == buy)"`

	Status		*string		`bson:"status,omitempty" json:"status,omitempty" validate:"omitempty,oneof=init auto" skip:"true"`
	InitTrans	*InitTransactionInfo `bson:"initTrans,omitempty" json:"initTrans,omitempty" description:"Информация о транзакции закупки (первой транзакции)" skip:"true"`
}
func (this TransactionUpdater) UpdateCollectionName() string {
	return transactionsCollection
}

// фильтр для поиска транзакций
type TransactionFilter struct {
	ID 			*uuid.UUID 	`bson:"_id,omitempty" json:"id,omitempty" query:"id" example:"715c0b81-bf1b-46c4-bf08-5c137cc6ec4d" description:"UUID записи"`
	Hash 		*string 	`bson:"hash,omitempty" json:"hash,omitempty" query:"hash" example:"009f801c3ab128fb53e5fca0ffe47b2dcfec3f6e28a07cf992ace5297363b72f" description:"Хэш первой операции цепочки транзакций"`
}
func (this TransactionFilter) FilterCollectionName() string {
	return transactionsCollection
}
