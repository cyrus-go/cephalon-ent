package enums

type MissionOrderStatus string

const (
	MissionOrderStatusUnknown   MissionOrderStatus = "unknown"
	MissionOrderStatusWaiting   MissionOrderStatus = "waiting"
	MissionOrderStatusCanceled  MissionOrderStatus = "canceled"
	MissionOrderStatusDoing     MissionOrderStatus = "doing"
	MissionOrderStatusSupplying MissionOrderStatus = "supplying"
	MissionOrderStatusFailed    MissionOrderStatus = "failed"
	MissionOrderStatusSucceed   MissionOrderStatus = "succeed"
)

func (obj MissionOrderStatus) Values() []string {
	return []string{
		string(MissionOrderStatusUnknown),
		string(MissionOrderStatusWaiting),
		string(MissionOrderStatusCanceled),
		string(MissionOrderStatusDoing),
		string(MissionOrderStatusSupplying),
		string(MissionOrderStatusFailed),
		string(MissionOrderStatusSucceed),
	}
}

func (obj MissionOrderStatus) Ptr() *MissionOrderStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
