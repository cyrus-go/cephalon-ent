package enums

type MissionCategory string

const (
	MissionCategoryUnknown MissionCategory = "unknown"
	MissionCategorySD      MissionCategory = "SD"
	MissionCategoryJP      MissionCategory = "JP"
	MissionCategoryWT      MissionCategory = "WT"
	MissionCategoryJPDK    MissionCategory = "JP_DK"
	MissionCategorySSH     MissionCategory = "SSH"
)

func (obj MissionCategory) Values() []string {
	return []string{
		string(MissionCategoryUnknown),
		string(MissionCategorySD),
		string(MissionCategoryJP),
		string(MissionCategoryWT),
		string(MissionCategoryJPDK),
		string(MissionCategorySSH),
	}
}

func (obj MissionCategory) Ptr() *MissionCategory {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
