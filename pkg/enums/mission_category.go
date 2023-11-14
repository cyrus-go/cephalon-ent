package enums

type MissionCategory string

const (
	MissionCategoryUnknown  MissionCategory = "unknown"
	MissionCategorySD       MissionCategory = "SD"
	MissionCategoryJP       MissionCategory = "JP"
	MissionCategoryWT       MissionCategory = "WT"
	MissionCategoryJPDK     MissionCategory = "JP_DK"
	MissionCategorySSH      MissionCategory = "SSH"
	MissionCategorySDTOMATO MissionCategory = "SD_TOMATO"
)

func (obj MissionCategory) Values() []string {
	return []string{
		string(MissionCategoryUnknown),
		string(MissionCategorySD),
		string(MissionCategoryJP),
		string(MissionCategoryWT),
		string(MissionCategoryJPDK),
		string(MissionCategorySSH),
		string(MissionCategorySDTOMATO),
	}
}

func (obj MissionCategory) Ptr() *MissionCategory {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
