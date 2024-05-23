package enums

type WithdrawType string

const (
	WithdrawTypeUnknown        WithdrawType = "unknown"
	WithdrawTypeWithdraw       WithdrawType = "withdraw"
	WithdrawTypeWithdrawWX     WithdrawType = "withdraw_vx"
	WithdrawTypeWithdrawAlipay WithdrawType = "withdraw_alipay"
	WithdrawTypeBankCard       WithdrawType = "withdraw_bank_card"
)

func (WithdrawType) Values() []string {
	return []string{
		string(WithdrawTypeUnknown),
		string(WithdrawTypeWithdraw),
		string(WithdrawTypeWithdrawWX),
		string(WithdrawTypeWithdrawAlipay),
		string(WithdrawTypeBankCard),
	}
}

func (obj WithdrawType) Ptr() *WithdrawType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
