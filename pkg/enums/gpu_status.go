package enums

type GpuStatus string

const (
	GpuStatusOnline  GpuStatus = "online"
	GpuStatusOffline GpuStatus = "offline"
	GpuStatusBusy    GpuStatus = "busy"
	GpuStatusFree    GpuStatus = "free"
	GpuStatusExit    GpuStatus = "exit" // 退出节点网络
)

func (GpuStatus) Values() []string {
	return []string{
		string(GpuStatusOnline),
		string(GpuStatusOffline),
		string(GpuStatusBusy),
		string(GpuStatusFree),
		string(GpuStatusExit),
	}
}

func (obj GpuStatus) Ptr() *GpuStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
