package enums

type MissionResult string

const (
	MissionResultPending MissionResult = "pending"
	MissionResultSucceed MissionResult = "succeed"
	MissionResultFailed  MissionResult = "failed"
)

func (obj MissionResult) Values() []string {
	return []string{
		string(MissionResultPending),
		string(MissionResultSucceed),
		string(MissionResultFailed),
	}
}

func (obj MissionResult) Ptr() *MissionResult {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
