package enums

type IncomeManageStatus string

const (
	IncomeManageStatusPending  IncomeManageStatus = "pending"
	IncomeManageStatusCanceled IncomeManageStatus = "canceled"
	IncomeManageStatusSucceed  IncomeManageStatus = "succeed"
	IncomeManageStatusFailed   IncomeManageStatus = "failed"
	IncomeManageStatusReject   IncomeManageStatus = "reject"
)

func (IncomeManageStatus) Values() []string {
	return []string{
		string(IncomeManageStatusPending),
		string(IncomeManageStatusCanceled),
		string(IncomeManageStatusSucceed),
		string(IncomeManageStatusFailed),
		string(IncomeManageStatusReject),
	}
}

func (obj IncomeManageStatus) Ptr() *IncomeManageStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
