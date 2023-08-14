package enums

type MissionCallWay string

const (
	MissionCallWayApi         MissionCallWay = "api"
	MissionCallWayYuanHui     MissionCallWay = "yuan_hui"
	MissionCallWayDevPlatform MissionCallWay = "dev_platform"
)

func (m MissionCallWay) Values() []string {
	return []string{
		string(MissionCallWayApi),
		string(MissionCallWayYuanHui),
		string(MissionCallWayDevPlatform),
	}
}
