package enums

type MissionBillingType string

const (
	MissionBillingTypeTime  MissionBillingType = "time"
	MissionBillingTypeCount MissionBillingType = "count"
)

func (obj MissionBillingType) Values() []string {
	return []string{
		string(MissionBillingTypeTime),
		string(MissionBillingTypeCount),
	}
}

func (obj MissionBillingType) Ptr() *MissionBillingType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
