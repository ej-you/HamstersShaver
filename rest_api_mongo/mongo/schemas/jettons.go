package schemas

import (
	"github.com/google/uuid"
)


const jettonsCollection = "jettons"


// структура для данных о монетах для покупки
type Jetton struct {
	ID 			uuid.UUID 	`bson:"_id" json:"id" validate:"required"`
	Symbol	 	string 		`bson:"symbol" json:"symbol" validate:"required"`
	JettonCA 	string 		`bson:"jettonCA" json:"jettonCA" validate:"required"`
	DEX 		string 		`bson:"dex" json:"dex" validate:"required,oneof=Ston.fi Dedust.io"`
}
func (this Jetton) CreatorCollectionName() string {
	return jettonsCollection
}
func (this Jetton) DataCollectionName() string {
	return jettonsCollection
}

// фильтр для поиска монет
type JettonFilter struct {
	ID 			*uuid.UUID 	`bson:"_id,omitempty" json:"id,omitempty"`
	Symbol	 	*string 	`bson:"symbol,omitempty" json:"symbol,omitempty"`
	JettonCA 	*string		`bson:"jettonCA,omitempty" json:"jettonCA,omitempty"`
	DEX 		*string		`bson:"dex,omitempty" json:"dex,omitempty" validate:"omitempty,oneof=Ston.fi Dedust.io"`
}
func (this JettonFilter) FilterCollectionName() string {
	return jettonsCollection
}
