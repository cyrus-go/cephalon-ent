package enums

type DeviceType string

const (
	DeviceTypeOfficial DeviceType = "official"
	DeviceTypeOrdinary DeviceType = "ordinary"
)

func (DeviceType) Values() []string {
	return []string{
		string(DeviceTypeOfficial),
		string(DeviceTypeOrdinary),
	}
}

func (obj DeviceType) Ptr() *DeviceType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
