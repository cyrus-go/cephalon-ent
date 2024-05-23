package enums

type WithdrawStatus string

const (
	WithdrawStatusPending      WithdrawStatus = "pending"
	WithdrawStatusCanceled     WithdrawStatus = "canceled"
	WithdrawStatusSucceed      WithdrawStatus = "succeed"
	WithdrawStatusFailed       WithdrawStatus = "failed"
	WithdrawStatusReexchange   WithdrawStatus = "reexchange"
	WithdrawStatusPendingOrder WithdrawStatus = "pending_order"
	WithdrawStatusApproved     WithdrawStatus = "approved"
	WithdrawStatusReject       WithdrawStatus = "reject"
)

func (WithdrawStatus) Values() []string {
	return []string{
		string(WithdrawStatusPending),
		string(WithdrawStatusCanceled),
		string(WithdrawStatusSucceed),
		string(WithdrawStatusFailed),
		string(WithdrawStatusReexchange),
		string(WithdrawStatusPendingOrder),
		string(WithdrawStatusApproved),
		string(WithdrawStatusReject),
	}
}

func (obj WithdrawStatus) Ptr() *WithdrawStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
