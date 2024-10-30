package enums

type DeviceStatus string

const (
	DeviceStatusOnline    DeviceStatus = "online"
	DeviceStatusOffline   DeviceStatus = "offline"
	DeviceStatusBusy      DeviceStatus = "busy"      // 任务已接满
	DeviceStatusAvailable DeviceStatus = "available" // 可以接任务
	DeviceStatusFree      DeviceStatus = "free"      // 空闲
	DeviceStatusExit      DeviceStatus = "exit"      // 退出节点
)

func (DeviceStatus) Values() []string {
	return []string{
		string(DeviceStatusOnline),
		string(DeviceStatusOffline),
		string(DeviceStatusBusy),
		string(DeviceStatusFree),
		string(DeviceStatusExit),
		string(DeviceStatusAvailable),
	}
}

func (obj DeviceStatus) Ptr() *DeviceStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
