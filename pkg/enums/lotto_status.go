package enums

type LottoStatus string

const (
	LottoStatusUnknown  LottoStatus = "unknown"
	LottoStatusNormal   LottoStatus = "normal"
	LottoStatusCanceled LottoStatus = "canceled"
)

func (LottoStatus) Values() []string {
	return []string{
		string(LottoStatusUnknown),
		string(LottoStatusNormal),
		string(LottoStatusCanceled),
	}
}

func (obj LottoStatus) Ptr() *LottoStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
