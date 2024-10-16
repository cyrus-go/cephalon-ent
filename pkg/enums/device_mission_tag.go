package enums

type DeviceMissionTag string

const (
	DeviceMissionTagNo   DeviceMissionTag = "no"   // 该标签表示只能接常规任务
	DeviceMissionTagAleo DeviceMissionTag = "aleo" // 该标签表示可以接 aleo 补贴任务
)

func (DeviceMissionTag) Values() []string {
	return []string{
		string(DeviceMissionTagNo),
		string(DeviceMissionTagAleo),
	}
}

func (obj DeviceMissionTag) Ptr() *DeviceMissionTag {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
