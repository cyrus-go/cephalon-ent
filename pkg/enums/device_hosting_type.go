package enums

type DeviceHostingType string

const (
	DeviceHostingTypeNo   DeviceHostingType = "no"   // 未托管
	DeviceHostingTypeHalf DeviceHostingType = "half" // 半托管
	DeviceHostingTypeFull DeviceHostingType = "full" // 全托管
)

func (DeviceHostingType) Values() []string {
	return []string{
		string(DeviceHostingTypeNo),
		string(DeviceHostingTypeHalf),
		string(DeviceHostingTypeFull),
	}
}

func (obj DeviceHostingType) Ptr() *DeviceHostingType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
