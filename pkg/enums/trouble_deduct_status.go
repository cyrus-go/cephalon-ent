package enums

type TroubleDeductStatus string

const (
	TroubleDeductStatusPending  TroubleDeductStatus = "pending"
	TroubleDeductStatusCanceled TroubleDeductStatus = "canceled"
	TroubleDeductStatusSucceed  TroubleDeductStatus = "succeed"
	TroubleDeductStatusFailed   TroubleDeductStatus = "failed"
	TroubleDeductStatusReject   TroubleDeductStatus = "reject"
)

func (TroubleDeductStatus) Values() []string {
	return []string{
		string(TroubleDeductStatusPending),
		string(TroubleDeductStatusCanceled),
		string(TroubleDeductStatusSucceed),
		string(TroubleDeductStatusFailed),
		string(TroubleDeductStatusReject),
	}
}

func (obj TroubleDeductStatus) Ptr() *TroubleDeductStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
