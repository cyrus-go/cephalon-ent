package enums

type MissionFailedFeedbackStatus string

const (
	MissionFailedFeedbackStatusInit MissionFailedFeedbackStatus = "init"
	MissionFailedFeedbackStatusDone MissionFailedFeedbackStatus = "done"
)

func (obj MissionFailedFeedbackStatus) Values() []string {
	return []string{
		string(MissionFailedFeedbackStatusInit),
		string(MissionFailedFeedbackStatusDone),
	}
}

func (obj MissionFailedFeedbackStatus) Ptr() *MissionFailedFeedbackStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
