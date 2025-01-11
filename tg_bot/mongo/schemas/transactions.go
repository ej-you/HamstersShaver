package schemas

import (
	"fmt"

	"github.com/google/uuid"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
)


type Transaction struct {
	UserID string `bson:"userID"`

	ID uuid.UUID `bson:"_id"`
	// (если завершена без ошибки)
	Hash string `bson:"hash,omitempty"`
	
	Action string `bson:"action"`
	DEX string `bson:"dex"`

	JettonCA string `bson:"jettonCA"`
	// использованные монеты (для продажи)
	UsedJettons string `bson:"usedJettons,omitempty"`
	// использованные TON (для покупки)
	UsedTON string `bson:"usedTon,omitempty"`

	// завершена ли транзакция
	Finished bool `bson:"finished"`
	// true/false (если завершена без ошибки)
	Success bool `bson:"success,omitempty"`
}
func (this Transaction) CollectionName() string {
	return "transactions"
}
func (this Transaction) Validate() error {
	if this.ID == [16]byte{} {
		internalErr := customErrors.InternalError("uuid is not set")
		return fmt.Errorf("validate struct: %w", internalErr)
	}
	return nil
}
