package enums

type MissionCallWay string

const (
	MissionCallWayApi         MissionCallWay = "api"
	MissionCallWayYuanHui     MissionCallWay = "yuan_hui"
	MissionCallWayDevPlatform MissionCallWay = "dev_platform"
)

func (obj MissionCallWay) Values() []string {
	return []string{
		string(MissionCallWayApi),
		string(MissionCallWayYuanHui),
		string(MissionCallWayDevPlatform),
	}
}

func (obj MissionCallWay) Ptr() *MissionCallWay {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
