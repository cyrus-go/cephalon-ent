package enums

type LottoCondition string

const (
	LottoConditionUnknown        LottoCondition = "unknown"
	LottoConditionRegister       LottoCondition = "register"
	LottoConditionInviteRegister LottoCondition = "invite_register"
	LottoConditionRecharge       LottoCondition = "recharge"
)

func (LottoCondition) Values() []string {
	return []string{
		string(LottoConditionUnknown),
		string(LottoConditionRegister),
		string(LottoConditionInviteRegister),
		string(LottoConditionRecharge),
	}
}

func (obj LottoCondition) Ptr() *LottoCondition {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
