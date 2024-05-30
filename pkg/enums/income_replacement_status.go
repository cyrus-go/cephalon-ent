package enums

type IncomeWalletOperateStatus string

const (
	IncomeWalletOperateStatusPending  IncomeWalletOperateStatus = "pending"
	IncomeWalletOperateStatusCanceled IncomeWalletOperateStatus = "canceled"
	IncomeWalletOperateStatusSucceed  IncomeWalletOperateStatus = "succeed"
	IncomeWalletOperateStatusFailed   IncomeWalletOperateStatus = "failed"
	IncomeWalletOperateStatusReject   IncomeWalletOperateStatus = "reject"
)

func (IncomeWalletOperateStatus) Values() []string {
	return []string{
		string(IncomeWalletOperateStatusPending),
		string(IncomeWalletOperateStatusCanceled),
		string(IncomeWalletOperateStatusSucceed),
		string(IncomeWalletOperateStatusFailed),
		string(IncomeWalletOperateStatusReject),
	}
}

func (obj IncomeWalletOperateStatus) Ptr() *IncomeWalletOperateStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
