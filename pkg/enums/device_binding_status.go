package enums

type DeviceBindingStatus string

const (
	DeviceBindingStatusInit      DeviceBindingStatus = "init"
	DeviceBindingStatusBound     DeviceBindingStatus = "bound"
	DeviceBindingStatusUnbound   DeviceBindingStatus = "unbound"
	DeviceBindingStatusRebinding DeviceBindingStatus = "rebinding"
)

func (DeviceBindingStatus) Values() []string {
	return []string{
		string(DeviceBindingStatusInit),
		string(DeviceBindingStatusBound),
		string(DeviceBindingStatusUnbound),
		string(DeviceBindingStatusRebinding),
	}
}
