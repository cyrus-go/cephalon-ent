package enums

type MissionStatus string

const (
	MissionStatusWaiting   MissionStatus = "waiting"
	MissionStatusCanceled  MissionStatus = "canceled"
	MissionStatusDoing     MissionStatus = "doing"
	MissionStatusSupplying MissionStatus = "supplying"
	MissionStatusClosing   MissionStatus = "closing"
	MissionStatusDone      MissionStatus = "done"
)

func (m MissionStatus) Values() []string {
	return []string{
		string(MissionStatusWaiting),
		string(MissionStatusCanceled),
		string(MissionStatusDoing),
		string(MissionStatusSupplying),
		string(MissionStatusClosing),
		string(MissionStatusDone),
	}
}
