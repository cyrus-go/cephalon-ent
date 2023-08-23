package enums

type MissionBillingType string

const (
	MissionBillingTypeTime  MissionBillingType = "time"
	MissionBillingTypeCount MissionBillingType = "count"
)

func (m MissionBillingType) Values() []string {
	return []string{
		string(MissionBillingTypeTime),
		string(MissionBillingTypeCount),
	}
}
