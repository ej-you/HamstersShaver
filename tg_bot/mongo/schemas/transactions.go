package schemas

import (
	"fmt"

	"github.com/google/uuid"

	customErrors "github.com/ej-you/HamstersShaver/tg_bot/errors"
)


type Transaction struct {
	UserID string `bson:"userID"`

	ID uuid.UUID `bson:"_id"`
	// (если завершена)
	Hash string `bson:"hash,omitempty"`
	
	Action string `bson:"action"`
	DEX string `bson:"dex"`

	JettonSymbol string `bson:"jettonSymbol"`
	JettonCA string `bson:"jettonCA"`
	// использованные монеты (для продажи)
	UsedJettons string `bson:"usedJettons,omitempty"`
	// использованные TON (для покупки)
	UsedTON string `bson:"usedTon,omitempty"`

	// завершена ли транзакция
	Finished bool `bson:"finished"`
	// true/false (если завершена без ошибки)
	Success bool `bson:"success,omitempty"`
	// (если возникла ошибка)
	Error string `bson:"error,omitempty"`
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
