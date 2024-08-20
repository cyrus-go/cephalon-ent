package enums

type DeviceRankType string

const (
	DeviceRankTypeBlack          DeviceRankType = "black"            // 正常黑名单
	DeviceRankTypeNoRelieveBlack DeviceRankType = "no_relieve_black" // 不可自动解除黑名单状态
	DeviceRankTypeNormal         DeviceRankType = "normal"           // 正常
)

func (DeviceRankType) Values() []string {
	return []string{
		string(DeviceRankTypeBlack),
		string(DeviceRankTypeNoRelieveBlack),
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
