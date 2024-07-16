package enums

type DeviceStabilityType string

const (
	DeviceStabilityTypeGreat DeviceStabilityType = "great"
	DeviceStabilityTypeGood  DeviceStabilityType = "good"
	DeviceStabilityTypeOk    DeviceStabilityType = "ok"
	DeviceStabilityTypeBad   DeviceStabilityType = "bad"
)

func (DeviceStabilityType) Values() []string {
	return []string{
		string(DeviceStabilityTypeGreat),
		string(DeviceStabilityTypeGood),
		string(DeviceStabilityTypeOk),
		string(DeviceStabilityTypeBad),
	}
}

func (obj DeviceStabilityType) Ptr() *DeviceStabilityType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
