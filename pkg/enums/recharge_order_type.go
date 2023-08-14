package enums

type RechargeOrderType string

const (
	RechargeOrderTypeManual RechargeOrderType = "manual"
	RechargeOrderTypeVX     RechargeOrderType = "vx"
	RechargeOrderTypeAlipay RechargeOrderType = "alipay"
)

func (m RechargeOrderType) Values() []string {
	return []string{
		string(RechargeOrderTypeManual),
		string(RechargeOrderTypeVX),
		string(RechargeOrderTypeAlipay),
	}
}
