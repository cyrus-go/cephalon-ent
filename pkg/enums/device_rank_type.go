package enums

type DeviceRankType string

const (
	DeviceRankTypeBlack  DeviceRankType = "blask"  //黑名单
	DeviceRankTypeNormal DeviceRankType = "normal" //正常
)

func (DeviceRankType) Values() []string {
	return []string{
		string(DeviceRankTypeBlack),
		string(DeviceRankTypeNormal),
	}
}

func (obj DeviceRankType) Ptr() *DeviceRankType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
