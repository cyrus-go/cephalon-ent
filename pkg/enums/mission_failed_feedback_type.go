package enums

type MissionFailedFeedbackType string

const (
	MissionFailedFeedbackTypeUnknown MissionFailedFeedbackType = "unknown"
	MissionFailedFeedbackTypeStart   MissionFailedFeedbackType = "start"
	MissionFailedFeedbackTypeRecover MissionFailedFeedbackType = "recover"
)

func (obj MissionFailedFeedbackType) Values() []string {
	return []string{
		string(MissionFailedFeedbackTypeUnknown),
		string(MissionFailedFeedbackTypeStart),
		string(MissionFailedFeedbackTypeRecover),
	}
}

func (obj MissionFailedFeedbackType) Ptr() *MissionFailedFeedbackType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
