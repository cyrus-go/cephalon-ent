package enums

type TransferOrderType string

const (
	TransferOrderTypeUnknown        TransferOrderType = "unknown"
	TransferOrderTypeRecharge       TransferOrderType = "recharge"
	TransferOrderTypeRechargeVX     TransferOrderType = "recharge_vx"
	TransferOrderTypeRechargeAlipay TransferOrderType = "recharge_alipay"
	TransferOrderTypeManual         TransferOrderType = "manual"
	TransferOrderTypeWithdraw       TransferOrderType = "withdraw"
)

func (TransferOrderType) Values() []string {
	return []string{
		string(TransferOrderTypeUnknown),
		string(TransferOrderTypeRecharge),
		string(TransferOrderTypeRechargeVX),
		string(TransferOrderTypeRechargeAlipay),
		string(TransferOrderTypeManual),
		string(TransferOrderTypeWithdraw),
	}
}

func (obj TransferOrderType) Ptr() *TransferOrderType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
