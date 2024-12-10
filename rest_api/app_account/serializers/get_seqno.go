package serializers

// seqno аккаунта
type GetSeqnoOut struct {
	Seqno uint32 `json:"seqno" example:"105" description:"порядковый номер версии кошелька аккаунта"`
}
