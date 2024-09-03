package enums

type MissionState string

const (
	MissionStateUnknown   MissionState = "unknown"
	MissionStateWaiting   MissionState = "waiting"
	MissionStateCanceled  MissionState = "canceled"
	MissionStateDoing     MissionState = "doing"
	MissionStateSupplying MissionState = "supplying"
	MissionStateClosing   MissionState = "closing"
	MissionStateSucceed   MissionState = "succeed"
	MissionStateFailed    MissionState = "failed"
	MissionStatePaused    MissionState = "paused"
)

func (obj MissionState) Values() []string {
	return []string{
		string(MissionStateUnknown),
		string(MissionStateWaiting),
		string(MissionStateCanceled),
		string(MissionStateDoing),
		string(MissionStateSupplying),
		string(MissionStateClosing),
		string(MissionStateSucceed),
		string(MissionStateFailed),
		string(MissionStatePaused),
	}
}

func (obj MissionState) Ptr() *MissionState {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
