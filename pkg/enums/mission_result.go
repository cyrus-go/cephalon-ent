package enums

type MissionResult string

const (
	MissionResultPending MissionResult = "pending"
	MissionResultSucceed MissionResult = "succeed"
	MissionResultFailed  MissionResult = "failed"
)

func (m MissionResult) Values() []string {
	return []string{
		string(MissionResultPending),
		string(MissionResultSucceed),
		string(MissionResultFailed),
	}
}
