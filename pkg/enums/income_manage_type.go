package enums

type IncomeManageType string

const (
	IncomeManageTypeUnknown           IncomeManageType = "unknown"
	IncomeManageTypeIncomeReplacement IncomeManageType = "income_replacement"
	IncomeManageTypeIncomeDeduct      IncomeManageType = "income_deduct"
)

func (obj IncomeManageType) Values() []string {
	return []string{
		string(IncomeManageTypeUnknown),
		string(IncomeManageTypeIncomeReplacement),
		string(IncomeManageTypeIncomeDeduct),
	}
}

func (obj IncomeManageType) Ptr() *IncomeManageType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
