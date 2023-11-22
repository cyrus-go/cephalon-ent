package enums

type GpuStatus string

const (
	GpuStatusOnline  GpuStatus = "online"
	GpuStatusOffline GpuStatus = "offline"
	GpuStatusBusy    GpuStatus = "busy"
	GpuStatusFree    GpuStatus = "free"
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
