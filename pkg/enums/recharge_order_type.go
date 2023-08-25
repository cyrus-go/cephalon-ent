package enums

type RechargeOrderType string

const (
	RechargeOrderTypeManual RechargeOrderType = "manual"
	RechargeOrderTypeVX     RechargeOrderType = "vx"
	RechargeOrderTypeAlipay RechargeOrderType = "alipay"
)

func (obj RechargeOrderType) Values() []string {
	return []string{
		string(RechargeOrderTypeManual),
		string(RechargeOrderTypeVX),
		string(RechargeOrderTypeAlipay),
	}
}

func (obj RechargeOrderType) Ptr() *RechargeOrderType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
