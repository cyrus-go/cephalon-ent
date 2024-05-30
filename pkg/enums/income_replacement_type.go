package enums

type IncomeWalletOperateType string

const (
	IncomeWalletOperateTypeUnknown           IncomeWalletOperateType = "unknown"
	IncomeWalletOperateTypeIncomeReplacement IncomeWalletOperateType = "income_replacement"
	IncomeWalletOperateTypeIncomeDeduct      IncomeWalletOperateType = "income_deduct"
)

func (obj IncomeWalletOperateType) Values() []string {
	return []string{
		string(IncomeWalletOperateTypeUnknown),
		string(IncomeWalletOperateTypeIncomeReplacement),
		string(IncomeWalletOperateTypeIncomeDeduct),
	}
}

func (obj IncomeWalletOperateType) Ptr() *IncomeWalletOperateType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
