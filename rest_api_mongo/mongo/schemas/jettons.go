package schemas

import (
	"github.com/google/uuid"
)


const jettonsCollection = "jettons"


// @Enum DEXesEnum
type DEXesEnum struct {
    DEXesEnum string `enum:"Ston.fi,Dedust.io" example:"Ston.fi" description:"допустимые значения DEX-биржи"`
}

// структура для данных о монетах для покупки
type Jetton struct {
	ID 			uuid.UUID 	`bson:"_id" json:"id" validate:"required" example:"715c0b81-bf1b-46c4-bf08-5c137cc6ec4d" description:"UUID записи" readOnly:"true"`
	Symbol	 	string 		`bson:"symbol" json:"symbol" validate:"required" example:"GRAM" description:"Название монеты"`
	JettonCA 	string 		`bson:"jettonCA" json:"jettonCA" validate:"required" example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"Мастер-адрес монеты (jetton_master)"`
	DEX 		string 		`bson:"dex" json:"dex" validate:"required,oneof=Ston.fi Dedust.io" example:"Ston.fi" description:"DEX-биржа" $ref:"DEXesEnum"`
}
func (this Jetton) CreatorCollectionName() string {
	return jettonsCollection
}
func (this Jetton) DataCollectionName() string {
	return jettonsCollection
}

// фильтр для поиска монет
type JettonFilter struct {
	ID 			*uuid.UUID 	`bson:"_id,omitempty" json:"id,omitempty" query:"id" example:"715c0b81-bf1b-46c4-bf08-5c137cc6ec4d" description:"UUID записи"`
	Symbol	 	*string 	`bson:"symbol,omitempty" json:"symbol,omitempty" query:"symbol" example:"GRAM" description:"Название монеты"`
	JettonCA 	*string		`bson:"jettonCA,omitempty" json:"jettonCA,omitempty" query:"jettonCA"example:"EQC47093oX5Xhb0xuk2lCr2RhS8rj-vul61u4W2UH5ORmG_O" description:"Мастер-адрес монеты (jetton_master)"`
	DEX 		*string		`bson:"dex,omitempty" json:"dex,omitempty" query:"dex" validate:"omitempty,oneof=Ston.fi Dedust.io" example:"Ston.fi" description:"DEX-биржа" $ref:"DEXesEnum"`
}
func (this JettonFilter) FilterCollectionName() string {
	return jettonsCollection
}
