package enums

type DeviceStatus string

const (
	DeviceStatusOnline  DeviceStatus = "online"
	DeviceStatusOffline DeviceStatus = "offline"
	DeviceStatusBusy    DeviceStatus = "busy"
	DeviceStatusFree    DeviceStatus = "free"
)

func (DeviceStatus) Values() []string {
	return []string{
		string(DeviceStatusOnline),
		string(DeviceStatusOffline),
		string(DeviceStatusBusy),
		string(DeviceStatusFree),
	}
}

func (obj DeviceStatus) Ptr() *DeviceStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
