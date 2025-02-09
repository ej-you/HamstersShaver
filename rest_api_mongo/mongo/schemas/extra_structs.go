package schemas


// для Swagger документации
// @Enum TypesEnum
type TypesEnum struct {
    TypesEnum string `enum:"trade,auto" example:"trade" description:"допустимые значения действия типа записи транзакции"`
}

// для Swagger документации
// @Enum DEXesEnum
type DEXesEnum struct {
    DEXesEnum string `enum:"Ston.fi,Dedust.io" example:"Ston.fi" description:"допустимые значения DEX-биржи"`
}

// для Swagger документации
// @Enum ActionsEnum
type ActionsEnum struct {
    ActionsEnum string `enum:"buy,cell" example:"buy" description:"допустимые значения действия для транзакции"`
}

// для Swagger документации
// @Enum StatusesEnum
type StatusesEnum struct {
    StatusesEnum string `enum:"init,auto" example:"init" description:"допустимые значения статуса auto транзакции"`
}
