package enums

type MissionStatus string

const (
	MissionStatusWaiting   MissionStatus = "waiting"
	MissionStatusCanceled  MissionStatus = "canceled"
	MissionStatusDoing     MissionStatus = "doing"
	MissionStatusSupplying MissionStatus = "supplying"
	MissionStatusClosing   MissionStatus = "closing"
	MissionStatusDone      MissionStatus = "done"
	MissionStatusPaused    MissionStatus = "paused"
	MissionStatusPaused2   MissionStatus = "paused2"
)

func (obj MissionStatus) Values() []string {
	return []string{
		string(MissionStatusWaiting),
		string(MissionStatusCanceled),
		string(MissionStatusDoing),
		string(MissionStatusSupplying),
		string(MissionStatusClosing),
		string(MissionStatusDone),
		string(MissionStatusPaused),
		string(MissionStatusPaused2),
	}
}

func (obj MissionStatus) Ptr() *MissionStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
