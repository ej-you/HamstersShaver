package serializers

// @Description Seqno аккаунта
type GetSeqnoOut struct {
	// порядковый номер версии кошелька аккаунта
	Seqno string `json:"seqno" example:"105"`
}
