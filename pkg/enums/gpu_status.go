package enums

type GpuStatus string

const (
	GpuStatusOnline  DeviceStatus = "online"
	GpuStatusOffline DeviceStatus = "offline"
	GpuStatusBusy    DeviceStatus = "busy"
	GpuStatusFree    DeviceStatus = "free"
)

func (GpuStatus) Values() []string {
	return []string{
		string(GpuStatusOnline),
		string(GpuStatusOffline),
		string(GpuStatusBusy),
		string(GpuStatusFree),
	}
}

func (obj GpuStatus) Ptr() *GpuStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
