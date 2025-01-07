package serializers


// хэш отловленной транзакции
type WaitNextOut struct {
	Hash string `json:"hash" example:"1a47b792640313b2d3ee6c05795364e95fedb9c6a55f8b9227fbfaa5c46c08ff" description:"хэш отловленной транзакции"`
}
