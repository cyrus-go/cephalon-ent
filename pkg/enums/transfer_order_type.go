package enums

type TransferOrderType string

const (
	TransferOrderTypeUnknown           TransferOrderType = "unknown"
	TransferOrderTypeRecharge          TransferOrderType = "recharge"
	TransferOrderTypeRechargeVX        TransferOrderType = "recharge_vx"
	TransferOrderTypeRefund            TransferOrderType = "recharge_refund"
	TransferOrderTypeRechargeAlipay    TransferOrderType = "recharge_alipay"
	TransferOrderTypeRechargeAirwallex TransferOrderType = "recharge_airwallex" // 使用 airwallex 付款
	TransferOrderTypeManual            TransferOrderType = "manual"
	TransferOrderTypeWithdraw          TransferOrderType = "withdraw"
	TransferOrderTypeWithdrawWX        TransferOrderType = "withdraw_vx"
	TransferOrderTypeWithdrawAlipay    TransferOrderType = "withdraw_alipay"
	TransferOrderTypeBankCard          TransferOrderType = "withdraw_bank_card"
	TransferOrderTypeCompany           TransferOrderType = "withdraw_company"
)

func (TransferOrderType) Values() []string {
	return []string{
		string(TransferOrderTypeWithdrawWX),
		string(TransferOrderTypeWithdrawAlipay),
		string(TransferOrderTypeBankCard),
		string(TransferOrderTypeUnknown),
		string(TransferOrderTypeRecharge),
		string(TransferOrderTypeRechargeVX),
		string(TransferOrderTypeRechargeAlipay),
		string(TransferOrderTypeRechargeAirwallex),
		string(TransferOrderTypeManual),
		string(TransferOrderTypeWithdraw),
		string(TransferOrderTypeRefund),
		string(TransferOrderTypeCompany),
	}
}

func (obj TransferOrderType) Ptr() *TransferOrderType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
