package enums

type MissionOrderStatus string

const (
	MissionOrderStatusWaiting   MissionOrderStatus = "waiting"
	MissionOrderStatusCanceled  MissionOrderStatus = "canceled"
	MissionOrderStatusDoing     MissionOrderStatus = "doing"
	MissionOrderStatusSupplying MissionOrderStatus = "supplying"
	MissionOrderStatusFailed    MissionOrderStatus = "failed"
	MissionOrderStatusSucceed   MissionOrderStatus = "succeed"
)

func (m MissionOrderStatus) Values() []string {
	return []string{
		string(MissionOrderStatusWaiting),
		string(MissionOrderStatusCanceled),
		string(MissionOrderStatusDoing),
		string(MissionOrderStatusSupplying),
		string(MissionOrderStatusFailed),
		string(MissionOrderStatusSucceed),
	}
}
