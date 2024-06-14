package enums

type WithdrawType string

const (
	WithdrawTypeUnknown        WithdrawType = "unknown"
	WithdrawTypeWithdraw       WithdrawType = "withdraw"
	WithdrawTypeWithdrawWX     WithdrawType = "withdraw_vx"
	WithdrawTypeWithdrawAlipay WithdrawType = "withdraw_alipay"
	WithdrawTypeBankCard       WithdrawType = "withdraw_bank_card"
	WithdrawTypeCompany        WithdrawType = "withdraw_company" // 对公提现
)

func (WithdrawType) Values() []string {
	return []string{
		string(WithdrawTypeUnknown),
		string(WithdrawTypeWithdraw),
		string(WithdrawTypeWithdrawWX),
		string(WithdrawTypeWithdrawAlipay),
		string(WithdrawTypeBankCard),
		string(WithdrawTypeCompany),
	}
}

func (obj WithdrawType) Ptr() *WithdrawType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
