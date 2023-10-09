package enums

type RenewalPayStatus string

const (
	RenewalPayStatusUnknow  RenewalPayStatus = "unknow"
	RenewalPayStatusWaiting RenewalPayStatus = "waiting"
	RenewalPayStatusSucceed RenewalPayStatus = "succeed"
	RenewalPayStatusFailed  RenewalPayStatus = "failed"
)

func (obj RenewalPayStatus) Values() []string {
	return []string{
		string(RenewalPayStatusUnknow),
		string(RenewalPayStatusWaiting),
		string(RenewalPayStatusSucceed),
		string(RenewalPayStatusFailed),
	}
}

func (obj RenewalPayStatus) Ptr() *RenewalPayStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
